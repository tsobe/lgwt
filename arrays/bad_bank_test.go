package arrays

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	assertEqual(t, BalanceFor(transactions, "Riya"), 100.0)
	assertEqual(t, BalanceFor(transactions, "Chris"), -75.0)
	assertEqual(t, BalanceFor(transactions, "Adil"), -25.0)
}
