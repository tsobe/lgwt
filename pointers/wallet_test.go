package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))
	wallet.Deposit(Bitcoin(11))

	got := wallet.Balance()
	want := Bitcoin(21)

	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}
