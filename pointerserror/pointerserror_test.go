package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didnt get an error but wanted error")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(1.0))
		want := Bitcoin(1.0)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{value: Bitcoin(20.0)}
		wallet.Withdraw(Bitcoin(5.0))
		want := Bitcoin(15.0)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{value: Bitcoin(20.0)}
		err := wallet.Withdraw(Bitcoin(100.0))
		want := Bitcoin(20.0)
		assertBalance(t, wallet, want)
		assertError(t, err, ErrInsufficientFunds)
	})

}
