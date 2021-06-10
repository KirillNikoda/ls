package models

import "log"

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID:    1,
		Title: "Lord Of The Rings",
		Author: Author{
			Name:     "J.R",
			LastName: "Tolkien",
			BornYear: 1892,
		},
		YearPublished: 1978,
	}
	log.Println(DB, book1)
	DB = append(DB, book1)
	log.Println(DB)
}

func FindBookById(id int) (*Book, bool) {
	var book Book
	var found bool

	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
		}
	}
	return &book, found

}
