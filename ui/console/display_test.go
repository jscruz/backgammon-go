package console

import (
	"strconv"
	"testing"
)

var fullWidthLine = [columns]string{}

func TestDirtyTest(t *testing.T) {

	w := &Window{}

	for k := 0; k < totalCells; k++ {
		w[k] = " "
	}

	for i := 0; i < columns; i ++ {
		fullWidthLine[i] = "-"
	}

	topRow := 0 + columns
	midRow := columns * ((rows - 1) / 2 )
	bottomRow := columns * (rows - 2)

	repeatHorizontal(topRow, "-", w)
	repeatHorizontal(midRow, "=", w)
	repeatHorizontal(bottomRow, "-", w)

	repeatVerticleFull(12, "|", w)
	repeatVerticleFull(13, " ", w)
	repeatVerticleFull(14, "|", w)

	setPoint(1, 15, "X", w)
	setPoint(2, 4, "O", w)
	setPoint(6, 15, "X", w)

	Draw(*w)
}

func setPoint(point, count int, char string, w *Window) {
	if count > 6 {
		for i := 0; i < 5; i++ {
			c := cells[point][i]
			w[c] = char
		}

		c := cells[point][5]
		w[c] = strconv.Itoa(count)
	} else {
		for i := 0; i < count; i++ {
			c := cells[point][i]
			w[c] = char
		}
	}
}

func repeatHorizontal(row int, char string, w *Window){
	for i := 0; i < columns; i += 2 {
		w[row + i] = char
	}
}

func repeatVerticleFull(col int, char string, w *Window) {
	i := col
	for {
		if i > totalCells {
			break
		}
		w[i] = char
		i += columns
	}
}