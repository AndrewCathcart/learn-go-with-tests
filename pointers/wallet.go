package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin is represented as an int.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet contains a balance in Bitcoin.
type Wallet struct {
	balance Bitcoin
}

// Deposit adds Bitcoin to the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds should be returned when attempting to withdraw more than the current wallet balance.
var ErrInsufficientFunds = errors.New("insufficient funds")

// Withdraw Bitcoin from the wallet. Returns an ErrInsufficientFunds error if attempting to withdraw more than the balance.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the number of Bitcoins in the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
