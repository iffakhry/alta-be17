package main

import "fmt"

func main() {
	// var today int = 3
	// switch today {
	// case 1:
	// 	fmt.Printf("Today is Monday")
	// case 2:
	// 	fmt.Printf("Today is Tuesday")
	// case 3:
	// 	fmt.Printf("Today is Wednesday")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	var today int = 2
	switch {
	case today == 1:
		fmt.Printf("Today is Monday")
	case today == 2:
		fmt.Printf("Today is Tuesday")
	default:
		fmt.Printf("Invalid Date")
	}

	fmt.Println()
	var day string = "senin"
	switch day {
	case "senin":
		fmt.Println("hari ini hari senin")
	case "selasa":
		fmt.Println("hari ini hari selasa")
	default:
		fmt.Println("invalid day")
	}

	value := 42
	switch value {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

}
