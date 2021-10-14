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
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "example@email.com",
			zipCode: 453,
		},
	}
	alex.print()
	alexPointer := &alex
	alexPointer.updateFirstName("updatedFirstName")
	alex.updateLastName("updatedLastName")
	alex.print()

	mySlice := []string{
		"Hi",
		"There",
		"How",
		"Are",
		"You",
	}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "Bill"
	fmt.Println(*&name)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointer *person) updateFirstName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p *person) updateLastName(newLastname string) {
	p.lastName = newLastname
}

func updateSlice(s []string) {
	s[0] = "Hello"
}
