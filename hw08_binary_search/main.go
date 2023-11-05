package main

import (
	"fmt"
)

func BinarySearch(s []int, searchElement int) (bool, error) {
	leftElement := 0
	rightElement := len(s) - 1
	if searchElement > s[rightElement] || searchElement < s[leftElement] {
		return false, fmt.Errorf("search element is out of range")
	}
	for leftElement <= rightElement {
		middleElement := (leftElement + rightElement) / 2
		switch {
		case searchElement < s[middleElement]:
			rightElement = middleElement
		case searchElement > s[middleElement]:
			leftElement = middleElement
		case searchElement == s[middleElement]:
			return true, nil
		}
	}
	return false, nil
}

func main() {
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6))
}
