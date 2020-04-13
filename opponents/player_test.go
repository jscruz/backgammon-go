package opponents

import (
	"backgammon_go/model"
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BoardStub struct {
	model.Model
	possibleMoves   []int
	initialPosition int
}

func (b *BoardStub) GetCurrentPlayer() model.Player {
	return model.White
}
func (b *BoardStub) GetPossibleMoves(p model.Player, die int) []int {
	return b.possibleMoves
}
func (b *BoardStub) Move(p model.Player, initialPosition int, die int) (bool, error) {
	b.initialPosition = initialPosition
	return true, nil
}

func TestHumanOpponentMove(t *testing.T) {
	bak := in
	in = bytes.NewReader([]byte("6\n"))
	defer func() { in = bak }()

	b := &BoardStub{}
	b.possibleMoves = []int{3, 6, 9}
	var h Human = Human{}
	move, error := h.Move(b, 3)
	assert.Equal(t, 6, move[0])
	assert.Equal(t, 3, move[1])
	assert.Nil(t, error)
}
func TestHumanOpponentEmptyMove(t *testing.T) {
	bak := in
	//This will make the position to 0
	in = bytes.NewReader([]byte("\n"))
	defer func() { in = bak }()

	b := &BoardStub{}
	b.possibleMoves = []int{3, 6, 9}
	var h Human = Human{}
	move, error := h.Move(b, 3)
	assert.NotNil(t, error)
	assert.Nil(t, move)
}

func TestHumanOpponentWrongCharMove(t *testing.T) {
	bak := in
	in = bytes.NewReader([]byte("c\n"))
	defer func() { in = bak }()

	b := &BoardStub{}
	b.possibleMoves = []int{3, 6, 9}
	var h Human = Human{}
	move, error := h.Move(b, 3)
	assert.NotNil(t, error)
	assert.Nil(t, move)
}

func TestCPUMove(t *testing.T) {
	rand.Seed(1)

	b := &BoardStub{}
	b.possibleMoves = []int{3, 6, 9}
	var cpu CPU = CPU{}
	move, error := cpu.Move(b, 3)
	assert.Equal(t, 9, move[0])
	assert.Equal(t, 6, move[1])
	assert.Nil(t, error)
}

func TestHumanOpponentNoPossibleMoves(t *testing.T) {
	bak := in
	in = bytes.NewReader([]byte("5\n"))
	defer func() { in = bak }()

	b := &BoardStub{}
	b.possibleMoves = []int{}
	var h Human = Human{}
	move, error := h.Move(b, 3)
	assert.NotNil(t, error)
	assert.Nil(t, move)
}

func TestCPUOpponentNoPossibleMoves(t *testing.T) {
	rand.Seed(1)

	b := &BoardStub{}
	b.possibleMoves = []int{}
	var cpu CPU = CPU{}
	move, error := cpu.Move(b, 3)
	assert.NotNil(t, error)
	assert.Nil(t, move)
}

func TestPips(t *testing.T) {
	var cpu CPU = CPU{}
	var human Human = Human{}

	cpu.pips = 3
	human.pips = 6

	assert.Equal(t, cpu.pips, cpu.GetPips())
	assert.Equal(t, human.pips, human.GetPips())
}
