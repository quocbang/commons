package mail

import (
	gomail "github.com/go-mail/mail"
)

// SendMail 寄信func
//func SendMail(subject string, message string，address [][]string, ccaddr[][]string) {
func SendMail(subject string, message string, address, ccaddr, bccaddr []Account, filepath []string) error {
	m := gomail.NewMessage()
	//設定寄件者資料
	m.SetHeader("From", m.FormatAddress("kenda@kenda.com.tw", "kenda_kd [建大訊息通知]"))
	//依傳入的header設定主旨，
	m.SetHeader("Subject", subject)

	//TO的人員可能是多筆的，用range來處理，提供人名設定
	for _, row := range address {
		//m.SetHeader("To", address[i][0])
		//m.SetHeader("To", m.FormatAddress(address[i][0], address[i][1]))
		m.SetAddressHeader("To", row.Address, row.Name)
	}

	//CC人員
	for _, row := range ccaddr {
		//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetAddressHeader("Cc", row.Address, row.Name)
	}

	//BCC人員
	for _, row := range bccaddr {
		//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetAddressHeader("Bcc", row.Address, row.Name)
	}

	// 內文，傳入的message，可使用html語法
	m.SetBody("text/html", message)

	// 附件檔案，正確的url或是WERP server上對應的路徑, 目前只管控了有傳入路徑時就附加
	// TODO: 管控檔案的大小，不能超過多少之類的吧~
	// 多檔的話，請用分號組字串進來，或是再作調整吧~
	for _, path := range filepath {
		m.Attach(path)
	}

	//剛開始使用時有發生X509憑證問題，後來加了TLSConfig設定後就OK了~
	d := gomail.Dialer{Host: "kmail.kenda.com.tw", Port: 25}
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
