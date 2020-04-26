package main

import (
	"fmt"
)

// Person represents details of a person.
type Person struct {
	FirstName    string
	LastName     string
	Age          uint8
	Emails       []string
	PhoneNumbers []string
	*Address
}

// Address represents address details.
type Address struct {
	Line1   string
	Line2   string
	City    string
	PINCode string
}

func main() {
	david := Person{
		FirstName: "David",
		LastName:  "Miller",
		Age:       29,
		Emails: []string{
			"david.miller@gmail.com",
		},
		PhoneNumbers: []string{
			"9908234543",
		},
		Address: &Address{
			Line1:   "3805A, MIG",
			City:    "Hyderabad",
			PINCode: "5343232",
		},
	}

	fmt.Printf("%+v", david)
}
