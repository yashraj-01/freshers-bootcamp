package main

import (
	"fmt"
	"sync"
)

type Transaction struct {
	id  int32
	t   int32 // 0: credit, 1: debit
	amt int32
}

type Account struct {
	mu      sync.Mutex
	balance int32
}

func (a *Account) withdraw(amt int32) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amt > a.balance {
		fmt.Println("Insufficient balance")
	} else {
		a.balance -= amt
		fmt.Println("Debited:", amt, "Balance:", a.balance)
	}
}

func (a *Account) deposit(amt int32) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amt
	fmt.Println("Credited:", amt, "Balance:", a.balance)
}

func (a *Account) transact(t Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("Transaction Id:", t.id, "## ")
	switch t.t {
	case 0:
		{
			a.deposit(t.amt)
		}
	case 1:
		{
			a.withdraw(t.amt)
		}
	}
}

func main() {
	acc := Account{
		balance: 500,
	}

	var wg sync.WaitGroup

	transactions := []Transaction{{0, 0, 24}, {1, 1, 212}, {2, 0, 14235254}, {3, 1, 24231}, {4, 1, 150020}}

	for _, t := range transactions {
		wg.Add(1)
		go acc.transact(t, &wg)
	}

	wg.Wait()

	fmt.Println("All transactions done")
}
