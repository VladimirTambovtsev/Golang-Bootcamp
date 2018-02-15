package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	//alex := person{firstname: "Alex", lastname: "Brown"}

	alex := person{
		firstname: "Alex",
		lastname:  "Brown",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 123456,
		},
	}

	alex.updateName("Bobby")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%v", p)
}
