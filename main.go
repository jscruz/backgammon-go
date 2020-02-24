package main

import "fmt"

func main() {
	var board = new(Board)
	board.setup()
	fmt.Printf(board.toString())
	board.move(WHITE, 24, 4)
	fmt.Printf(board.toString())
	board.move(WHITE, 24, 4)
	fmt.Printf(board.toString())
	board.move(WHITE, 20, 1)
	fmt.Printf(board.toString())
}
