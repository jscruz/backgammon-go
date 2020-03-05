package cli

import (
	"backgammon_go/model"
	"fmt"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ShowBoard(board model.Board) {
	fmt.Printf("%v\n%v\n", board.Board[0], board.Board[1])
}

func ShowPossibleMoves(board model.Board, moves int) {
	fmt.Printf("Possible moves: %v\n", board.GetPossibleMoves(board.GetCurrentPlayer(), moves))
}

func ShowPips(board model.Board) {
	fmt.Printf("Pips: White [%d] Red [%d]\n", board.Pips[model.White], board.Pips[model.Red])
}

func Prompt(board model.Board, die int) int {
	var position int
	fmt.Printf("Moves: %d >> ", die)
	if board.IsHit(board.GetCurrentPlayer()) {
		fmt.Printf("You have been hit. Moving to recover")
		position = model.BarPosition
	} else {
		fmt.Printf("Select checker [1-24]:")
		fmt.Scanf("%d", &position)
	}
	return position
}
