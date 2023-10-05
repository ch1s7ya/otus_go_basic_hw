package main

import (
	"fmt"
	"strings"
)

const (
	black = "#"
	white = " "
)

func drawBoardChess(boardSize int) string {
	var boardChess strings.Builder

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

	return boardChess.String()
}

func main() {
	boardSize := 8
	fmt.Printf("Enter size of board: ")
	fmt.Scanln(&boardSize)

	boardChess := drawBoardChess(boardSize)

	fmt.Println(boardChess)
}
