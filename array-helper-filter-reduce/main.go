package main

import (
	"fmt"
)

// Filter returns a subset of the provided array based on criteria
func Filter[T any](in []T, fn func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range in {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

// Reduce converts a list of type T into a value of type R.
func Reduce[T any, R any](input []T, reduceFunc func(T, R) R) R {
	result := *new(R)
	for _, v := range input {
		result = reduceFunc(v, result)
	}
	return result
}

type Account struct {
	Type    string
	Balance int
}

// main contains an example of how Filter and Reduce could help with working with data.
func main() {
	data := []Account{
		{
			Type:    "chequing",
			Balance: 10803,
		},
		{
			Type:    "savings",
			Balance: 54,
		},
		{
			Type:    "chequing",
			Balance: 23467,
		},
		{
			Type:    "chequing",
			Balance: 765,
		},
		{
			Type:    "savings",
			Balance: 463416,
		},
	}

	chequingAccounts := Filter(data, func(a Account) bool {
		return a.Type == "chequing"
	})

	balance := Reduce(chequingAccounts, func(a Account, b int) int {
		return a.Balance + b
	})

	fmt.Printf("The balance of all chequing accounts: %d\n", balance)
}
