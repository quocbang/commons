package mail

import (
	"io"
	"os"
	"strconv"
	"strings"
)

// defines content type of the message
const (
	ContentHTML = "text/html"
)

type file struct {
	name     string
	readFunc func() (io.Reader, error)
}

// Message defines the message to be sent.
type Message struct {
	Recipients        []*Account
	CarbonCopies      []*Account
	BlindCarbonCopies []*Account
	Subject           string
	Body              string

	contentType string
	attachments []*file
	embeds      []*file
}

// NewMessage creates a message.
func NewMessage(contentType string) *Message {
	return &Message{
		contentType: contentType,
	}
}

// Attach attaches a file to the message.
func (m *Message) Attach(filename string) {
	m.attachments = append(m.attachments, &file{
		name: filename,
		readFunc: func() (io.Reader, error) {
			f, err := os.Open(filename)
			if err != nil {
				return nil, err
			}
			return f, nil
		},
	})
}

// AttachReader attaches a file using an io.Reader
func (m *Message) AttachReader(filename string, r io.Reader) {
	m.attachments = append(m.attachments, &file{
		name: filename,
		readFunc: func() (io.Reader, error) {
			return r, nil
		},
	})
}

// Embed embeds an image to the message and returns a Content-ID representing this image.
func (m *Message) Embed(filename string) string {
	name := m.renameEmbedFile(filename)
	m.embeds = append(m.embeds, &file{
		name: name,
		readFunc: func() (io.Reader, error) {
			f, err := os.Open(filename)
			if err != nil {
				return nil, err
			}
			return f, nil
		},
	})
	return name
}

// EmbedReader embeds an image using an io.Reader and returns a Content-ID representing this image.
func (m *Message) EmbedReader(filename string, r io.Reader) string {
	name := m.renameEmbedFile(filename)
	m.embeds = append(m.embeds, &file{
		name: name,
		readFunc: func() (io.Reader, error) {
			return r, nil
		},
	})
	return name
}

func (m *Message) renameEmbedFile(filename string) string {
	return "image" + strconv.Itoa(len(m.embeds)) + extractFilenameExtension(filename)
}

func extractFilenameExtension(filename string) string {
	i := strings.LastIndex(filename, ".")
	if i < 0 {
		return ""
	}
	return filename[i:]
}
