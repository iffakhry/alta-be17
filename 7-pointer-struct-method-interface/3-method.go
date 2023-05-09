package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

type Student struct {
	FirstName string
	LastName  string
}

type x string

// function
func fullNameFunc(firstname, lastname string) string {
	full := firstname + " " + lastname
	return full
}

// method
func (emp Employee) fullName(gelar string) string {
	full := emp.FirstName + " " + emp.LastName + " " + gelar
	return full
}

func (s x) CetakKata() string {
	return string(s)
}

func main() {
	employee1 := Employee{
		FirstName: "Bambang",
		LastName:  "Pamungkas",
	}

	student1 := Student{
		FirstName: "Bambang",
		LastName:  "Pamungkas",
	}
	// fmt.Println("hasil func:", fullNameFunc(employee1.FirstName, employee1.LastName))
	fmt.Println("hasil method:", employee1.fullName("ST"))
	// fmt.Println("hasil method:", student1.fullName())
	fmt.Println("hasil func:", fullNameFunc(student1.FirstName, student1.LastName))
}
