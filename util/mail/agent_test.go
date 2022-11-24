package mail

import (
	"os"
	"testing"
)

func TestSend(t *testing.T) {
	// registers a email agent
	a := NewWriterAgent(os.Stdout)
	RegisterAgent(a)

	// defines the message
	m := NewMessage(ContentHTML)
	m.Recipients = []*Account{
		&Account{Address: "aaa@kenda.com.tw"},
		&Account{Address: "abcd@kenda.com.tw"},
	}
	m.CarbonCopies = []*Account{
		&Account{Address: "cca@kenda.com.tw"},
	}
	m.BlindCarbonCopies = []*Account{
		&Account{Address: "xyz@kenda.com.tw"},
		&Account{Address: "zz@kenda.com.tw"},
	}
	m.Subject = "[test]go mail test"
	cid0 := m.Embed("C:/Users/user/Desktop/測試0.png")
	cid1 := m.Embed("C:/Users/user/Desktop/測試1.png")
	m.Body = "<p>test</p><img src=\"cid:" + cid0 + "\" alt=\"my test image0\" /><p>測試</p><img src=\"cid:" + cid1 + "\" alt=\"my test image1\" />"
	m.Attach("C:/Users/user/Desktop/測試附件.png")

	err := A().Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
