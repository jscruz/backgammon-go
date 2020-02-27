package console

import (
	"backgammon_go/model"
	"fmt"
	"html/template"
	"os"
	"strconv"
)

const (
	BottomBar = 26
	TopBar = 25

	BottomHome = 0
	TopHome = 27
)

type Model struct{
	points [28]point
}
type point [6]string

// NewModel generates a new ui Model and instantiates is points field
func NewModel() *Model {

	pts := [28]point{}

	m := &Model{pts}

	return m
}

// SetPointCount updates the model and sets the number of men on a point
// and the string that represents them. Normally "X" or "O"
//TODO: if the board model in game library is handling moving and updating the number of
//  men on certain points, then this func can be private, as its main job is to make sure strings are
//  formatted to be 2 characters wide.
func (m *Model) SetPointCount(point, count int, s string) {

	for i := 0; i < count && i < 6; i++ {
		m.points[point][i] = s
	}

	if count > 6 {
		m.points[point][5] = strconv.Itoa(count - 5)
	}
}


// Draw renders the Model in the template and draws it out to the console
func Draw(m *Model) error {

	tpl, err := template.New("board").Funcs(template.FuncMap{
		"c": func(p, i int, m *Model) string {
			return fmt.Sprintf("%-2v", m.points[p][i])
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
func GenerateModelFromBoard(b model.Board) error {
	m := NewModel()

	for i, p := range b.Board[0] {
		m.SetPointCount(i + 1, p, "X")
	}

	for i, p := range b.Board[1] {
		m.SetPointCount(i + 1, p, "O")
	}

	err := Draw(m)
	if err != nil {
		return err
	}

	return nil
}