package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
)

var CONST_A, CONST_Z, CONST_S, CONST_E int

type Configuration struct {
	start     string
	end       string
	startCell *Cell
	endCell   *Cell
	matrix    *AdjacencyMatrix
}

func (me *Configuration) isNavigable(currentCell *Cell, targetCell *Cell, currentPath *utility.SimpleStack) bool {

	ok := false
	if currentPath != nil && currentPath.Contains(targetCell) {
		ok = false
	} else {
		currentHeight := 0
		targetHeight := me.letterAsInt(targetCell.z)

		if currentCell.z == "S" {
			currentHeight = me.letterAsInt("a")
		} else {
			currentHeight = me.letterAsInt(currentCell.z)
		}

		/*
			same height
		*/

		/*
			current height + 1 is target height
		*/

		/*
			target height is less than the current height
		*/

		if currentHeight == targetHeight {
			ok = true
		} else if (currentHeight + 1) == targetHeight {
			ok = true
		} else if targetHeight < currentHeight {
			if targetCell.id == "S" {
				ok = false
			} else {
				ok = true
			}

		}
	}

	return ok
}

type NodeList struct {
	nodes map[string]*Cell
}

func (me *NodeList) gather() []*Cell {
	values := make([]*Cell, 0)
	for _, aValue := range me.nodes {
		values = append(values, aValue)
	}
	return values
}

func (me *NodeList) add(cell *Cell) {
	me.nodes[cell.id] = cell
}

type AdjacencyMatrix struct {
	//cells               [][]interface{}
	//cellsCrossReference map[string]int
	nodeListMap map[string]*NodeList
}

type Node struct {
	node  *Cell
	nodes []*Node
}

func (me *Node) containsChildNode(node *Node) bool {
	itDoes := false
	for _, aChildNode := range me.nodes {
		if aChildNode.node.id == node.node.id {
			itDoes = true
			break
		}
	}
	return itDoes
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

func (me *Configuration) letterAsInt(aString string) int {
	anInt := 0
	runes := []rune(aString)
	anInt = me.runeAsInt(runes[0])
	return anInt
}

func runeAsInt(aRune rune) int {
	return int(aRune)
}

func intAsLetter(aNumber int) string {
	var aString string
	aString = string(rune(aNumber))
	return aString
}

func letterAsInt(aString string) int {
	anInt := 0
	runes := []rune(aString)
	anInt = runeAsInt(runes[0])
	return anInt
}

func (me *Configuration) runeAsInt(aRune rune) int {
	return int(aRune)
}

func (me *Configuration) intAsLetter(aNumber int) string {
	var aString string
	aString = string(rune(aNumber))
	return aString
}

func (me *Configuration) shortestBFS() *utility.SimpleStack {
	visited := utility.NewSimpleStack()
	queue := utility.NewSimpleQueue()
	shortest := me.bfs(me.startCell, me.endCell, visited, queue)

	if !queue.IsEmpty() || shortest == nil {
		panic("not a clue")
	}
	return shortest
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

type Cell struct {
	x  int
	y  int
	z  string
	id string
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

func next(sourceCell *Cell, targetCell *Cell) bool {

	ok := false

	currentHeight := 0
	targetHeight := letterAsInt(targetCell.z)

	if sourceCell.z == "S" {
		currentHeight = letterAsInt("a")
	} else {
		currentHeight = letterAsInt(sourceCell.z)
	}

	/*
		same height
	*/

	/*
		current height + 1 is target height
	*/

	/*
		target height is less than the current height
	*/

	if currentHeight == targetHeight {
		ok = true
	} else if (currentHeight + 1) == targetHeight {
		ok = true
	} else if targetHeight < currentHeight {
		if targetCell.id == "S" {
			ok = false
		} else {
			ok = true
		}

	}

	return ok
}
