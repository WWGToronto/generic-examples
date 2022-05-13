package main

import "fmt"

type Money int64

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Add[T integer](m Money, v T) Money {
	fmt.Printf("calling Add(%d, %d) with %T\n", m, v, *new(T))
	return Money(uint64(m) + uint64(v))
}

func Mul[T integer](m Money, v T) Money {
	fmt.Printf("calling Mul(%d, %d) with %T\n", m, v, *new(T))
	return Money(uint64(m) * uint64(v))
}

// main contains an example of how to allow multiple data types to be passed to a function with casting handled internally.
func main() {
	aValue := 5
	aOtherValue := int32(8)
	aMoney := Money(800)

	aMoney = Add(aMoney, 1)
	aMoney = Add(aMoney, aOtherValue)
	aMoney = Add(aMoney, Money(1))
	aMoney = Add(aMoney, aValue)
	aMoney = Mul(aMoney, aValue)

	fmt.Printf("Result: %v\n", aMoney)
}
