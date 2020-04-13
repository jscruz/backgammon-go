package opponents

import (
	"backgammon_go/model"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
)

// Player interface for a player
type Player interface {
	Move(board *model.Model, die int) ([]int, error)
	GetPips() int
}

//CPU player
type CPU struct {
	pips int
}

//Human player
type Human struct {
	pips int
}

var in io.Reader = os.Stdin

//Move CPU moves randomly
func (c CPU) Move(board model.Model, die int) ([]int, error) {
	possibleMoves := board.GetPossibleMoves(board.GetCurrentPlayer(), die)
	if len(possibleMoves) > 0 {
		min := 0
		max := len(possibleMoves)
		initialPosition := possibleMoves[rand.Intn(max-min)+min]
		board.Move(board.GetCurrentPlayer(), initialPosition, die)
		return []int{initialPosition, initialPosition - die}, nil
	}
	return nil, errors.New("No moves possible")
}

//Move human
func (h Human) Move(board model.Model, die int) ([]int, error) {
	possibleMoves := board.GetPossibleMoves(board.GetCurrentPlayer(), die)
	if len(possibleMoves) > 0 {
		fmt.Print("Enter initial position: ")
		var input int
		fmt.Fscanln(in, &input)
		if contains(input, possibleMoves) {
			board.Move(board.GetCurrentPlayer(), input, die)
			return []int{input, input - die}, nil
		}
	}
	return nil, errors.New("No moves possible")
}

//GetPips get the pips calculation for CPU
func (c CPU) GetPips() int {
	return c.pips
}

//GetPips get the pips calculation for Human
func (h Human) GetPips() int {
	return h.pips
}

func contains(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
