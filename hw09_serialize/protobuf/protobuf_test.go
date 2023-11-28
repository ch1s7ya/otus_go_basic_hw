package protobuf

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestBook_Marshal(t *testing.T) {
	type fields struct {
		Id     int64
		Title  string
		Author string
		Year   int64
		Size   int64
		Rate   float32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Murshul protobuf book",
			fields: fields{
				Id:     1,
				Title:  "All Quiet on the Western Front",
				Author: "Erich Maria Remarque",
				Year:   1929,
				Size:   200,
				Rate:   5.0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Id:     tt.fields.Id,
				Title:  tt.fields.Title,
				Author: tt.fields.Author,
				Year:   tt.fields.Year,
				Size:   tt.fields.Size,
				Rate:   tt.fields.Rate,
			}
			got, err := b.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Book.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			book := &Book{}
			err = book.Unmarshal(got)
			if err != nil {
				t.Errorf("Book.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !proto.Equal(b, book) {
				t.Errorf("Objects isn't equal")
			}
		})
	}
}

func TestBookshelf_MarshalSlice(t *testing.T) {
	type fields struct {
		Book []*Book
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Marshal/Unmarshal slice of books",
			fields: fields{
				Book: []*Book{
					{
						Id:     1,
						Title:  "All Quiet on the Western Front",
						Author: "Erich Maria Remarque",
						Year:   1929,
						Size:   200,
						Rate:   5.0,
					},
					{
						Id:     2,
						Title:  "Three Comrades",
						Author: "Erich Maria Remarque",
						Year:   1936,
						Size:   498,
						Rate:   5.0,
					},
					{
						Id:     3,
						Title:  "Animal Farm",
						Author: "George Orwell",
						Year:   1945,
						Size:   92,
						Rate:   5.0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bookshelf{
				Book: tt.fields.Book,
			}
			got, err := b.MarshalSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bookshelf.MarshalSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			bookshelf := Bookshelf{}
			err = bookshelf.UnmarshalSlice(got)
			if err != nil {
				t.Errorf("Bookshelf.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !proto.Equal(b, &bookshelf) {
				t.Errorf("Objects isn't equal")
			}
		})
	}
}