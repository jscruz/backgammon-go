package model

import (
	"errors"
	"math/rand"
)

type Player int

// Players
const (
	White Player = iota
	Red

	Spaces = 26

	BarPosition = Spaces - 1

	HomePosition = 0
)

// Board backgammon board
type Board struct {
	Turn  int
	Board [Spaces][Spaces]int
	Hit   [2]int
	Borne [2]int
	Pips  [2]int
}

func reverse(numbers [Spaces]int) [Spaces]int {
	var newNumbers [Spaces]int
	j := 0
	for i := Spaces - 1; i >= 0; i-- {
		newNumbers[j] = numbers[i]
		j++
	}
	return newNumbers
}

// Setup the backgammon board with the initial position for players
func NewBoard() *Board {
	b := Board{}

	var redPosition = [Spaces]int{0, 0, 0, 0, 0, 0, 5, 0, 3, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0}
	var whitePosition = [Spaces]int{0, 0, 0, 0, 0, 0, 5, 0, 3, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0}
	b.Board = [Spaces][Spaces]int{whitePosition, redPosition}
	b.Turn = 0
	b.GetPips()
	return &b
}

func (b *Board) Move(player Player, initialPos int, moves int) (bool, error) {
	possibleMoves := b.GetPossibleMoves(player, moves)
	if len(possibleMoves) < 1 {
		return false, errors.New("No moves possible")
	}
	adversaryBoard := reverse(b.Board[adversary(player)])
	hasMoved := false
	if contains(possibleMoves, initialPos) {
		b.Board[player][initialPos]--
		if initialPos-moves <= HomePosition {
			b.Board[player][HomePosition]++
		} else {
			b.Board[player][initialPos-moves]++
			if adversaryBoard[getAdversaryPosition(initialPos)] == 1 {
				b.Board[adversary(player)][getAdversaryPosition(initialPos)] = 0
				b.Board[adversary(player)][BarPosition]++
			}
		}
		hasMoved = true
	}
	b.GetPips()
	return hasMoved, nil
}

func (b *Board) GetPossibleMoves(player Player, moves int) []int {
	currentPlayerBoard := b.Board[player]
	adversaryBoard := reverse(b.Board[adversary(player)])
	var possibleMoves []int

	if currentPlayerBoard[BarPosition] > 0 {
		if adversaryBoard[BarPosition-moves] < 2 {
			possibleMoves = append(possibleMoves, BarPosition)
		}
		return possibleMoves
	}

	for i := 0; i < Spaces; i++ {
		if i-moves < HomePosition && b.IsPlayerHome(player) && i > HomePosition {
			possibleMoves = append(possibleMoves, i)
		}
		if i-moves > HomePosition && currentPlayerBoard[i] > 0 && adversaryBoard[i-moves] < 2 {
			possibleMoves = append(possibleMoves, i)
		}
	}
	return possibleMoves
}

func (b *Board) RollDie() int {
	min := 1
	max := 6
	return rand.Intn(max-min) + min
}

func (b *Board) NextTurn() Player {
	b.Turn++
	return Player(b.Turn % 2)
}

func (b *Board) GetCurrentPlayer() Player {
	return Player(b.Turn % 2)
}

func adversary(player Player) Player {
	switch player {
	case White:
		return Red
	default:
		return White
	}
}

func (b *Board) IsPlayerHome(player Player) bool {
	checkers := []int{}
	sum := 0
	checkers = b.Board[player][:18]
	for _, num := range checkers {
		sum += num
	}
	return sum == 0
}

func (b *Board) IsHit(player Player) bool {
	return b.Board[player][BarPosition] > 0
}

func (board *Board) GetPips() {
	board.Pips[White] = CalculatePips(board.Board[White])
	board.Pips[Red] = CalculatePips(board.Board[Red])
}

func CalculatePips(position [Spaces]int) int {
	pips := 0
	for i := 0; i < len(position); i++ {
		pips += i * position[i]
	}
	return pips
}

func getAdversaryPosition(position int) int {
	adversary := []int{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	return adversary[position]
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
