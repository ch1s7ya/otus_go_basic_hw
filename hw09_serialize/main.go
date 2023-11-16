package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float32 `json:"rate"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(*b)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type tmpBook Book
	var book tmpBook
	err := json.Unmarshal(data, &book)
	if err != nil {
		return err
	}
	*b = Book(book)
	return nil
}

func main() {
	book1 := Book{
		ID:     1,
		Title:  "All Quiet on the Western Front",
		Author: "Erich Maria Remarque",
		Year:   1929,
		Size:   200,
		Rate:   5.0,
	}

	d, _ := book1.MarshalJSON()
	fmt.Printf("%s\n", string(d))

	var book2 Book

	err := book2.UnmarshalJSON(d)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%#v\n", book2)
}
