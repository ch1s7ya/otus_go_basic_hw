package main

import (
	"reflect"
	"testing"
)

func Test_countWords(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want map[string]int
	}{
		{
			name: "Text from one word",
			str:  "Hello",
			want: map[string]int{"hello": 1},
		},
		{
			name: "Text from two word",
			str:  "Hello Bender",
			want: map[string]int{"hello": 1, "bender": 1},
		},
		{
			name: "Text with punctuation symbols",
			str:  "Hello, Bender!",
			want: map[string]int{"hello": 1, "bender": 1},
		},
		{
			name: "Complex text",
			str:  "Hello, Bender! Bender is robot. His full name is Bender Bending Rodríguez. ",
			want: map[string]int{
				"hello": 1, "bender": 3, "is": 2,
				"robot": 1, "his": 1, "full": 1,
				"name": 1, "bending": 1, "rodríguez": 1,
			},
		},
		{
			name: "Text only with punctuation symbols",
			str:  ", !",
			want: map[string]int{},
		},
		{
			name: "Empty string",
			str:  "",
			want: map[string]int{},
		},
		{
			name: "Words with lowercase and uppercase letters",
			str:  "Bender, bender, BeNdEr",
			want: map[string]int{"bender": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWords(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
