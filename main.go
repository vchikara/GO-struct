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

	jim := person{firstName: "vineet",
		lastName: "chikara",
		contact: contactInfo{
			email:   "vineet.chikara@gmail.com",
			zipCode: 120989,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("ram")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
