package mail

import (
	"bytes"
	"io"
	"strconv"
)

// WriterAgent it does not send any email and uses io.Writer to write the message instead.
type WriterAgent struct {
	printer io.Writer
}

// NewWriterAgent creates a writer mail agent.
func NewWriterAgent(p io.Writer) *WriterAgent {
	return &WriterAgent{printer: p}
}

// Send sends the messages.
func (a *WriterAgent) Send(ms ...*Message) (err error) {
	var buf bytes.Buffer
	for _, m := range ms {
		a.writeAccountTo(&buf, "Recipients", m.Recipients)
		buf.WriteByte('\n')
		a.writeAccountTo(&buf, "Cc", m.CarbonCopies)
		buf.WriteByte('\n')
		a.writeAccountTo(&buf, "Bcc", m.BlindCarbonCopies)
		buf.WriteByte('\n')
		a.writeAttachmentAmountTo(&buf, len(m.attachments))
		buf.WriteByte('\n')
		a.writeSubjectTo(&buf, m.Subject)
		buf.WriteByte('\n')
		a.writeBodyTo(&buf, m.contentType, m.Body)
		buf.WriteByte('\n')
	}
	_, err = a.printer.Write(buf.Bytes())
	return
}

func (a *WriterAgent) writeAccountTo(buf *bytes.Buffer, name string, accounts []*Account) {
	buf.WriteString(name)
	buf.WriteString(": ")
	for _, account := range accounts {
		buf.WriteString(account.Address)
		buf.WriteString("; ")
	}
}

func (a *WriterAgent) writeAttachmentAmountTo(buf *bytes.Buffer, amount int) {
	buf.WriteString("Attachment Amount: ")
	buf.WriteString(strconv.Itoa(amount))
}

func (a *WriterAgent) writeSubjectTo(buf *bytes.Buffer, subject string) {
	buf.WriteString("Subject: ")
	buf.WriteString(subject)
}

func (a *WriterAgent) writeBodyTo(buf *bytes.Buffer, contentType, body string) {
	buf.WriteString("Body(")
	buf.WriteString(contentType)
	buf.WriteString("): ")
	buf.WriteString(body)
}
