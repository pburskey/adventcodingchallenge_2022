package main

import "adventcodingchallenge_2022/utility"

type Part1 struct {
	answer int
}

// x = row
// y = column
type Grid struct {
	data [][]int
}

func parse(data []string) *Grid {

	gridData := utility.ParseRowsToInts(data)
	grid := &Grid{data: gridData}
	return grid
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	grid := parse(data)
	if grid != nil {
		//alg.answer += grid.perimeter()
		for x := 0; x < len(grid.data); x++ {
			for y := 0; y < len(grid.data[x]); y++ {

				visible := grid.isPositionAtVisibleFromOutside(x, y)
				if visible {
					alg.answer++
				}
			}

		}

	}
	return nil, alg.answer
}

func (g *Grid) perimeter() int {
	outside := len(g.data)
	outside = (outside * 2) + ((outside - 2) * 2)
	return outside
}

func (g *Grid) isPositionAtVisibleFromOutside(x int, y int) bool {
	visible := false
	if x == 0 || y == 0 {
		visible = true
	} else if x == len(g.data)-1 {
		visible = true
	} else if y == len(g.data[0])-1 {
		visible = true
	} else {
		targetValue := g.data[x][y]
		above, below, left, right := g.collectElementsFromPosition(x, y)

		least, max := utility.LeastAndMax(above)
		if least < targetValue && max < targetValue {
			visible = visible || true
		}

		least, max = utility.LeastAndMax(below)
		if least < targetValue && max < targetValue {
			visible = visible || true
		}

		least, max = utility.LeastAndMax(right)
		if least < targetValue && max < targetValue {
			visible = visible || true
		}

		least, max = utility.LeastAndMax(left)
		if least < targetValue && max < targetValue {
			visible = visible || true

		}
	}

	return visible
}

func (g *Grid) collectElementsFromPosition(x int, y int) (above []int, below []int, left []int, right []int) {

	above = g.collectElementsAbovePosition(x, y)

	below = g.collectElementsBelowPosition(x, y)

	right = g.collectElementsRightOfPosition(x, y)

	left = g.collectElementsLeftOfPosition(x, y)

	return
}

func (g *Grid) collectElementsAbovePosition(x int, y int) []int {

	data := make([]int, 0)
	for xx := x - 1; xx >= 0; xx-- {
		data = append(data, g.data[xx][y])
	}

	return data

}

func (g *Grid) collectElementsBelowPosition(x int, y int) []int {

	data := make([]int, 0)
	for xx := x + 1; xx < len(g.data); xx++ {
		data = append(data, g.data[xx][y])
	}

	return data

}

func (g *Grid) collectElementsRightOfPosition(x int, y int) []int {

	data := make([]int, 0)
	for yy := y + 1; yy < len(g.data[x]); yy++ {
		data = append(data, g.data[x][yy])
	}

	return data

}

func (g *Grid) collectElementsLeftOfPosition(x int, y int) []int {

	data := make([]int, 0)
	for yy := y - 1; yy >= 0; yy-- {
		data = append(data, g.data[x][yy])
	}

	return data

}
