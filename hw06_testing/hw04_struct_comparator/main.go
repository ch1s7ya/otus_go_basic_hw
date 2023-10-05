package main

import "fmt"

type ComparatorMode int

const (
	Year ComparatorMode = iota
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

type Comparator struct {
	mode ComparatorMode
}

func NewBook(id int, title string, author string, year int, size int, rate float64) *Book {
	return &Book{id, title, author, year, size, rate}
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float64 {
	return b.rate
}

func NewComparator(mode ComparatorMode) *Comparator {
	return &Comparator{mode: mode}
}

func (c *Comparator) SetComparisonMode(mode ComparatorMode) {
	c.mode = mode
}

func (c *Comparator) CompareBooks(book1, book2 Book) bool {
	var ComparisonResult bool
	switch {
	case c.mode == 0:
		ComparisonResult = book1.year >= book2.year
	case c.mode == 1:
		ComparisonResult = book1.size >= book2.size
	case c.mode == 2:
		ComparisonResult = book1.rate >= book2.rate
	}
	return ComparisonResult
}

func main() {
	book1 := NewBook(1, "1984", "George Orwell", 1949, 328, 5.0)
	book2 := NewBook(2, "Three Comrades", "Erich Maria Remarque", 1937, 498, 5.0)

	fmt.Printf("First book is: %v\nSecond book is: %v\n", book1, book2)

	c := NewComparator(Year)

	fmt.Println("Is year of book1 more year of book2: ", c.CompareBooks(*book1, *book2)) // True

	c.SetComparisonMode(Size)
	fmt.Println("Is size of book1 more size of book2: ", c.CompareBooks(*book1, *book2)) // False

	c.SetComparisonMode(Rate)
	fmt.Println("Is rate of book1 more rate of book2: ", c.CompareBooks(*book1, *book2)) // True
}
