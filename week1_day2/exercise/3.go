package main

import (
	"errors"
	"fmt"
	"sync"
)

// Exercise link: https://docs.google.com/document/d/1LZ7q3zF9FmQNb2jGD78v5Az5qUtexcFxHWDCTqgtlRA/edit#heading=h.r2aa02zg243u
// Problem 3:

type Transaction struct {
	id              int32
	transactionType int32 // 0: credit, 1: debit
	amount          int32
}

type Account struct {
	mu      sync.Mutex
	balance int32
}

// Deducts the specified amount from the account balance and returns a message or an error.
// params: a pointer to an Account (a) and the withdrawal amount (amt)
// Returns a success message if the withdrawal is successful, or an error message if the balance is insufficient
func (a *Account) withdraw(amt int32) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amt > a.balance {
		return "", errors.New("Insufficient balance")
	} else {
		a.balance -= amt
		return fmt.Sprint("Debited:", amt, " Balance:", a.balance), nil
	}
}

// Adds the specified amount to the account balance and returns a message.
// params: a pointer to an Account (a) and the deposit amount (amt)
// Returns a message indicating the credit and the updated balance.
func (a *Account) deposit(amt int32) string {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amt
	return fmt.Sprint("Credited:", amt, " Balance:", a.balance)
}

// Processes a transaction and updates the account accordingly.
// It takes a Transaction (t) and a sync.WaitGroup (wg) as parameters.
// The sync.WaitGroup is used to signal when the transaction processing is complete.
func (a *Account) txn(t Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	message := fmt.Sprint("Transaction Id:", t.id, " ## ")

	switch t.transactionType {
	case 0:
		{
			response := a.deposit(t.amount)
			message += response
			fmt.Println(message)
		}
	case 1:
		{
			response, err := a.withdraw(t.amount)
			if err != nil {
				message += err.Error()
				fmt.Println(message)
			} else {
				message += response
				fmt.Println(message)
			}
		}
	}
}

func main() {
	acc := Account{
		balance: 500,
	}

	transactions := []Transaction{
		{0, 0, 24},
		{1, 1, 212},
		{2, 0, 14235254},
		{3, 1, 24231},
		{4, 1, 150020},
	}

	var wg sync.WaitGroup

	for _, t := range transactions {
		wg.Add(1)
		go acc.txn(t, &wg)
	}

	wg.Wait()

	fmt.Println("All transactions done")
}
