package arrays

import "testing"

func TestBadBank(t *testing.T) {
	riya := Account{Name: "Riya", Balance: 100}
	chris := Account{Name: "Chris", Balance: 75}
	adil := Account{Name: "Adil", Balance: 200}

	transactions := []Transaction{
		NewTransaction(chris, riya, 100),
		NewTransaction(adil, chris, 25),
	}

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	assertEqual(t, newBalanceFor(riya), 200.0)
	assertEqual(t, newBalanceFor(chris), 0.0)
	assertEqual(t, newBalanceFor(adil), 175.0)
}
