package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Cannot Withdraw. Insufficient funds. Check Balance")

//Bitcoin struct
type Bitcoin float64

type Stringer interface {
	String() string
}

//Method to get the string form of Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

// Wallet struct
type Wallet struct {
	value Bitcoin
}

// Method to get wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.value
}

// Method to deposit amount into Wallet
func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount > 0.0 {
		w.value += amount
	}
	return nil
}

// Method to withdraw amount from Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.value { // && amount > 0
		return ErrInsufficientFunds
	}
	w.value -= amount
	return nil
}
func main() {
	w := Wallet{Bitcoin(10.00)}
	fmt.Printf("Initial balance BitCoin is %s \n", w.Balance())

	depositAmount := Bitcoin(100.0)
	w.Deposit(depositAmount)
	fmt.Printf("Balance after deposit of %s is %s \n", depositAmount, w.Balance())

	withDrawAmount := Bitcoin(1200.0)
	err := w.Withdraw(withDrawAmount)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Balance after withdraw of %s is %s \n", withDrawAmount, w.Balance())
	}
}
