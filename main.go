package main

func main() {
	var board = new(Board)
	board.Setup()
	PrintBoard(board)
	board.Move(WHITE, 24, 4)
	PrintBoard(board)
	board.Move(WHITE, 24, 4)
	PrintBoard(board)
	board.Move(WHITE, 20, 1)
	PrintBoard(board)
}
