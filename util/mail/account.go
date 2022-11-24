package mail

// defines kenda mail sender.
const (
	KD = iota
)

var (
	senders = []*Account{
		KD: &Account{Address: "kenda_kd@kenda.com.tw", Name: "kenda_kd[建大訊息通知]"},
	}
)

//Account defines email address and name.
type Account struct {
	Address, Name string
}
