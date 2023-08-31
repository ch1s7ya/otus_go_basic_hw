package main

import "fmt"

type Mode int

const (
	Year Mode = iota
	Size
	Rate
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

type ComparisonMode struct {
	mode Mode
}

func (b *Book) NewBook(id int, title string, author string, year int, size int, rate float64) {
	b.id = id
	b.title = title
	b.author = author
	b.year = year
	b.size = size
	b.rate = rate
}

func (b Book) GetBook() Book {
	return b
}

func (cm *ComparisonMode) SetComparisonMode(mode Mode) {
	cm.mode = mode
}

func (cm *ComparisonMode) CompareBooks(book1, book2 Book) bool {
	var ComparisonResult bool
	switch {
	case cm.mode == 0:
		ComparisonResult = book1.year >= book2.year
	case cm.mode == 1:
		ComparisonResult = book1.size >= book2.size
	case cm.mode == 2:
		ComparisonResult = book1.rate >= book2.rate
	}
	return ComparisonResult
}

func main() {
	var book1, book2 Book
	book1.NewBook(1, "1984", "George Orwell", 1949, 328, 5.0)
	book2.NewBook(2, "Three Comrades", "Erich Maria Remarque", 1937, 498, 5.0)

	fmt.Printf("First book is: %v\nSecond book is: %v\n", book1.GetBook(), book2.GetBook())

	var cm ComparisonMode
	cm.SetComparisonMode(Year)

	fmt.Println("Is year of book1 more year of book2: ", cm.CompareBooks(book1, book2)) // True

	cm.SetComparisonMode(Size)
	fmt.Println("Is size of book1 more size of book2: ", cm.CompareBooks(book1, book2)) // False

	cm.SetComparisonMode(Rate)
	fmt.Println("Is rate of book1 more rate of book2: ", cm.CompareBooks(book1, book2)) // True
}
