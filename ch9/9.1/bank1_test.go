package bank_test

import (
	"fmt"
	"testing"

	"github.com/cyberroadie/golangexercises/ch9/9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Olivier
	go func() {
		bank.Withdraw(100)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Wait for first withdrawel within limit
	<-done

	if got, want := bank.Balance(), 400; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Olivier
	var ok bool
	go func() {
		ok = bank.Withdraw(1000)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Wait for withdrawel out of limit
	<-done

	if ok {
		t.Error("To high a withdrawel was possible")
	}

	if got, want := bank.Balance(), 400; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

}
