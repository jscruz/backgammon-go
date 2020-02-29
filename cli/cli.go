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

func ShowPips(board model.Board) {
	fmt.Printf("Pips: White [%d] Red [%d]\n", board.Pips[model.White], board.Pips[model.Red])
}

func Prompt(board model.Board, die int) int {
	var position int
	fmt.Printf("Moves: %d >> ", die)
	if board.IsHit() {
		fmt.Printf("You have been hit. Moving to recover")
		position = board.Turn.HitPosition()
	} else {
		fmt.Printf("Move from [1-24]:")
		fmt.Scanf("%d", &position)
	}
	return position
}
