package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. len string
	sentence := "hello guys"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)

	// 2. compare string
	str1 := "abc"
	str2 := "abc"
	fmt.Println(str1 == str2)

	// 3. Contains
	var str = "something"
	var substr = "ethi"

	res := strings.Contains(str, substr)
	fmt.Println(res) // true

	// 4. substring
	value := "cat;dog"
	// Take substring from index 4 to length of string.
	// substring := value[4:len(value)]
	// substring := value[2:5]
	// substring := value[:3]
	substring := value[2:]
	fmt.Println(substring)

	// 5. Replace
	s := "this[things]I would li[ke to re[move"
	// t := strings.Replace(s, "[", "-", -1)
	text := strings.Replace(s, "[", "-", 1)
	text = strings.Replace(text, "]", "*", 1)
	fmt.Printf("%s\n", text)

	// 6. Insert
	p := "green"
	index := 2
	q := p[:index] + "HI" + p[index:]
	fmt.Println(p, q)

	var data = "1234"
	var key = data[0] // ascii
	fmt.Println("key:", key)
	var huruf = string(data[0])
	fmt.Println("huruf:", huruf)
}
