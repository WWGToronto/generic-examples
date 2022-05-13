package main

import (
	"fmt"
	"sort"
)

// TypedSort wraps the standard library sort but has compile-time type checking.
func TypedSort[T any](input []T, less func(T, T) bool) {
	sort.Slice(input, func(i, j int) bool {
		return less(input[i], input[j])
	})
}

type myObj struct {
	Name string
	Age  int
}

func main() {
	data := []myObj{
		{
			Name: "Joe",
			Age:  25,
		},
		{
			Name: "Sue",
			Age:  30,
		},
		{
			Name: "May",
			Age:  20,
		},
		{
			Name: "Bob",
			Age:  40,
		},
	}

	for _, v := range data {
		fmt.Printf("(Before) %s is %d\n", v.Name, v.Age)
	}

	// Sort the above array by age.
	TypedSort(data, func(a, b myObj) bool {
		return a.Age < b.Age
	})

	fmt.Println()
	for _, v := range data {
		fmt.Printf("(After) %s is %d\n", v.Name, v.Age)
	}
}
