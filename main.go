package main

import (
	"backgammon_go/cli"
	"backgammon_go/model"
)

func main() {
	board := model.Board{}
	var die int 
	var position int
	cli.Clear()
	board.Setup()
	for {
		cli.Print(board)
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
