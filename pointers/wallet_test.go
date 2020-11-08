package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	actual := wallet.Balance()
	expected := Bitcoin(10)

	if actual != expected {
		t.Errorf("got %s expected %s", actual, expected)
	}
}
