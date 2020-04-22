package main

import (
	"backgammon_go/cli"
	"backgammon_go/model"
	"backgammon_go/ui/console"
)

func main() {
	board := model.NewBoard()
	cli.Clear()
	for {
		dice := model.RollDice()
		for i := 0; i < len(dice); i++ {
			for {
				m := console.GenerateViewModelFromBoard(board)
				m.Draw()
				cli.ShowPips(board)
				cli.ShowPossibleMoves(board, dice[i])
				cli.ShowTurn(board)
				position := cli.Prompt(board, dice[i])
				err := board.Move(board.GetCurrentPlayer(), position, dice[i])
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
