package main

import "fmt"

type Author struct {
	Nama     string
	PenName  string
	Umur     uint
	IsActive bool
}

type Book struct {
	Title       string
	PublishYear uint
	Publisher   string
	Authors     []Author
}

func main() {
	// var tere = Author{
	// 	Nama:     "Tere liye",
	// 	PenName:  "Tere liye",
	// 	Umur:     30,
	// 	IsActive: true,
	// }

	var book1 Book
	book1.Title = "Bumi Matahari"
	book1.PublishYear = 2020
	book1.Publisher = "Gramedia"
	// book1.Authors = tere
	book1.Authors = append(book1.Authors, Author{
		Nama:     "Tere liye",
		PenName:  "Tere liye",
		Umur:     30,
		IsActive: true,
	})

	fmt.Println("book:", book1)
	fmt.Println("book title:", book1.Title)
	fmt.Println("book publish year:", book1.PublishYear)
	fmt.Println("book author name:", book1.Authors[0].Nama)

	// var books []Books

}
