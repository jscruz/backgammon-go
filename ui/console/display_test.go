package console

import (
	"backgammon_go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirtyTest(t *testing.T) {

	b := model.Board{}
	b.Setup()


	m := NewModel()

	m.SetPointCount(9, 15, "X")
	m.SetPointCount(9, 2, "Y")

	//for i, p := range b.Board[0] {
	//	m.SetPointCount(i + 1, p, "X")
	//}
	//
	//for i, p := range b.Board[1] {
	//	m.SetPointCount(i + 1, p, "O")
	//}

 	err := m.draw()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdatingSamePoint(t *testing.T) {
	m := NewModel()

	//m.SetPointCount(9, 10, "X")
	m.SetPointCount(13, 20, "Y")
	m.SetPointCount(BottomHome, 25, "X")

	assert.Equal(t, "", m.points[9][2])
	assert.Equal(t, "Y", m.points[9][1])

	m.draw()
}