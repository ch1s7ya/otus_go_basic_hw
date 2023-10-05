package main

import "testing"

func TestComparator_CompareBooks(t *testing.T) {
	type fields struct {
		mode ComparatorMode
	}
	type args struct {
		book1 Book
		book2 Book
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Compare two books by year",
			fields: fields{mode: Year},
			args: args{
				book1: *NewBook(1, "1984", "George Orwell", 1949, 328, 5.0),
				book2: *NewBook(2, "Three Comrades", "Erich Maria Remarque", 1937, 498, 5.0),
			},
			want: true,
		},
		{
			name:   "Compare two books by size",
			fields: fields{mode: Size},
			args: args{
				book1: *NewBook(1, "1984", "George Orwell", 1949, 328, 5.0),
				book2: *NewBook(2, "Three Comrades", "Erich Maria Remarque", 1937, 498, 5.0),
			},
			want: false,
		},
		{
			name:   "Compare two books by rate",
			fields: fields{mode: Rate},
			args: args{
				book1: *NewBook(1, "1984", "George Orwell", 1949, 328, 5.0),
				book2: *NewBook(2, "Three Comrades", "Erich Maria Remarque", 1937, 498, 5.0),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Comparator{
				mode: tt.fields.mode,
			}
			if got := c.CompareBooks(tt.args.book1, tt.args.book2); got != tt.want {
				t.Errorf("Comparator.CompareBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
