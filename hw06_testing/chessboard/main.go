package main

import "fmt"

func main() {
	var size int
	fmt.Print("Input length side chessboard: ")
	fmt.Scanf("%d", &size)
	fmt.Println(CreateChessboard(size))
}

func CreateChessboard(size int) string {
	var board string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
				board += fmt.Sprint(" ")
			} else if i%2 == 0 && j%2 != 0 || i%2 != 0 && j%2 == 0 {
				board += fmt.Sprint("#")
			}
		}
		board += fmt.Sprintln("")
	}
	return board
}
