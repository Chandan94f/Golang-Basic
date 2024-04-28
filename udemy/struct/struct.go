package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstName string
	lastName  string

	// contact contactInfo
	contactInfo
}

func fStruct() {
	avenger := person{
		firstName: "Spider",
		lastName:  "Man",
		//contact: contactInfo{
		contactInfo: contactInfo{
			email:   "spider@marvel.com",
			pincode: 560025,
		},
	}

	// avenger.updateName("Iron")

	// one way
	// avengerPointer := &avenger

	// avengerPointer.updateName("Iron")

	//other way -- even this work as it is sending as memomary address also,

	// all depend upon receiver's end how it takes as value or as reference.

	avenger.updateName("Iron")

	avenger.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (p person) print() {
	fmt.Println(p)
}

// every single has to be a comma
/*
func print() {

	alex1 := person{"Harry", "Potter"}

	alex2 := person{firstName: "Harry", lastName: "Potter"}

	var alex3 person

	fmt.Println(alex1)

	fmt.Println(alex2)

	fmt.Println(alex3)

	fmt.Printf("%+v\n", alex3)

	alex3.firstName = "Tim"
	alex3.lastName = "Cook"

	fmt.Println(alex3)
}
*/
