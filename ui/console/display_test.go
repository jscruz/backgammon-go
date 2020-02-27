package console_test

import (
	"backgammon_go/model"
	"backgammon_go/ui/console"
	"testing"
)

func TestDirtyTest(t *testing.T) {

	b := model.Board{}
	b.Setup()


	m := console.NewModel()

	//m.SetPointCount(9, 15, "X")

	for i, p := range b.Board[0] {
		m.SetPointCount(i + 1, p, "X")
	}

	for i, p := range b.Board[1] {
		m.SetPointCount(i + 1, p, "O")
	}

 	err := console.Draw(m)
	if err != nil {
		t.Fatal(err)
	}
}