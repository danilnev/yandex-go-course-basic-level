package main

import (
	"errors"
)

var (
	Balance               = 0.0
	IncorrectAmountError  = errors.New("amount is incorrect")
	IncorrectBalanceError = errors.New("balance is incorrect")
)

func topUpBalance(amount float64) error {
	if amount <= 0 {
		return IncorrectAmountError
	}
	if Balance < 0 {
		return IncorrectBalanceError
	}
	Balance += amount
	return nil
}

func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return IncorrectAmountError
	}
	if amount > Balance {
		return IncorrectBalanceError
	}
	if Balance < 0 {
		return IncorrectBalanceError
	}
	Balance -= amount
	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, IncorrectAmountError
	}
	if Balance < 0 {
		return 0.0, IncorrectBalanceError
	}
	topUpBalance(amount)
	return Balance, nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	if amount > Balance {
		return 0.0, IncorrectBalanceError
	}
	if amount <= 0 {
		return 0.0, IncorrectAmountError
	}
	if Balance < 0 {
		return 0.0, IncorrectBalanceError
	}
	chargeFromBalance(amount)
	return Balance, nil
}
