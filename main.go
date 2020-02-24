package main

import (
	"backgammon_go/cli"
	"backgammon_go/model"
)

func main() {
	board := model.Board{}
	board.Setup()
	cli.Print(board)
	board.Move(model.WHITE, 24, 4)
	cli.Print(board)
	board.Move(model.WHITE, 24, 4)
	cli.Print(board)
	board.Move(model.WHITE, 20, 1)
	cli.Print(board)
}
