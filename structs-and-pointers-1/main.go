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

	alex := person{"Alex", "Anderson", contactInfo{"alex@go.lang", 94141}}

	frank := person{firstName: "Frank", lastName: "Franklin", contactInfo: contactInfo{email: "frank@go.lang", zipCode: 55555}}

	joe := person{
		firstName: "Joe",
		lastName:  "Josephson",
		contactInfo: contactInfo{
			email:   "joe@go.lang",
			zipCode: 10110,
		},
	}

	fmt.Printf("%+v\n", alex)
	fmt.Printf("%+v\n", joe)
	fmt.Printf("%+v\n", frank)

	alex.firstName = "Alexis"
	alex.lastName = "Andreeeeeson"
	alex.contactInfo.email = "alexis@go.lang"
	alex.contactInfo.zipCode = 20000

	fmt.Printf("%+v\n", alex)

	joePointer := &joe
	joePointer.updateName("fakeName")
	alex.print()
	joe.print()
	frank.print()
	fmt.Println(joePointer)

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
