package main

import "fmt"

type EmployeeStruct struct {
	name   string
	salary int
}

func (e *EmployeeStruct) changeName(newName string) {
	(*e).name = newName
}

func main() {
	e := EmployeeStruct{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)
	// create pointer to `e`
	ep := &e
	// change name
	ep.changeName("Monica Geller")
	// e after name change
	fmt.Println("e after name change =", e)
}
