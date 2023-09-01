package main

import (
	"fmt"
	"strings"
)

func main() {
	boardSize := 8
	black := "#"
	white := " "
	var boardChess strings.Builder

	fmt.Printf("Enter size of board: ")
	fmt.Scanln(&boardSize)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if (i+j)%2 == 0 {
				boardChess.WriteString(white)
			} else {
				boardChess.WriteString(black)
			}
		}
		boardChess.WriteString("\n")
	}

	fmt.Println(boardChess.String())
}
