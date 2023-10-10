package main

import (
	"fmt"
	"strings"
)

func parseCommands(data []string, grid *Grid) *Grid {

	parsedData := make([][]string, 0)

	for _, aRow := range data {
		characters := strings.Split(aRow, "")
		parsedData = append(parsedData, characters)

	}

	if parsedData != nil {
		grid.cells = make([][]*Cell, 0)
		for y, ydata := range parsedData {
			grid.cells = append(grid.cells, make([]*Cell, 0))
			for x, xdata := range ydata {

				runes := []rune(xdata)
				height := int(runes[0] - '0')
				//guid := uuid.New()
				aCell := &Cell{
					x:  x,
					y:  y,
					z:  height,
					id: fmt.Sprintf("y%dx%d", y, x),
				}
				grid.cells[y] = append(grid.cells[y], aCell)
			}
		}
	}

	for y, _ := range grid.cells {
		for x, aCell := range grid.cells[y] {
			if aCell.y != y || aCell.x != x {
				panic("Danger will robinson, danger")
			}
		}
	}

	return grid
}
