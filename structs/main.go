package main

import "fmt"

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

	jim.updateName("jim")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
