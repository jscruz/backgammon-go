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

// Make sure when you update a point with a lower count, the old marks are removed
func TestUpdatingSamePoint(t *testing.T) {
	m := NewModel()

	m.SetPointCount(9, 10, "X")
	m.SetPointCount(9, 2, "Y")

	// 3rd+ man on the point, it should be back to empty
	assert.Equal(t, "", m[9][2])
	assert.Equal(t, "", m[9][3])
	assert.Equal(t, "", m[9][4])

	//1st and 2nd man should have been changed
	assert.Equal(t, "Y", m[9][0])
	assert.Equal(t, "Y", m[9][1])
}