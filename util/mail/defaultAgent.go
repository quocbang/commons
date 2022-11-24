package mail

import (
	"crypto/tls"
	"errors"
	"io"

	gomail "github.com/go-mail/mail"
)

// DefaultAgent a default email agent.
type DefaultAgent struct {
	InsecureSkipVerify bool

	developMode bool
	developer   []*Account

	host     string
	port     int
	username string
	password string
	sender   *Account
}

// NewDefaultAgent creates a default email agent, the sender is always kenda_kd@kenda.com.tw.
func NewDefaultAgent(host string, port int, username, password string) *DefaultAgent {
	return &DefaultAgent{
		host:     host,
		port:     port,
		username: username,
		password: password,
		sender:   senders[KD],
	}
}

// DevelopmentMode uses the develop mode, all messages will be sent to developer.
func (a *DefaultAgent) DevelopmentMode(developer []*Account) error {
	if len(developer) == 0 {
		return errors.New("the developer is absent")
	}
	a.developMode = true
	a.developer = developer
	return nil
}

// Send sends the messages.
func (a *DefaultAgent) Send(ms ...*Message) (err error) {
	msgs := make([]*gomail.Message, len(ms))
	for i, m := range ms {
		msg := gomail.NewMessage()

		msg.SetHeader(headerFrom, msg.FormatAddress(a.sender.Address, a.sender.Name))

		if a.developMode {
			if recipient := a.handleAccount(a.developer, msg); len(recipient) > 0 {
				msg.SetHeader(headerRecipient, recipient...)
			}
		} else {
			if recipient := a.handleAccount(m.Recipients, msg); len(recipient) > 0 {
				msg.SetHeader(headerRecipient, recipient...)
			}

			if cc := a.handleAccount(m.CarbonCopies, msg); len(cc) > 0 {
				msg.SetHeader(headerCarbonCopy, cc...)
			}

			if bcc := a.handleAccount(m.BlindCarbonCopies, msg); len(bcc) > 0 {
				msg.SetHeader(headerBlindCarbonCopy, bcc...)
			}
		}

		msg.SetHeader(headerSubject, m.Subject)
		msg.SetBody(m.contentType, m.Body)

		for _, att := range m.attachments {
			var r io.Reader
			if r, err = att.readFunc(); err != nil {
				return
			}
			msg.AttachReader(att.name, r)
		}

		for _, embed := range m.embeds {
			var r io.Reader
			if r, err = embed.readFunc(); err != nil {
				return
			}
			msg.EmbedReader(embed.name, r)
		}

		msgs[i] = msg
	}

	d := gomail.NewDialer(a.host, a.port, a.username, a.password)
	if a.InsecureSkipVerify {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return d.DialAndSend(msgs...)
}

func (a *DefaultAgent) handleAccount(accs []*Account, msg *gomail.Message) []string {
	arr := make([]string, len(accs))
	for i, acc := range accs {
		arr[i] = msg.FormatAddress(acc.Address, acc.Name)
	}
	return arr
}
