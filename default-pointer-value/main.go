package main

import "fmt"

func Ref[T any](v T) *T {
	return &v
}

// OrDefault returns the de-referenced value of the first parameter, or a default value if the first parameter is nil.
func OrDefault[T any](v *T, or T) T {
	if v == nil {
		return or
	}
	return *v
}

type Profile struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func main() {
	profiles := map[string]Profile{
		"NotNull": {
			FirstName:  "Mary",
			MiddleName: Ref("Abcde"),
			LastName:   "Sue",
		},
		"Null": {
			FirstName: "Mary",
			LastName:  "Sue",
		},
	}

	for k, v := range profiles {
		fmt.Println(k)
		fmt.Printf("\tFirst:  %s\n", v.FirstName)
		fmt.Printf("\tMiddle: %s\n", OrDefault(v.MiddleName, "<null>"))
		fmt.Printf("\tLast:   %s\n", v.LastName)
	}
}
