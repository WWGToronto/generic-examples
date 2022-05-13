package main

import "fmt"

// Map converts a list of T into a list of R using the provided function.
func Map[T any, R any](input []T, mapFunc func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = mapFunc(v)
	}
	return result
}

func main() {
	input := []int{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}

	// Convert the array of integers to an array of bytes which can in-turn be converted to a string.
	output := Map(input, func(t int) byte {
		return byte(t)
	})

	fmt.Printf("%+v\n", string(output))
}
