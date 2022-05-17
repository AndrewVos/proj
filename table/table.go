package table

import (
	"fmt"
	"golang.org/x/term"
)

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignRight
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
	Rows          [][]Cell
	cellAlignment map[int]Alignment
	stretchIndex  int
}

func New() *Table {
	return &Table{stretchIndex: -1, cellAlignment: map[int]Alignment{}}
}

func (t *Table) Row(cells []Cell) {
	t.Rows = append(t.Rows, cells)
}

func (t *Table) SetCellStretch(index int) {
	t.stretchIndex = index
}

func (t *Table) SetCellAlignment(index int, alignment Alignment) {
	t.cellAlignment[index] = alignment
}

func terminalWidth() int {
	width, _, err := term.GetSize(0)

	if err != nil {
		defaultWidth := 100
		return defaultWidth
	}

	return width
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
		stretch := 0
		if t.stretchIndex >= 0 {
			fullRowWidth := 0
			for cellIndex, cell := range row {
				columnWidth := maxColumnWidths[cellIndex]
				extraPadding := columnWidth - cell.Width()
				fullRowWidth += extraPadding
				fullRowWidth += cell.Width()

				if cellIndex > 0 {
					fullRowWidth += 1
				}
			}
			stretch = terminalWidth() - fullRowWidth
		}

		for cellIndex, cell := range row {
			columnWidth := maxColumnWidths[cellIndex]
			extraPadding := columnWidth - cell.Width()

			if cellIndex > 0 {
				fmt.Printf(" ")
			}
			alignment := t.cellAlignment[cellIndex]

			if alignment == AlignRight {
				fmt.Printf(padding(extraPadding, " "))
				if t.stretchIndex == cellIndex {
					fmt.Print(padding(stretch, " "))
				}
			}

			cell.Render()

			if alignment == AlignLeft {
				fmt.Printf(padding(extraPadding, " "))
				if t.stretchIndex == cellIndex {
					fmt.Print(padding(stretch, " "))
				}
			}

		}
		fmt.Println("")
	}
}

func padding(width int, str string) string {
	result := ""
	for i := 0; i < width; i++ {
		result = result + str
	}
	return result
}
