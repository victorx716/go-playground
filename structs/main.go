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
	fmt.Printf("%+v", jim)
}
