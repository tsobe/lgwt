package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		wallet.Deposit(Bitcoin(11))

		assertBalance(t, wallet, Bitcoin(21))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}

		err := wallet.Withdraw(Bitcoin(30))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func assertError(t *testing.T, got error, wantMsg error) {
	t.Helper()
	if got == nil {
		t.Errorf("Expected error")
	}
	if got != wantMsg {
		t.Errorf("Got %q, expected %q", got, wantMsg)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Error not expected, got %q", err)
	}
}
