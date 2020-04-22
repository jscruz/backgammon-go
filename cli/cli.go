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

func ShowPossibleMoves(board *model.Board, moves int) {
	fmt.Printf("Possible moves: %v\n", board.GetPossibleMoves(board.GetCurrentPlayer(), moves))
}

func ShowPips(board *model.Board) {
	fmt.Printf("Pips: White [%d] Red [%d]\n", board.Pips[model.White], board.Pips[model.Red])
}

func ShowTurn(board *model.Board) {
	fmt.Printf("Turn %v ", board.Turn)
	if board.GetCurrentPlayer() == model.White {
		fmt.Printf(" [X]\n")
	} else {
		fmt.Printf(" [O]\n")
	}
}

func Prompt(board *model.Board, die int) int {
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

// func Setup() []opponents.player {
// 	var option int
// 	for {
// 		fmt.Printf("1. human vs human\n")
// 		fmt.Printf("2. cpu vs human\n")
// 		fmt.Printf("Select one option:\n")
// 		fmt.Scanf("%d", &option)
// 		if option >=1 && option <=2 {
// 			break
// 		}
// 	}
// 	if option == 1 {
// 		return opponents.player [] {opponents.human, opponents.human}
// 	}
// 	return [] {opponents.cpu, opponents.human}
// }
