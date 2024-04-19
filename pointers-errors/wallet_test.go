package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))
		if err != nil {
			t.Fatalf("got an error but didn't want one: %v", err)
		}
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		starting_balance := Bitcoin(20)
		wallet := Wallet{balance: starting_balance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, starting_balance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
