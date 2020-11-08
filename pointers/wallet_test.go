package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertbalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertbalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertbalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertbalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	actual := wallet.Balance()

	if actual != expected {
		t.Errorf("got %s expected %s", actual, expected)
	}
}

func assertError(t *testing.T, actual error, expected error) {
	t.Helper()
	if actual == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if actual != expected {
		t.Errorf("got %q, expected %q", actual, expected)
	}
}
