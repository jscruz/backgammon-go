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
}

// Move piece from player
func (b *Board) Move(initialPos int, moves int) error {
	var player = b.Turn
	var adversary, direction = adversary(player)
	var position = initialPos - 1
	moves = moves * direction

	if (!b.IsHit() && b.Board[player][position] < 1) || b.Board[adversary][position+moves] > 1 {
		return fmt.Errorf("!illegal move from %d to %d", position, moves)
	}

	if b.Board[adversary][position+moves] == 1 {
		b.Board[adversary][position+moves]--
		b.Hit[adversary]++
	}

	b.Board[player][position+moves]++
	if b.IsHit() {
		b.Hit[player]--
	} else {
		b.Board[player][position]--
	}
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
	switch b.Turn {
	case White:
		b.Turn = Red
	default:
		b.Turn = White
	}
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
