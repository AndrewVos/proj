package table

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

type Table struct {
	Rows    [][]string
	Colours [][]*color.Color
}

func (t *Table) Row(cells []string) {
	t.Rows = append(t.Rows, cells)
}

func (t *Table) ColouriseRow(colours []*color.Color) {
	t.Colours = append(t.Colours, colours)
}

func (t *Table) colourForCell(row int, cell int) *color.Color {
	if len(t.Colours) > row {
		if len(t.Colours[row]) > cell {
			return t.Colours[row][cell]
		}
	}

	return color.New()
}

func (t *Table) Print() {
	maxCellCount := 0
	for _, row := range t.Rows {
		if len(row) > maxCellCount {
			maxCellCount = len(row)
		}
	}

	maxColumnWidths := make([]int, maxCellCount)

	for _, row := range t.Rows {
		for i, cell := range row {
			if len(cell) > maxColumnWidths[i] {
				maxColumnWidths[i] = len(cell)
			}
		}
	}

	for rowIndex, row := range t.Rows {
		for cellIndex, cell := range row {
			columnWidth := maxColumnWidths[cellIndex]

			colour := t.colourForCell(rowIndex, cellIndex)

			if cellIndex != 0 {
				fmt.Printf(" ")
			}

			colour.Printf("%-"+strconv.Itoa(columnWidth)+"s", cell)

			fmt.Printf(" ")
		}
		fmt.Println("")
	}
}

func New() *Table {
	return &Table{}
}
