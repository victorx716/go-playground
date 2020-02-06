package main

import "fmt"

// structs are pass by value (into new memory address in RAM)
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// creating one struct in another
	jim := person{
		firstName: "Jimmy",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jimmy@gmail.com",
			zipCode: 90210,
		},
	}
	// & operator retrieves memory address of jim (where the struct sits in RAM)
	// jimPointer points to original mem address

	jim.updateName("jim")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// * asks for value sitting at such mem address.
	// & is address, * is value at address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// value types: int float string bool structs
// reference types: slices maps channels pointers functions
