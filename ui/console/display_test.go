package console

import (
	"backgammon_go/model"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// Make sure when you update a point with a lower count, the old marks are removed
func TestUpdatingSamePoint(t *testing.T) {
	m := NewModel()

	m.SetPointCount(9, 10, "X")
	m.SetPointCount(9, 2, "Y")

	// 3rd+ man on the point, it should be back to empty
	assert.Equal(t, "", m.points[9][2])
	assert.Equal(t, "", m.points[9][3])
	assert.Equal(t, "", m.points[9][4])

	//1st and 2nd man should have been changed
	assert.Equal(t, "Y", m.points[9][0])
	assert.Equal(t, "Y", m.points[9][1])
}

func TestD(t *testing.T) {

	b := model.Board{}
	b.Setup()

	m := GenerateViewModelFromBoard(b)

	m.BtmPlyrName = "Dan"
	m.TopPlyrName = "Jorge"

	err := m.Draw()
	if err != nil {
		log.Fatal(err)
	}
}