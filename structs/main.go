package main

import "fmt"

type contactInfo struct {
	email    string
	postcode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "The Lion",
		contactInfo: contactInfo{
			email:    "alex@madagasc.ar",
			postcode: "JU12 ALX",
		}}
	// Shortcut to *person
	alex.updateName("Alexander")
	alex.print()
}

func (pn person) print() {
	fmt.Printf("%+v\n", pn)
}

func (pn *person) updateName(newFirstName string) {
	pn.firstName = newFirstName
}
