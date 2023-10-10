package main

import (
	"fmt"
)

var CONST_A, CONST_Z, CONST_S, CONST_E int

type Configuration struct {
	start     int
	end       int
	startCell *Cell
	endCell   *Cell
}

type NodeList struct {
	nodes map[string]*Cell
}

func (me *NodeList) add(cell *Cell) {
	me.nodes[cell.id] = cell
}

type AdjacencyMatrix struct {
	//cells               [][]interface{}
	//cellsCrossReference map[string]int
	nodeListMap map[string]*NodeList
}

func (me *AdjacencyMatrix) label(y, x int) string {
	label := fmt.Sprintf("y%dx%d", y, x)
	return label
}

func (me *AdjacencyMatrix) from(grid *Grid) {
	if grid != nil {
		//area := grid.area()
		//me.cells = make([][]interface{}, area)
		//for y, _ := range me.cells {
		//	me.cells[y] = make([]interface{}, area)
		//}

		me.nodeListMap = make(map[string]*NodeList)

		for y, _ := range grid.cells {
			for _, aCell := range grid.cells[y] {
				label := aCell.id
				me.nodeListMap[label] = &NodeList{nodes: make(map[string]*Cell)}
			}
		}

		for y, _ := range grid.cells {
			for _, sourceCell := range grid.cells[y] {
				//sourceCell := grid.cells[y][x]
				neighbors := sourceCell.Neighbors(grid)
				for _, aNeighborCell := range neighbors {
					if aNeighborCell != nil {
						/*
							adjacency dictionary
						*/
						if aNodeList, ok := me.nodeListMap[sourceCell.id]; ok {
							aNodeList.add(aNeighborCell)
						}
					}

				}
			}
		}

	}
}

func (me *Configuration) LoadStartAndEnd(aGrid *Grid) {
	for y, _ := range aGrid.cells {
		for _, aCell := range aGrid.cells[y] {
			if aCell.z == me.start {
				me.startCell = aCell
			} else if aCell.z == me.end {
				me.endCell = aCell
			}
		}
	}
}

func (me *Configuration) nextHeightFrom(aCell *Cell) int {
	targetHeight := 0
	if aCell.z == me.letterAsInt('S') {
		targetHeight = me.letterAsInt('a')
	} else {
		targetHeight = aCell.z + 1
	}
	return targetHeight
}

func (me *Configuration) letterAsInt(aRune rune) int {
	return int(aRune - '0')
}

type Grid struct {
	cells [][]*Cell
}

func (me *Grid) area() int {
	var area int
	if me.cells != nil {
		if 0 < len(me.cells) && 0 < len(me.cells[0]) {
			area = len(me.cells) * len(me.cells[0])
		}
	}
	return area
}

type CompleteGraph struct {
	cells [][]*CellPathFrom
}

func (me *CompleteGraph) load(grid *Grid) {

}

type CellPathFrom struct {
	cell              *Cell
	nextNeighbors     []*Cell
	previousNeighbors []*Cell
}

type Cell struct {
	x         int
	y         int
	z         int
	id        string
	neighbors []*Cell
}

func cellAtYAndXOrNil(y, x int, cells [][]*Cell) *Cell {
	var cell *Cell
	if y > -1 && y < len(cells) && x > -1 && x < len(cells[y]) {
		cell = cells[y][x]
	}
	return cell
}

func (me *Cell) Neighbors(aGrid *Grid) []*Cell {
	neighbors := make([]*Cell, 4)

	var x, y int
	if x == 0 && y == 0 {
		//wtf
	}
	/*
		left
	*/
	y = me.y
	x = me.x - 1

	neighbors[0] = cellAtYAndXOrNil(y, x, aGrid.cells)

	/*
		up
	*/
	y = me.y - 1
	x = me.x
	neighbors[1] = cellAtYAndXOrNil(y, x, aGrid.cells)

	/*
		right
	*/
	y = me.y
	x = me.x + 1
	neighbors[2] = cellAtYAndXOrNil(y, x, aGrid.cells)

	/*
		down
	*/
	y = me.y + 1
	x = me.x
	neighbors[3] = cellAtYAndXOrNil(y, x, aGrid.cells)

	return neighbors
}
