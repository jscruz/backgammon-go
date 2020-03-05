package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoardSetup(t *testing.T) {
	board := NewBoard()

	assert.Equal(t, 0, board.Turn)
	assert.Equal(t, 167, board.Pips[White])
	assert.Equal(t, 167, board.Pips[Red])
}

func TestPipsCalculation(t *testing.T) {
	position := [Spaces]int{0, 0, 0, 0, 0, 0, 5, 0, 3, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0}
	barPosition := [Spaces]int{}
	barPosition[BarPosition] = 1
	homePosition := [Spaces]int{}
	homePosition[HomePosition] = 1

	valuePosition := CalculatePips(position)
	valuebarPosition := CalculatePips(barPosition)
	valuehomePosition := CalculatePips(homePosition)

	assert.Equal(t, 167, valuePosition)
	assert.Equal(t, 25, valuebarPosition)
	assert.Equal(t, 0, valuehomePosition)
}

func TestNextTurn(t *testing.T) {
	board := NewBoard()
	player := board.NextTurn()
	assert.Equal(t, Red, player)
	assert.Equal(t, 1, board.Turn)

	player = board.NextTurn()
	assert.Equal(t, White, player)
	assert.Equal(t, 2, board.Turn)

	player = board.NextTurn()
	assert.Equal(t, Red, player)
	assert.Equal(t, 3, board.Turn)
}

func TestPlayerIsHome(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 14}
	board.NextTurn()

	assert.Equal(t, true, board.IsPlayerHome(White))
	assert.Equal(t, false, board.IsPlayerHome(Red))
}

func TestPossibleMoves(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 1, 2, 2, 5}
	board.Board[Red] = [Spaces]int{0, 0, 0, 0, 0, 0, 2, 3, 0, 1, 0, 0, 2, 0, 0, 0, 0, 0, 1, 1, 0, 5, 0, 0, 0, 0}

	moves := board.GetPossibleMoves(Red, 3)

	assert.Equal(t, 6, len(moves))
	assert.True(t, contains(moves, 6))
	assert.True(t, contains(moves, 9))
	assert.True(t, contains(moves, 12))
	assert.True(t, contains(moves, 18))
	assert.True(t, contains(moves, 19))
	assert.True(t, contains(moves, 21))
}

func TestPossibleMovesWhenInBar(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 1, 2, 2, 5}
	board.Board[Red] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}

	moves := board.GetPossibleMoves(White, 6)

	assert.Equal(t, 1, len(moves))
	assert.True(t, contains(moves, 25))
}

func TestNoPossibleMovesWhenInBar(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 1, 2, 2, 5}
	board.Board[Red] = [Spaces]int{0, 0, 0, 0, 0, 0, 2, 1, 0, 1, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	moves := board.GetPossibleMoves(White, 6)
	fmt.Printf("%v", moves)
	assert.Equal(t, 0, len(moves))
}
