package console

import (
	"fmt"
	"html/template"
	"os"
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

func NewModel() *Model {

	pts := [28]point{}

	m := &Model{pts}

	for pi, _ := range m.points {
		m.SetPointCount(pi,6,"")
	}

	return m
}

func (m *Model) SetPointCount(point, count int, s string) {

	for i := 0; i < count && i < 6; i++ {
		m.points[point][i] = fmt.Sprintf("%-2v", s)
	}

	if count > 6 {
		m.points[point][5] = fmt.Sprintf("%-2v", count - 5)
	}
}

func Draw(m *Model)error {

	tpl, err := template.New("board").Funcs(template.FuncMap{
		"c": func(p, i int, m *Model) string {
			return m.points[p][i]
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