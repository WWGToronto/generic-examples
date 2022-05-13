package main

import "fmt"

// Ref returns a reference to the value passed.
func Ref[T any](v T) *T {
	return &v
}

// DeRef returns a de-referenced value, or the empty value of the type if nil.
func DeRef[T any](v *T) T {
	if v == nil {
		return *new(T)
	}
	return *v
}

type Profile struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func main() {
	p := Profile{
		FirstName:  "Joe",
		MiddleName: Ref("A."),
		LastName:   "Blow",
	}

	fmt.Printf("Profile: %+v\n", p)
	fmt.Printf("Middle name (ver 1): %[1]s (Type: %[1]T)\n", p.MiddleName)
	fmt.Printf("Middle name (ver 2): %[1]s (Type: %[1]T)\n", DeRef(p.MiddleName))
}
