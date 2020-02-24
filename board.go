package main

import (
	"fmt"
)

// Players
const (
	WHITE int = iota
	RED
)

// Board backgammon board
type Board struct {
	board [][]int
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
	b.board = [][]int{whiteBoard, redBoard}
}

// Move piece from player
func (b *Board) Move(player int, initialPos int, moves int) error {
	var adversary = RED
	moves = moves * -1
	initialPos--
	if player != WHITE {
		adversary = WHITE
		moves = moves * -1
	}
	if b.board[adversary][initialPos+moves] > 1 {
		return fmt.Errorf("!illegal move from %d to %d", initialPos, moves)
	}
	b.board[player][initialPos+moves]++
	b.board[player][initialPos]--
	return nil
}
