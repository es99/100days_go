package main

import (
	"fmt"
)

func main() {
	fmt.Println(YearsBeforeDesiredBalance(200.75, 214.88))
}

func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	} else if balance >= 0 && balance < 1000 {
		return 0.5
	} else if balance >= 1000 && balance < 5000 {
		return 1.621
	} else {
		return 2.475
	}
}

func Interest(balance float64) float64 {
	return (float64(InterestRate(balance)) / 100) * balance
}

func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	anos := 0

	for balance < targetBalance {
		anos++
		balance += Interest(balance)
	}

	return anos
}
