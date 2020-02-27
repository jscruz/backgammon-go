package console

import "fmt"

const (
	columns    = 27
	rows       = 21
	totalCells = columns * rows
)

type Window [rows * columns]string

func Draw(w Window) {

	for i := 0; i < rows * columns; i += columns {

		for j := i; j < i + columns; j++ {
			fmt.Printf("%v ", w[j])
		}

		fmt.Print("\n")
	}
}