package main

import "fmt"

func main() {
	boardSize := 8
	black := "#"
	white := " "
	var boardChess string

	fmt.Printf("Enter size of board: ")
	fmt.Scanln(&boardSize)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if (i+j)%2 == 0 {
				boardChess += white
			} else {
				boardChess += black
			}
		}
		boardChess += "\n"
	}

	fmt.Println(boardChess)
}
