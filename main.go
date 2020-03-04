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
	board.NewBoard()
	for {
		for i:=0; i < 2; i++ {
			m := console.GenerateViewModelFromBoard(board)
			die = board.RollDie()
			for {
				m.Draw()
				cli.ShowTurn(board)
				cli.ShowPips(board)
				position = cli.Prompt(board, die)
				err := board.Move(position, die)
				if err == nil {
					break
				}
				cli.Clear()
			}
			cli.Clear()
		}
		board.NextTurn()
	}
}
