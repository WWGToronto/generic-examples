package main

import (
	"fmt"
)

// TypedSort returns a sorted version of the input using a hastily-written bubble sort.
func TypedSort[T any](input []T, less func(T, T) bool) []T {
	res := make([]T, len(input))
	for k, v := range input {
		res[k] = v
	}

	for i := 0; i < len(res)-1; i++ {
		for j := i + 1; j < len(res); j++ {
			if less(res[j], res[i]) {
				tmp := res[j]
				res[j] = res[i]
				res[i] = tmp
			}
		}
	}

	return res
}

type myObj struct {
	Name string
	Age  int
}

// main contains an example of using a typed sort method that returns a copy of the list.
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
	sorted := TypedSort(data, func(a, b myObj) bool {
		return a.Age < b.Age
	})

	fmt.Println()
	for _, v := range data {
		fmt.Printf("(Original Not Sorted) %s is %d\n", v.Name, v.Age)
	}

	fmt.Println()
	for _, v := range sorted {
		fmt.Printf("(Sorted) %s is %d\n", v.Name, v.Age)
	}
}
