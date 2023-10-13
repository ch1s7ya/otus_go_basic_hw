package main

import (
	"fmt"
	"strings"
)

func countWords(str string) map[string]int {
	str = strings.TrimSpace(str)
	words := strings.Split(str, " ")
	m := make(map[string]int)
	for _, word := range words {
		word = strings.Trim(word, "!?.;,-:'\"()[]{}")
		word = strings.ToLower(word)
		if word != "" {
			m[word]++
		}
	}
	return m
}

func main() {
	text := "The origins of the term software engineering have been attributed to various sources. "
	fmt.Println(countWords(text))
}
