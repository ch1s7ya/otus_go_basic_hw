package json

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook_MarshalJSON(t *testing.T) {
	type fields struct {
		ID     int
		Title  string
		Author string
		Year   int
		Size   int
		Rate   float32
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Marshal book",
			fields: fields{
				ID:     1,
				Title:  "All Quiet on the Western Front",
				Author: "Erich Maria Remarque",
				Year:   1929,
				Size:   200,
				Rate:   5.0,
			},
			want:    []byte(`{"id":1,"title":"All Quiet on the Western Front","author":"Erich Maria Remarque","year":1929,"size":200,"rate":5}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				ID:     tt.fields.ID,
				Title:  tt.fields.Title,
				Author: tt.fields.Author,
				Year:   tt.fields.Year,
				Size:   tt.fields.Size,
				Rate:   tt.fields.Rate,
			}
			got, err := b.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Book.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Book.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_UnmarshalJSON(t *testing.T) {
	type fields struct {
		ID     int
		Title  string
		Author string
		Year   int
		Size   int
		Rate   float32
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Unmarshal book",
			fields: fields{
				ID:     1,
				Title:  "All Quiet on the Western Front",
				Author: "Erich Maria Remarque",
				Year:   1929,
				Size:   200,
				Rate:   5.0,
			},
			args: args{
				data: []byte(`{"id":1,"title":"All Quiet on the Western Front","author":"Erich Maria Remarque","year":1929,"size":200,"rate":5}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				ID:     tt.fields.ID,
				Title:  tt.fields.Title,
				Author: tt.fields.Author,
				Year:   tt.fields.Year,
				Size:   tt.fields.Size,
				Rate:   tt.fields.Rate,
			}
			if err := b.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Book.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, "Erich Maria Remarque", b.Author)
		})
	}
}

func TestMarshalSlice(t *testing.T) {
	type args struct {
		books []Book
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Marshal slice of book",
			args: args{
				books: []Book{
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
				},
			},
			want:    []byte(`[{"id":1,"title":"All Quiet on the Western Front","author":"Erich Maria Remarque","year":1929,"size":200,"rate":5},{"id":2,"title":"Three Comrades","author":"Erich Maria Remarque","year":1936,"size":498,"rate":5},{"id":3,"title":"Animal Farm","author":"George Orwell","year":1945,"size":92,"rate":5}]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalSlice(tt.args.books)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalSlice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []Book
		wantErr bool
	}{
		{
			name: "Unmarshal slice of book",
			args: args{
				b: []byte(`[{"id":1,"title":"All Quiet on the Western Front","author":"Erich Maria Remarque","year":1929,"size":200,"rate":5},{"id":2,"title":"Three Comrades","author":"Erich Maria Remarque","year":1936,"size":498,"rate":5},{"id":3,"title":"Animal Farm","author":"George Orwell","year":1945,"size":92,"rate":5}]`),
			},
			want: []Book{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalSlice(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalSlice() = %v, want %v", got, tt.want)
			}
			assert.Equal(t, "Erich Maria Remarque", got[0].Author)
		})
	}
}
