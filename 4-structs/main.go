package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	alex := person{
		"Alex",
		"Anderson",
		33,
		contactInfo{
			email: "AlexAnderson@Anderson.matrix",
			zip:   91210,
		},
	}
	fmt.Println(alex, alex.firstName, alex.lastName)

	tom := person{firstName: "Tom", lastName: "Fisher", age: 99}
	fmt.Println(tom, tom.firstName, tom.lastName)

	var peter person
	// peter has assigned zero value: string - "", int - 0
	fmt.Println(peter, peter.firstName, peter.lastName)
	fmt.Printf("%+v\n", peter)

	peter.firstName = "Peter"
	peter.lastName = "Parker"
	peter.age = 55

	fmt.Println(peter)

	jim := person{
		firstName: "Jim",
		lastName:  "Earth",
		age:       1,
		contact: contactInfo{
			email: "earth@jim.worm",
			zip:   94000,
		},
	}

	fmt.Printf("%+v\n", jim)

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.getData()

	(&jim).updateName("Gimmy")
	jim.getData()

	jim.updateName("NoPointerJimmy")
	jim.getData()

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	name := "Bill"
	// &name get memory address and * returns the value.
	fmt.Println(*&name)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p person) getData() {
	fmt.Println("---------")
	fmt.Println(p.firstName, p.lastName, p.age, p.contact.zip)
	fmt.Println("---------")
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}
