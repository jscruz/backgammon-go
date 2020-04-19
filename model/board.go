package model

import (
	"errors"
	"math/rand"
)

// Player it could be White: 0 or Red: 1
type Player int

// Players
const (
	White Player = iota
	Red
	Spaces       = 26
	BarPosition  = Spaces - 1
	HomePosition = 0
	Checkers     = 15
)

// Model interface for the board model
type Model interface {
	NewBoard()
	Copy(*Model)
	Move(player Player, initialPos int, moves int) (bool, error)
	GameFinished() bool
	GetPossibleMoves(player Player, moves int) []int
	NextTurn() Player
	GetCurrentPlayer() Player
	IsPlayerHome(player Player) bool
	IsHit(player Player) bool
}

// Board backgammon board
type Board struct {
	Turn  int
	Board [2][Spaces]int
	Pips  [2]int
}

// Copy function to do a deep copy of the board game
func (b *Board) Copy(c *Board) {
	b = c
	b.Board = c.Board
	b.Pips = c.Pips
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

// NewBoard setups the backgammon board with the initial position for players
func NewBoard() *Board {
	b := &Board{}

	var redPosition = [Spaces]int{0, 0, 0, 0, 0, 0, 5, 0, 3, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0}
	var whitePosition = [Spaces]int{0, 0, 0, 0, 0, 0, 5, 0, 3, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0}
	b.Board = [2][Spaces]int{whitePosition, redPosition}
	b.Turn = 0
	b.UpdatePips()
	return b
}

// Move changes the position of checkers in the board given a initial position initialPos and moves we want to take.
// It will keep it in the Home Position if the number of moves exceeds that position.
// It checks as well if as an action of moving the other player is hit
func (b *Board) Move(player Player, initialPos int, moves int) error {
	if b.isValidMove(player, initialPos, moves) {
		b.Board[player][initialPos]--
		if initialPos-moves <= HomePosition {
			b.Board[player][HomePosition]++
		} else {
			b.Board[player][initialPos-moves]++
			opponent := adversary(player)
			opponentPosition := getAdversaryPosition(initialPos - moves)
			if b.Board[opponent][opponentPosition] == 1 {
				b.Board[opponent][opponentPosition]--
				b.Board[opponent][BarPosition]++
			}
		}
		b.UpdatePips()
		return nil
	}
	return errors.New("Illegal move")
}

func (b *Board) isValidMove(player Player, initialPos int, moves int) bool {
	if moves > 0 && moves < 7 && initialPos > 0 && initialPos < BarPosition && b.Board[player][initialPos] > 0 {
		return true
	}
	return false
}

//GameFinished returns true if any of the players has borne 15 checkers
func (b *Board) GameFinished() bool {
	return b.Board[White][0] == Checkers || b.Board[Red][0] == Checkers
}

//GetPossibleMoves returns a list of possible initial position to move the number of player moves
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

	for i := 1; i < Spaces; i++ {
		//Check player is home and moves are equal or minor than HomePosition
		if b.IsPlayerHome(player) && (i-moves) <= HomePosition && currentPlayerBoard[i] > 0 {
			possibleMoves = append(possibleMoves, i)
		}
		if (i-moves) > HomePosition && currentPlayerBoard[i] > 0 && adversaryBoard[i-moves] < 2 {
			possibleMoves = append(possibleMoves, i)
		}
	}
	return possibleMoves
}

//RollDie returns a random integer value within 1 to 6
func RollDie() int {
	min := 1
	max := 6
	return rand.Intn(max-min) + min
}

//RollDice returns a pair of numbers representing dice in the game.
func RollDice() []int {
	var dice []int
	die1 := RollDie()
	die2 := RollDie()
	if die1 == die2 {
		dice = []int{die1, die2, die1, die2}
	} else {
		dice = []int{die1, die2}
	}
	return dice
}

//NextTurn returns which player is next and add a turn to the internal counter
func (b *Board) NextTurn() Player {
	b.Turn++
	return Player(b.Turn % 2)
}

//GetCurrentPlayer return the current player
func (b *Board) GetCurrentPlayer() Player {
	return Player(b.Turn % 2)
}

//IsPlayerHome checks is a given player is at home board given that there is no checkers in the rest of the board.
func (b *Board) IsPlayerHome(player Player) bool {
	checkers := []int{}
	sum := 0
	checkers = b.Board[player][7:]
	for _, num := range checkers {
		sum += num
	}
	return sum == 0
}

//IsHit checks is a player has any checker in the bar position
func (b *Board) IsHit(player Player) bool {
	return b.Board[player][BarPosition] > 0
}

//UpdatePips calculate pips and assign them to each player pips
func (b *Board) UpdatePips() {
	b.Pips[White] = calculatePips(b.Board[White])
	b.Pips[Red] = calculatePips(b.Board[Red])
}

//calculatePips calculates pips based on the player positions
func calculatePips(position [Spaces]int) int {
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

func adversary(player Player) Player {
	switch player {
	case White:
		return Red
	default:
		return White
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
