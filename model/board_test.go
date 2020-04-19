package model

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
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

	valuePosition := calculatePips(position)
	valuebarPosition := calculatePips(barPosition)
	valuehomePosition := calculatePips(homePosition)

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
	board.Board[White] = [Spaces]int{12, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	board.NextTurn()

	assert.Equal(t, true, board.IsPlayerHome(White))
	assert.Equal(t, false, board.IsPlayerHome(Red))

	board.Board[Red] = [Spaces]int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	board.Board[White] = [Spaces]int{0, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12}
	assert.Equal(t, true, board.IsPlayerHome(Red))
	assert.Equal(t, false, board.IsPlayerHome(White))
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

func TestPossibleMovesWhenAtHome(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 1, 2, 2, 5}
	board.Board[Red] = [Spaces]int{2, 3, 5, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	moves := board.GetPossibleMoves(Red, 3)

	assert.Equal(t, 3, len(moves))
	assert.True(t, contains(moves, 1))
	assert.True(t, contains(moves, 2))
	assert.True(t, contains(moves, 3))
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
	assert.Equal(t, 0, len(moves))
}

func TestAdversaryPosition(t *testing.T) {
	a := [Spaces]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	b := reverse(a)
	for i := 0; i < Spaces; i++ {
		assert.Equal(t, b[i], getAdversaryPosition(i))
	}
}

func TestContains(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6}

	assert.True(t, contains(a, 1))
	assert.False(t, contains(a, 19))
}

func TestCopy(t *testing.T) {
	a := NewBoard()
	b := NewBoard()

	b.Board[White] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 1, 2, 2, 5}
	b.Board[Red] = [Spaces]int{0, 0, 0, 0, 0, 0, 2, 1, 0, 1, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	b.Pips[White] = 33
	b.Pips[Red] = 66

	assert.Equal(t, a.Pips[White], 167)
	assert.Equal(t, a.Pips[Red], 167)
	assert.Equal(t, a.Board[White][25], 0)
	assert.Equal(t, a.Board[Red][25], 0)

	a.Copy(b)

	assert.Equal(t, b.Pips[White], 33)
	assert.Equal(t, b.Pips[Red], 66)
	assert.Equal(t, b.Board[White][25], 5)
	assert.Equal(t, b.Board[Red][25], 0)

	a.Board[White][0] = 88

	assert.Equal(t, b.Board[White][0], 0)
}

func TestDice(t *testing.T) {
	var numbers []int

	rand.Seed(1)
	numbers = RollDice()
	assert.Equal(t, 2, len(numbers))
	assert.Equal(t, 2, numbers[0])
	assert.Equal(t, 3, numbers[1])

	// Initialise the seed to force a double dice
	// I found this value iterating to check RollDice values and increasing the seed
	// Fortunately, the seed 2 was the one offering a double dice
	rand.Seed(2)
	numbers = RollDice()
	assert.Equal(t, 4, len(numbers))
	assert.Equal(t, 2, numbers[0])
	assert.Equal(t, 2, numbers[1])
	assert.Equal(t, 2, numbers[2])
	assert.Equal(t, 2, numbers[3])
}

func TestGameFinished(t *testing.T) {
	board := NewBoard()
	assert.False(t, board.GameFinished())
	board.Board[White] = [Spaces]int{15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.True(t, board.GameFinished())
}

func TestMoveToEmptyLocation(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{1, 5, 4, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	board.UpdatePips()
	assert.Equal(t, board.Pips[White], 30)
	//Try to move from a empty location
	assert.Error(t, board.Move(White, 20, 3))
}

func TestMoveAndBear(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{1, 5, 4, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	board.UpdatePips()
	//Bear a checker, reduce 1 in position 4, increase 1 in position 0 (Home Position), pips recalculated
	assert.NoError(t, board.Move(White, 4, 5))
	assert.Equal(t, board.Board[White][0], 2)
	assert.Equal(t, board.Board[White][4], 1)
	assert.Equal(t, board.Pips[White], 26)
}

func TestMoveAndHit(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 1, 2, 2, 0}
	board.Board[Red] = [Spaces]int{0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}

	//White moves to a occupied location and hit Red player
	assert.NoError(t, board.Move(White, 20, 1))
	assert.Equal(t, 1, board.Board[White][19])
	assert.Equal(t, 0, board.Board[Red][6])
	assert.Equal(t, 1, board.Board[Red][BarPosition])
}

func TestInvalidMove(t *testing.T) {
	board := NewBoard()
	board.Board[White] = [Spaces]int{1, 5, 4, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	assert.True(t, board.isValidMove(White, 4, 2))
	assert.True(t, board.isValidMove(White, 4, 5))
	assert.False(t, board.isValidMove(White, 20, 3))
	assert.False(t, board.isValidMove(White, -4, 2))
	assert.False(t, board.isValidMove(White, 4, -2))
	assert.False(t, board.isValidMove(White, 4, 8))
}

func TestPlayerHasBeenHit(t *testing.T) {
	board := NewBoard()
	assert.False(t, board.IsHit(White))
	board.Board[White] = [Spaces]int{1, 4, 4, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	assert.True(t, board.IsHit(White))
}

func TestGetCurrentPlayer(t *testing.T) {
	board := NewBoard()
	board.Turn = 1
	assert.Equal(t, Red, board.GetCurrentPlayer())

	board.Turn = 4
	assert.Equal(t, White, board.GetCurrentPlayer())

	board.Turn = 5
	assert.Equal(t, Red, board.GetCurrentPlayer())
}
