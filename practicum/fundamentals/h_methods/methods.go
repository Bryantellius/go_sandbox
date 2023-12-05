package main

import (
	"fmt"
)

const LESSON_NUMBER int = 8

type person struct {
	name string
}

type bankAccount struct {
	accountHolder string
	balance       float64
}

func createPerson(name string) *person {
	newPerson := person{name: name}
	return &newPerson
}

func createBankAccount(accountHolderName string) *bankAccount {
	newBankAccount := bankAccount{accountHolder: accountHolderName, balance: 0}
	return &newBankAccount
}

func (p person) sayHello(name string) {
	fmt.Printf("Hello %s! My name is %s.", name, p.name)
}

func (ba *bankAccount) reportTransaction(amount float64) float64 {
	ba.balance += amount
	return ba.balance
}

func PrintWhatILearned() {
	me := createPerson("Ben")
	me.sayHello("Gandalf")

	ba := createBankAccount(me.name)
	ba.reportTransaction(100.00)
	fmt.Printf("Balance after payday: $%.2f", ba.balance)
}

func main() {
	fmt.Printf("\n%d. Methods.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
