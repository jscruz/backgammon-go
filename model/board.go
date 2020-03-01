package model

import (
	"fmt"
	"math/rand"
)

type Player int

// Players
const (
	White Player = iota
	Red
)

// Board backgammon board
type Board struct {
	Turn  Player
	Board [][]int
	Hit   [2]int
	Borne [2]int
	Pips [2]int
}

func reverse(numbers []int) []int {
	newNumbers := make([]int, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

// Setup the backgammon board with the initial position for players
func (b *Board) Setup() {
	firstPosition := []int{0, 0, 0, 0, 0, 5, 0, 3, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}
	whiteBoard := firstPosition
	redBoard := reverse(firstPosition)
	b.Board = [][]int{whiteBoard, redBoard}
	b.Turn = White
	b.GetPips()
}

// Move piece from player
func (b *Board) Move(initialPos int, moves int) error {
	var player = b.Turn
	var adversary, direction = adversary(player)
	var position = initialPos - 1
	moves = moves * direction

	if (!b.IsPlayerHome() && (position+moves > 23 || position+moves < 0)) || (!b.IsHit() && b.Board[player][position] < 1) || b.Board[adversary][position+moves] > 1 {
		return fmt.Errorf("!illegal move from %d to %d", position, moves)
	}

	if b.Board[adversary][position+moves] == 1 {
		b.Board[adversary][position+moves]--
		b.Hit[adversary]++
	}

	if b.IsPlayerHome() && (position+moves > 23 || position+moves < 0) {
		b.Borne[player]++
	} else {
		b.Board[player][position+moves]++
	} 
	
	if b.IsHit() {
		b.Hit[player]--
	} else {
		b.Board[player][position]--
	}
	b.GetPips()
	return nil
}

func (b *Board) IsHit() bool {
	return b.Hit[b.Turn] > 0
}

func (b *Board) RollDie() int {
	min := 1
	max := 6
	return rand.Intn(max-min) + min
}

func (b *Board) NextTurn() {
	b.Turn = 1 - b.Turn
}

func (player Player) HitPosition() int {
	switch player {
	case White:
		return 25
	default:
		return 0
	}
}
func adversary(player Player) (Player, int) {
	switch player {
	case White:
		return Red, -1
	default:
		return White, 1
	}
}

func (b *Board) IsPlayerHome () bool {
	checkers := []int{}
	sum := 0 
	checkers = b.Board[b.Turn][:18]
	if b.Turn == Red {
		checkers = b.Board[b.Turn][6:]
	}
	for _, num := range checkers {
        sum += num
	}
	return sum == 0
}

func (board *Board) GetPips () {
	board.Pips[White] = CalculatePips(board.Board[White]) + board.Hit[White] * 25
	board.Pips[Red] = CalculatePips(reverse(board.Board[Red])) + board.Hit[Red] * 25
}

func CalculatePips (position []int) int {
	pips := 0
	for i:=0; i<len(position); i++ {
		pips += (i+1) * position[i]
	}
	return pips
}
