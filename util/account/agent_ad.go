package account

import (
	"crypto/tls"
	"errors"
	"fmt"
	"strings"

	ldap "github.com/go-ldap/ldap"
)

// AgentViaADConfig the config for user agent via active directory.
var AgentViaADConfig ADConfig // deprecated with CheckAgentAD

var (
	ErrInternal            = errors.New("server error=")
	ErrInternalWithPayload = func(err error) error { return fmt.Errorf("%w%s", ErrInternal, err.Error()) }

	ErrLoginFailed            = errors.New("login filed error=")
	ErrLoginFailedWithPayLoad = func(err error) error { return fmt.Errorf("%w%s", ErrLoginFailed, err.Error()) }

	ErrQueryUserBindingFailed = errors.New("failed to bind binding user") // deprecated with CheckAgentAD
	ErrUserNotExist           = errors.New("user not exists")
	ErrMissingPassword        = errors.New("missing password")
	ErrMissingUser            = errors.New("missing user ID")
	ErrMissingRole            = errors.New("missing role ID")
	ErrMissingAttribute       = errors.New("missing attribute")
	ErrUserRolesNotExist      = errors.New("user roles not exists")
	ErrRoleMemberNotExist     = errors.New("role members not exists")
	ErrNotExistSearchSetting  = errors.New("not exist search setting")

	ErrCanNotFindADEntries    = errors.New("can't find any ldap entry")
	ErrCanNotFindValueOfCN    = errors.New("can't find any CN (common name) value from null ldap entry")
)

var (
	searchFilterByUserID = func(id string) string {
		return fmt.Sprintf("(&(objectClass=top)(objectClass=person)(objectClass=organizationalPerson)(objectClass=user)(sAMAccountName=%s))", id)
	}
	searchFilterByRole = func(role string) string {
		return fmt.Sprintf("(&(objectCategory=group)(cn=%s))", role)
	}
	searchTimeLimit = 10
)

// ADConfig is a config for active-directory.
type ADConfig struct {
	Host          string `long:"ad-host" description:"ad server hostname" env:"AD_HOST"`
	Port          int    `long:"ad-port" description:"ad server port" env:"AD_PORT"`
	DN            string `long:"ad-dn" description:"distinguished name of domain" env:"AD_DN"`
	QueryUser     string `long:"ad-user" description:"read only user name" env:"AD_QUERY_USER"`
	QueryPassword string `long:"ad-pwd" description:"read only user password" env:"AD_QUERY_PWD"`
	WithTLS       bool   `long:"ad-tls" description:"ad server security" env:"AD_TLS"`
}

// ADAgent is an agent via active-directory.
type ADAgent struct {
	config ADConfig
}

// NewADAgent creates a new active-directory agent. NewADAgent attempts to
// bind the QueryUser and the QueryPassword, returns ErrQueryUserBindingFailed
// if failed.
func NewADAgent(cfg ADConfig) (*ADAgent, error) {
	agent := &ADAgent{config: cfg}

	ad, err := agent.dialAndBind()
	if err != nil {
		return nil, err
	}
	ad.Close()

	return agent, nil
}

// Auth authenticates the id and the password. Auth returns an ErrUserNotExist error
// when the user is not found. If the search request returns too many entries, Auth
// will return an error because Auth does not know which entry to choose for authentication.
func (agent *ADAgent) Auth(id, pwd string) error {
	if strings.TrimSpace(id) == "" {
		return ErrMissingUser
	}
	if strings.TrimSpace(pwd) == "" {
		return ErrMissingPassword
	}

	ad, err := agent.dialAndBind()
	if err != nil {
		return err
	}
	defer ad.Close()

	// search login user
	result, err := ad.Search(&ldap.SearchRequest{
		BaseDN:       agent.config.DN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		TimeLimit:    searchTimeLimit,
		Filter:       searchFilterByUserID(id),
		Attributes:   []string{},
	})

	if err != nil {
		return ErrInternalWithPayload(err)
	}
	if len(result.Entries) == 0 {
		return ErrUserNotExist
	}
	if len(result.Entries) > 1 {
		return ErrInternalWithPayload(errors.New("too many entries returned"))
	}

	// bind as login user to verify their password
	if err = ad.Bind(result.Entries[0].DN, pwd); err != nil {
		return ErrLoginFailedWithPayLoad(err)
	}

	return nil
}

// CheckAgentAD authenticate user via ad. The function will be deprecated, use
// ADAgent.Auth instead.
func CheckAgentAD(id, pwd string) error {
	if strings.TrimSpace(id) == "" {
		return ErrMissingUser
	}
	if strings.TrimSpace(pwd) == "" {
		return ErrMissingPassword
	}

	// 0. setting ad connection
	var ad *ldap.Conn
	var err error
	address := fmt.Sprintf("%s:%d", AgentViaADConfig.Host, AgentViaADConfig.Port)
	if !AgentViaADConfig.WithTLS {
		ad, err = ldap.Dial("tcp", address)
	} else {
		ad, err = ldap.DialTLS("tcp", address, &tls.Config{ServerName: AgentViaADConfig.Host})
	}
	if err != nil {
		return err
	}
	defer ad.Close()

	// 1. bind read only user to get connection for search user
	err = ad.Bind(AgentViaADConfig.QueryUser, AgentViaADConfig.QueryPassword)
	if err != nil {
		return ErrQueryUserBindingFailed
	}

	// 2. search login user
	result, err := ad.Search(&ldap.SearchRequest{
		BaseDN:       AgentViaADConfig.DN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		TimeLimit:    searchTimeLimit,
		Filter:       searchFilterByUserID(id),
		Attributes:   []string{},
	})
	if err != nil {
		return ErrInternalWithPayload(err)
	}
	if len(result.Entries) == 0 {
		return ErrUserNotExist
	}
	if len(result.Entries) > 1 {
		return ErrInternalWithPayload(errors.New("too many entries returned"))
	}

	// 3. bind as login user to verify their password
	if err = ad.Bind(result.Entries[0].DN, pwd); err != nil {
		return ErrLoginFailedWithPayLoad(err)
	}
	return nil
}

// SearchUserRoles searches roles of the user. SearchUserRoles returns an
// ErrUserRolesNotExist error if there is no role to the user.
func (agent *ADAgent) SearchUserRoles(id string) (roles []string, err error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrMissingUser
	}

	result, err := agent.standardSearch(searchFilterByUserID(id), "memberOf")
	if err != nil {
		return nil, err
	}

	roles, err = findAttributeCNValues(result.Entries, "memberOf")
	if err != nil {
		return nil, ErrInternalWithPayload(err)
	}

	if len(roles) == 0 {
		return nil, ErrUserRolesNotExist
	}

	return roles, nil
}

// SearchRoleMembers searches members of the role.
// ErrRoleMemberNotExist error if there is no member to the role.
func (agent *ADAgent) SearchRoleMembers(role string) ([]string, error) {
	if strings.TrimSpace(role) == "" {
		return nil, ErrMissingRole
	}

	result, err := agent.standardSearch(searchFilterByRole(role), "member")
	if err != nil {
		return nil, ErrInternalWithPayload(err)
	}

	users, err := findAttributeCNValues(result.Entries, "member")
	if err != nil {
		return nil, ErrInternalWithPayload(err)
	}

	if len(users) == 0 {
		return nil, ErrRoleMemberNotExist
	}

	return users, nil
}

// standardSearch performs the standard search request with given filter and attribute
func (agent *ADAgent) standardSearch(filter string, attribute string) (*ldap.SearchResult, error) {
	ad, err := agent.dialAndBind()
	if err != nil {
		return nil, err
	}
	defer ad.Close()
	return ad.Search(&ldap.SearchRequest{
		BaseDN:       agent.config.DN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		TimeLimit:    searchTimeLimit,
		Filter:       filter,
		Attributes:   []string{attribute},
	})
}

// findAttributeCNValues find attribute CN (common Name) value from ldap entries.
// return an ErrCanNotFindADEntries error if there is no entry.
// return an ErrCanNotFindValueOfCN error if there is no common name of the attribute.
func findAttributeCNValues(entries []*ldap.Entry, attribute string) (values []string, err error) {
	if entries == nil {
		return nil, ErrCanNotFindADEntries
	}
	
	for _, entry := range entries {
		for _, attribute := range entry.GetAttributeValues(attribute) {
			for _, value := range strings.Split(attribute, ",") {
				// example: "CN=huei,OU=資訊一課,OU=資訊部,OU=管理本部,OU=總公司,OU=KENDA,DC=kenda,DC=com,DC=tw"
				if strings.Split(value, "=")[0] == "CN" {
					values = append(values, strings.Split(value, "=")[1])
					break
				}
			}
		}
	}

	if len(values) == 0 {
		return values, ErrCanNotFindValueOfCN
	}

	return values, nil
}

func (agent *ADAgent) dialAndBind() (*ldap.Conn, error) {
	ad, err := agent.dial()
	if err != nil {
		return nil, err
	}

	if err := ad.Bind(agent.config.QueryUser, agent.config.QueryPassword); err != nil {
		return nil, err
	}
	return ad, nil
}

func (agent *ADAgent) dial() (*ldap.Conn, error) {
	address := fmt.Sprintf("%s:%d", agent.config.Host, agent.config.Port)
	if !agent.config.WithTLS {
		return ldap.Dial("tcp", address)
	}
	return ldap.DialTLS("tcp", address, &tls.Config{ServerName: agent.config.Host})
}
