package main

import (
	"encoding/json"
	"fmt"
)

type Optional[T any] struct {
	data   T
	exists bool
}

func (o Optional[T]) Get() (T, bool) {
	return o.data, o.exists
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.data)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.data); err != nil {
		return err
	}
	o.exists = true
	return nil
}

type PatchRequest struct {
	FirstName Optional[string] `json:"first_name"`
}

func main() {
	examples := map[string]string{
		"Empty JSON Object": `{}`,
		"Null":              `{ "first_name": null }`,
		"Empty":             `{ "first_name": "" }`,
		"Valid":             `{ "first_name": "Joe" }`,
		"Int":               `{ "first_name": 123 }`,
	}

	for k, v := range examples {
		fmt.Println(k)
		result := PatchRequest{}
		fmt.Printf("\tString is: %s\n", v)
		if err := json.Unmarshal([]byte(v), &result); err != nil {
			fmt.Printf("\tError: %+v\n", err)
		} else {
			value, exists := result.FirstName.Get()
			fmt.Printf("\tFirst name: %[1]q (Type: %[1]T), (Exists: %[2]t)\n", value, exists)
		}
		fmt.Println()
	}
}
