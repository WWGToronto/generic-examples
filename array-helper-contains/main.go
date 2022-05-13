package main

import (
	"fmt"
	"reflect"
)

// Contains returns true if the haystack contains the needle.
func Contains[T any](haystack []T, needle T) bool {
	for _, v := range haystack {
		if reflect.DeepEqual(v, needle) {
			return true
		}
	}
	return false
}

type Card struct {
	LastFourDigits string
	Status         string
	Network        string
}

// main contains an example of how Contains may be useful for asserting something in tests.
func main() {
	cards := []Card{
		{
			LastFourDigits: "1234",
			Status:         "active",
			Network:        "visa",
		},
		{
			LastFourDigits: "9087",
			Status:         "active",
			Network:        "visa",
		},
		{
			LastFourDigits: "6644",
			Status:         "active",
			Network:        "mastercard",
		},
	}

	primaryCard := Card{ // 2nd in the list
		LastFourDigits: "9087",
		Status:         "active",
		Network:        "visa",
	}

	hasCertainCard := Contains(cards, primaryCard)
	fmt.Printf("Created contains certain card: %t\n", hasCertainCard)
}
