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

// FullName makes full name from first name and last name.
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// AddressString makes a string from Address fields.
func (addr *Address) AddressString() string {
	return addr.Line1 + " " + addr.Line2 + " " + addr.City + " " + addr.PINCode
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

	fmt.Println(david.FullName())
	fmt.Println(david.AddressString())
}
