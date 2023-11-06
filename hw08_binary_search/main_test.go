package main

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		s             []int
		searchElement int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sorted array from 0 to 9",
			args: args{
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				7,
			},
			want: 7,
		},
		{
			name: "Search element is out of range",
			args: args{
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				42,
			},
			want: elementNotFound,
		},
		{
			name: "Boundary condition",
			args: args{
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.args.s, tt.args.searchElement)
			if got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
