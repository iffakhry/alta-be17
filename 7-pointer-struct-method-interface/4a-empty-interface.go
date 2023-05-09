package main

import "fmt"

//	func describe(i interface{}) {
//		fmt.Printf("(%v, %T)\n", i, i)
//	}
func describe(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	// var i interface{}
	var i any
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = true
	describe(i)
}
