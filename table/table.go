package table

import (
	"fmt"
	"strconv"
)

type Cell interface {
	Width() int
	Render()
}

type SimpleCell struct {
	Value string
}

func (c SimpleCell) Width() int {
	return len(c.Value)
}

func (c SimpleCell) Render() {
	fmt.Printf(c.Value)
}

type Table struct {
	Rows [][]Cell
}

func (t *Table) Row(cells []Cell) {
	t.Rows = append(t.Rows, cells)
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
			if cell.Width() > maxColumnWidths[i] {
				maxColumnWidths[i] = cell.Width()
			}
		}
	}

	for _, row := range t.Rows {
		for cellIndex, cell := range row {
			columnWidth := maxColumnWidths[cellIndex]
			extraPadding := columnWidth - cell.Width()

			if cellIndex != 0 {
				fmt.Printf(" ")
			}

			cell.Render()
			fmt.Printf("%-"+strconv.Itoa(extraPadding)+"s", "")
			fmt.Printf(" ")
		}
		fmt.Println("")
	}
}

func New() *Table {
	return &Table{}
}
