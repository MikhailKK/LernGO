package main

import "fmt"

type bankAccount struct {
	number  int
	balance int
}

func NewBankAccount(number, balance int) bankAccount {
	return bankAccount{number: number, balance: balance}
}

func (a bankAccount) IncreaseBalance(amount int) {
	a.balance += amount
}

func main() {
	ba := &bankAccount{number: 123456, balance: 100}

	fmt.Println("first", ba.balance)

	ba.IncreaseBalance(200)

	fmt.Println("second", ba.balance)

	fmt.Printf("%T\n", NewBankAccount(100, 200))

}
