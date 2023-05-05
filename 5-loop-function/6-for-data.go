package main

import "fmt"

func main() {
	var dataslice = []string{"adi", "budi", "cindy", "dodi", "edi"}
	dataslice = append(dataslice, "frodi")
	for i := 0; i < len(dataslice); i++ {
		fmt.Println("user:", dataslice[i])
	}

	var usermap = map[string]string{}
	usermap["adi"] = "jogja"
	usermap["budi"] = "malang"
	usermap["cindy"] = "surabaya"
	// for i := 0; i < len(usermap); i++ {
	// 	fmt.Println(usermap[i])
	// }

	for keyMap, valueMap := range usermap {
		fmt.Printf("nama: %s, lahir: %s\n", keyMap, valueMap)
		// fmt.Println("nama:", keyMap, ", lahir:", valueMap)
	}
	// nama: adi, lahir: jogja

}
