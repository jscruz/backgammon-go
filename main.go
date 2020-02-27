package main

import (
	"backgammon_go/cli"
	"backgammon_go/model"
	"backgammon_go/ui/console"
)

func main() {
	board := model.Board{}
	var die int 
	var position int
	cli.Clear()
	board.Setup()
	for {
		//cli.Print(board)

		console.GenerateModelFromBoard(board)

		die = board.RollDie()
		for {
			position = cli.Prompt(board, die)
			err := board.Move(position, die)
			if err == nil {
				break;
			}
		}
		board.NextTurn()
		cli.Clear()
	}
}
