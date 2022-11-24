package account

import (
	"crypto/tls"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"bou.ke/monkey"
	ldap "github.com/go-ldap/ldap"
	"github.com/stretchr/testify/assert"
)

var userList = map[string]string{
	"query": "12345",
	"user1": "11111",
}

func configureConfig(
	server string,
	port int,
	queryPassword, queryUser, DN string,
	withTLS bool) {
	AgentViaADConfig.Host = server
	AgentViaADConfig.Port = port
	AgentViaADConfig.QueryPassword = queryPassword
	AgentViaADConfig.QueryUser = queryUser
	AgentViaADConfig.DN = DN
	AgentViaADConfig.WithTLS = withTLS
}

func checkAgentMockPatch(c *ldap.Conn) {
	monkey.Patch(ldap.Dial, func(network, addr string) (*ldap.Conn, error) {
		address := strings.Split(addr, ":")
		if address[0] != "host" {
			return nil, fmt.Errorf("not found %s", address[0])
		}
		return &ldap.Conn{}, nil
	})

	monkey.Patch(ldap.DialTLS, func(network, addr string, config *tls.Config) (*ldap.Conn, error) {
		address := strings.Split(addr, ":")
		if address[0] != "host" {
			return nil, fmt.Errorf("not found %s", address[0])
		}
		return &ldap.Conn{}, nil
	})

	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Bind", func(c *ldap.Conn, username, password string) error {
		pwd, ok := userList[username]
		if !ok {
			return fmt.Errorf("not found username:%s", username)
		}
		if pwd != password {
			return fmt.Errorf("wrong password")
		}
		return nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "SimpleBind", func(c *ldap.Conn, request *ldap.SimpleBindRequest) (*ldap.SimpleBindResult, error) {
		pwd, ok := userList[request.Username]
		if !ok {
			return nil, fmt.Errorf("not found username:%s", request.Username)
		}
		if pwd != request.Password {
			return nil, fmt.Errorf("wrong password")
		}
		return nil, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Close", func(c *ldap.Conn) {})
}

// searchWithError with internal error.
func searchWithError(c *ldap.Conn) {
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Search", func(c *ldap.Conn, request *ldap.SearchRequest) (*ldap.SearchResult, error) {
		return nil, fmt.Errorf("search error")
	})
}

// searchUser1 will return user1 as result.
func searchUser1(c *ldap.Conn) {
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Search", func(c *ldap.Conn, request *ldap.SearchRequest) (*ldap.SearchResult, error) {
		return &ldap.SearchResult{
			Entries: []*ldap.Entry{
				{
					DN: "user1",
				},
			},
		}, nil
	})
}

// searchMultipleEntries return multiple entries result.
func searchMultipleEntries(c *ldap.Conn) {
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Search", func(c *ldap.Conn, request *ldap.SearchRequest) (*ldap.SearchResult, error) {
		return &ldap.SearchResult{
			Entries: []*ldap.Entry{
				{
					DN: "user1",
				},
				{
					DN: "xxxx",
				},
			},
		}, nil
	})
}

// searchWithoutEntries with zero-length entry.
func searchWithoutEntries(c *ldap.Conn) {
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Search", func(c *ldap.Conn, request *ldap.SearchRequest) (*ldap.SearchResult, error) {
		return &ldap.SearchResult{}, nil
	})
}

func TestCheckAgentAD(t *testing.T) {
	configureConfig("host", 8888, "12345", "query", "cd=com,cd=tw", false)
	var c *ldap.Conn
	var err error
	checkAgentMockPatch(c)
	defer monkey.UnpatchAll()
	asserts := assert.New(t)
	{ // success without TLS
		searchUser1(c)
		id, password := "user1", "11111"
		err = CheckAgentAD(id, password)
		asserts.NoError(err)
	}
	{ // missing user
		searchUser1(c)
		id, password := "", "0000"
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, ErrMissingUser.Error())
	}
	{ // missing password
		searchUser1(c)
		id, password := "user1", ""
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, ErrMissingPassword.Error())
	}
	{ // wrong password
		searchUser1(c)
		id, password := "user1", "12345"
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, ErrLoginFailedWithPayLoad(fmt.Errorf("wrong password")).Error())
	}
	{ // too many entries
		searchMultipleEntries(c)
		id, password := "user1", "11111"
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, ErrInternalWithPayload(errors.New("too many entries returned")).Error())
	}
	{ // user does not exist
		searchWithoutEntries(c)
		id, password := "user1", "11111"
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, ErrUserNotExist.Error())
	}
	{ // failed to check user
		searchWithError(c)
		id, password := "user1", "11111"
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, ErrInternalWithPayload(fmt.Errorf("search error")).Error())
	}
	{ // failed binding user
		searchUser1(c)
		id, password := "user1", "11111"
		configureConfig("host", 8888, "12345", "qad", "cd=com,cd=tw", false)
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, ErrQueryUserBindingFailed.Error())
	}
	{ // not found server
		searchUser1(c)
		id, password := "user1", "11111"
		configureConfig("server", 8888, "12345", "query", "cd=com,cd=tw", false)
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, fmt.Errorf("not found server").Error())
	}
	{ // success case with TLS
		searchUser1(c)
		id, password := "user1", "11111"
		configureConfig("host", 8888, "12345", "query", "cd=com,cd=tw", true)
		err := CheckAgentAD(id, password)
		asserts.NoError(err)
	}
	{ // not found server
		searchUser1(c)
		id, password := "user1", "11111"
		configureConfig("server", 8888, "12345", "query", "cd=com,cd=tw", true)
		err = CheckAgentAD(id, password)
		asserts.EqualError(err, fmt.Errorf("not found server").Error())
	}
}

func TestADAgent(t *testing.T) {
	assert := assert.New(t)
	var c *ldap.Conn
	checkAgentMockPatch(c)
	defer monkey.UnpatchAll()
	{ // success without TLS
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.NoError(err)
		id, password := "user1", "11111"
		err = ad.Auth(id, password)
		assert.NoError(err)
	}
	{ // missing user
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.NoError(err)
		id, password := "", "0000"
		err = ad.Auth(id, password)
		assert.EqualError(err, ErrMissingUser.Error())
	}
	{ // missing password
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.NoError(err)
		id, password := "user1", ""
		err = ad.Auth(id, password)
		assert.EqualError(err, ErrMissingPassword.Error())
	}
	{ // wrong password
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.NoError(err)
		id, password := "user1", "12345"
		err = ad.Auth(id, password)
		assert.EqualError(err, ErrLoginFailedWithPayLoad(fmt.Errorf("wrong password")).Error())
	}
	{ // too many entries
		searchMultipleEntries(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.NoError(err)
		id, password := "user1", "11111"
		err = ad.Auth(id, password)
		assert.EqualError(err, ErrInternalWithPayload(errors.New("too many entries returned")).Error())
	}
	{ // user does not exist
		searchWithoutEntries(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.NoError(err)
		id, password := "user1", "11111"
		err = ad.Auth(id, password)
		assert.EqualError(err, ErrUserNotExist.Error())
	}
	{ // failed to check user
		searchWithError(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.NoError(err)
		id, password := "user1", "11111"
		err = ad.Auth(id, password)
		assert.EqualError(err, ErrInternalWithPayload(fmt.Errorf("search error")).Error())
	}
	{ // failed binding user
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "qad",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.EqualError(err, "not found username:qad")
		assert.Nil(ad)
	}
	{ // not found server
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "server",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       false,
		})
		assert.EqualError(err, "not found server")
		assert.Nil(ad)
	}
	{ // success case with TLS
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "host",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       true,
		})
		assert.NoError(err)
		id, password := "user1", "11111"
		err = ad.Auth(id, password)
		assert.NoError(err)
	}
	{ // not found server
		searchUser1(c)
		ad, err := NewADAgent(ADConfig{
			Host:          "server",
			Port:          8888,
			DN:            "cd=com,cd=tw",
			QueryUser:     "query",
			QueryPassword: "12345",
			WithTLS:       true,
		})
		assert.EqualError(err, "not found server")
		assert.Nil(ad)
	}
}
