package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	*Address
}

type Address struct {
	Street  string
	City    string
	State   string
	Country string
	ZipCode int
}

func main() {
	a := Address{}
	p := Person{Address: &a}
	fmt.Printf("%+v", p)
}
