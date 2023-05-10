package main

import "fmt"

/*
5! = 5 * 4 * 3 * 2 * 1

	= 5 * 4!

4! = 4 * 3 * 2 * 1

	= 4 * 3!

2! = 2 * 1
1!  = 1
*/
func factorial(n int) int {
	if n == 1 { //base case
		return 1
	} else { //recurrent relations
		return n * factorial(n-1)
	}
}

func factorialLoop(n int) int {
	var hasil int = 1
	for i := n; i >= 1; i-- {
		// hasil *= i
		hasil = hasil * i
	}
	return hasil
}

func main() {
	// result := factorial(5)
	result := factorialLoop(5)
	fmt.Println("result", result)
}
