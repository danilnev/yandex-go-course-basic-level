package main

import (
	"errors"
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(o string) *Account {
	return &Account{balance: 0, owner: o}
}

func (a *Account) SetBalance(value float64) error {
	if value >= 0 {
		a.balance = value
		return nil
	}
	return errors.New("Balance must be positive")
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(value float64) error {
	if value >= 0 {
		a.balance += value
		return nil
	}
	return errors.New("Deposit amount must be positive")
}

func (a *Account) Withdraw(value float64) error {
	if value >= 0 && value <= a.balance {
		a.balance -= value
		return nil
	}
	return errors.New("Withdrawal amount must be positive and less than or equal to the balance")
}
