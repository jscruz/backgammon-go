package console

import (
	"backgammon_go/model"
	"fmt"
	"html/template"
	"os"
	"strconv"
)

const (
	cellsPerPoint = 6 // in order to change this, just the template needs lines added or removed
	topCell = cellsPerPoint - 1 // this is the cell that might have numbers if there are too many men

	BottomBar = 26
	TopBar = 25

	BottomHome = 0
	TopHome = 27
)

type Model [28]point

type point [cellsPerPoint]string

// NewModel generates a new ui Model and instantiates is points field
func NewModel() *Model {

	m := &Model{}

	return m
}

// SetPointCount updates the model and sets the number of men on a point
// and the string that represents them. Normally "X" or "O"
//TODO: if the board model in game library is handling moving and updating the number of
//  men on certain points, then this func can be private, as its main job is to make sure strings are
//  formatted to be 2 characters wide.
func (m *Model) SetPointCount(point, count int, s string) {

	m.resetPoint(point)

	for i := 0; i < count && i < cellsPerPoint; i++ {
		m[point][i] = s
	}

	if count > cellsPerPoint {
		m[point][topCell] = strconv.Itoa(count - topCell)
	}
}

func (m *Model) resetPoint(point int) {
	for i := 0; i < cellsPerPoint; i++ {
		m[point][i] = ""
	}
}


// Draw renders the Model in the template and draws it out to the console
func (m *Model) draw() error {

	tpl, err := template.New("board").Funcs(template.FuncMap{
		"c": func(p, i int, m *Model) string {
			return fmt.Sprintf("%-2v", m[p][i])
		},
	}).Parse(outputTemplate)
	if err != nil {
		return err
	}

	err = tpl.Execute(os.Stdout, m)
	if err != nil {
		return err
	}

	return nil
}


// GenerateModelFromBoard converts a board object [][]int into a Model to be rendered
// by the template
func Draw(b model.Board) error {
	m := NewModel()

	for i, p := range b.Board[0] {
		m.SetPointCount(i + 1, p, "X")
	}

	for i, p := range b.Board[1] {
		m.SetPointCount(i + 1, p, "O")
	}

	err := m.draw()
	if err != nil {
		return err
	}

	return nil
}