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
	// alex := person{"Alex", "Anderson"}

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Print(alex)

	// another way
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// 40. Embedding Struct
	steven := person{
		firstName: "Steven",
		lastName:  "Theodore",
		contact: contactInfo{
			email:   "steventheodore98@gmail.com",
			zipCode: 40114,
		},
	}

	// 41. Receiver function
	// steven.print()
	// steven.updateName("SteveHehe")
	// steven.print()

	//42. Pass by value

	// 43. Pointer function
	// stevenPointer := &steven
	// stevenPointer.updateName("Steve")
	// stevenPointer.print()

	// 44. Pointer shortcut
	steven.updateName("Sbeve")
	steven.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
