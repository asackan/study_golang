package bank

type Account struct {
	owner   string
	balance int
}

func (a Account) NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
