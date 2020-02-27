package console_test

import (
	"backgammon_go/ui/console"
	"testing"
)

func TestDirtyTest(t *testing.T) {

	m := console.NewModel()
	m.SetPointCount(12, 4, "x")

	m.SetPointCount(console.TopBar, 1, "x")
	m.SetPointCount(console.BottomHome, 20, "x")
	m.SetPointCount(console.TopHome, 10, "o")

 	err := console.Draw(m)
	if err != nil {
		t.Fatal(err)
	}
}