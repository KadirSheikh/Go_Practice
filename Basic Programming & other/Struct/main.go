package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	kadir := person{

		firstName: "Kadir",
		lastName:  "Sheikh",
		contact: contactInfo{
			email:   "kadir@gmail.com",
			zipCode: 12345,
		},
	}

	kadirPointer := &kadir
	kadirPointer.updateFirstName("Kamil")
	kadir.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (personPointer *person) updateFirstName(newFirstName string) {
	(personPointer).firstName = newFirstName
}
