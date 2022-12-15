package accounts

// Account
type Account struct {
	owner   string
	balance int
}

// /NewAccount creat Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
