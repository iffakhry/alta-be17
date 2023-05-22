package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// select name, height, mass, hair_color, skin_color from peoples limit 10 offset 10

type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	Films     []string `json:"films"`
}

type PeopleResponse struct {
	Count   int              `json:"count"`
	Next    string           `json:"next"`
	Results []StarWarsPeople `json:"results"`
}

func main() {
	// responses, err := http.Get("https://swapi.dev/api/people/2")
	// if err != nil {
	// 	log.Fatal("error consume API")
	// }

	// responsesBody, errBody := ioutil.ReadAll(responses.Body)
	// defer responses.Body.Close()
	// if errBody != nil {
	// 	log.Fatal("error read response body")
	// }

	// var people1 StarWarsPeople
	// json.Unmarshal(responsesBody, &people1)

	// fmt.Println(people1.Name)
	// fmt.Println(people1.Height)
	// fmt.Println(people1.HairColor)
	// // fmt.Println(people1.Films)
	// for _, value := range people1.Films {
	// 	fmt.Println("people film:", value)
	// }

	// MEMBACA JSON OBJECT NESTED
	responses, err := http.Get("https://swapi.dev/api/people/?page=9")
	if err != nil {
		log.Fatal("error consume API")
	}

	responsesBody, errBody := ioutil.ReadAll(responses.Body)
	defer responses.Body.Close()
	if errBody != nil {
		log.Fatal("error read response body")
	}

	var people PeopleResponse
	// mengubah dari format json ke object
	json.Unmarshal(responsesBody, &people)
	fmt.Println(people.Count)
	fmt.Println(people.Next)
	fmt.Println(people.Results)
	// for _, value := range people.Results {
	// 	fmt.Println("name people:", value.Name)
	// }
	for i := 0; i < len(people.Results); i++ {
		fmt.Println("name people:", people.Results[i].Name)

	}

}
