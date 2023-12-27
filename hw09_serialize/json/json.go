package json

import "encoding/json"

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

func MarshalSlice(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func UnmarshalSlice(b []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(b, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
