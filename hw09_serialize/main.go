package main

import (
	"fmt"

	"github.com/ch1s7ya/otus_go_basic_hw/hw09_serialize/json"
	"github.com/ch1s7ya/otus_go_basic_hw/hw09_serialize/protobuf"
	"google.golang.org/protobuf/proto"
)

func main() {
	book1 := json.Book{
		ID:     1,
		Title:  "All Quiet on the Western Front",
		Author: "Erich Maria Remarque",
		Year:   1929,
		Size:   200,
		Rate:   5.0,
	}

	d, err := book1.MarshalJSON()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%s\n", d)
	}

	var book2 json.Book

	err = book2.UnmarshalJSON(d)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%#v\n", book2)
	}

	book3 := protobuf.Book{
		ID:     1,
		Title:  "All Quiet on the Western Front",
		Author: "Erich Maria Remarque",
		Year:   1929,
		Size:   200,
		Rate:   5.0,
	}

	marshalled, err := proto.Marshal(&book3)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("Marshaled:\n%s\n", marshalled)
	}

	books := []json.Book{
		{
			ID:     1,
			Title:  "All Quiet on the Western Front",
			Author: "Erich Maria Remarque",
			Year:   1929,
			Size:   200,
			Rate:   5.0,
		},
		{
			ID:     2,
			Title:  "Three Comrades",
			Author: "Erich Maria Remarque",
			Year:   1936,
			Size:   498,
			Rate:   5.0,
		},
		{
			ID:     3,
			Title:  "Animal Farm",
			Author: "George Orwell",
			Year:   1945,
			Size:   92,
			Rate:   5.0,
		},
	}

	b, err := json.MarshalSlice(books)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("Marshaled slice of book:\n%s\n", b)
	}

	// booksp := []protobuf.Book{
	// 	{
	// 		ID:     1,
	// 		Title:  "All Quiet on the Western Front",
	// 		Author: "Erich Maria Remarque",
	// 		Year:   1929,
	// 		Size:   200,
	// 		Rate:   5.0,
	// 	},
	// 	{
	// 		ID:     2,
	// 		Title:  "Three Comrades",
	// 		Author: "Erich Maria Remarque",
	// 		Year:   1936,
	// 		Size:   498,
	// 		Rate:   5.0,
	// 	},
	// 	{
	// 		ID:     3,
	// 		Title:  "Animal Farm",
	// 		Author: "George Orwell",
	// 		Year:   1945,
	// 		Size:   92,
	// 		Rate:   5.0,
	// 	},
	// }

	// bp, _ := json.MarshalSlice(&booksp)
	// fmt.Printf("Marshaled slice of book:\n%s\n", string(bp))
}
