package bank

var deposits = make(chan int)    // send amount to deposit
var balances = make(chan int)    // receive balance
var withdrawels = make(chan int) // withdraw money

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	if <-balances-amount < 0 {
		return false
	}
	deposits <- amount
	return true
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amountd := <-deposits:
			balance += amountd
		case amountw := <-withdrawels:
			balance -= amountw
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
