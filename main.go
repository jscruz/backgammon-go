package main

import (
	"backgammon_go/cli"
	"backgammon_go/model"
	"backgammon_go/ui/console"
)

func main() {
	board := model.NewBoard()
	var die int
	var position int
	cli.Clear()
	for {
		for i := 0; i < 2; i++ {
			m := console.GenerateViewModelFromBoard(board)
			die = model.RollDie()
			for {
				m.Draw()
				cli.ShowPips(board)
				cli.ShowPossibleMoves(board, die)
				position = cli.Prompt(board, die)
				allowedMove, err := board.Move(board.GetCurrentPlayer(), position, die)
				if allowedMove || err != nil {
					break
				}
				cli.Clear()
			}
			cli.Clear()
		}
		board.NextTurn()
	}
}
