package utility

import (
	"fmt"
	"strconv"
)

type Coordinate struct {
	X     int
	Y     int
	Value interface{}
}

func (c *Coordinate) DistanceBetween(another *Coordinate) *Coordinate {

	xDifference := AbsDiffInt(c.X, another.X)
	yDifference := AbsDiffInt(c.Y, another.Y)

	return &Coordinate{X: xDifference, Y: yDifference}
}

func (c *Coordinate) IsAdjacent(another *Coordinate) bool {
	adjacent := false
	xDifference := AbsDiffInt(c.X, another.X)
	yDifference := AbsDiffInt(c.Y, another.Y)

	adjacent = xDifference <= 1 && yDifference <= 1

	return adjacent
}

func (c *Coordinate) IsInSamePositionAs(other *Coordinate) bool {
	same := false
	if other != nil {
		same = c.X == other.X && c.Y == other.Y
	}
	return same
}

func (c *Coordinate) IsInOneOfPositions(otherPositions []*Coordinate) bool {
	same := false
	if otherPositions != nil {
		for _, position := range otherPositions {
			if position != nil && position.IsInSamePositionAs(c) {
				return true
			}
		}
	}
	return same
}

func (c *Coordinate) IsInSameXAs(other *Coordinate) bool {
	same := false
	if other != nil {
		same = c.X == other.X
	}
	return same
}

func (c *Coordinate) IsInSameYAs(other *Coordinate) bool {
	same := false
	if other != nil {
		same = c.Y == other.Y
	}
	return same
}

func (c *Coordinate) MoveTo(targetCoordinate *Coordinate) {
	if targetCoordinate != nil {
		c.X = targetCoordinate.X
		c.Y = targetCoordinate.Y
	}

}

func Boundaries(coordinates []*Coordinate) (minX int, minY int, maxX int, maxY int) {
	for _, aCoordinate := range coordinates {
		if aCoordinate.X > maxX {
			maxX = aCoordinate.X
		}
		if aCoordinate.X < minX {
			minX = aCoordinate.X
		}

		if aCoordinate.Y > maxY {
			maxY = aCoordinate.Y
		}
		if aCoordinate.Y < minY {
			minY = aCoordinate.Y
		}

	}
	return
}

type CoordinateRange struct {
	Start *Coordinate
	End   *Coordinate
}

func (c *CoordinateRange) FilterStartAndEnd(coordinates []*Coordinate) []*Coordinate {
	filteredCoordinates := make([]*Coordinate, 0)
	if coordinates != nil {
		aSlice := []*Coordinate{c.Start, c.End}
		for _, aCoordinate := range coordinates {
			if !aCoordinate.IsInOneOfPositions(aSlice) {
				filteredCoordinates = append(filteredCoordinates, aCoordinate)
			}
		}
	}
	return filteredCoordinates
}

func (c *CoordinateRange) DetermineCoordinatesInRange(includeDiag bool) []*Coordinate {
	coordinates := make([]*Coordinate, 0)
	if c.Start.X == c.End.X {
		x := c.Start.X
		numbers := []int{c.Start.Y, c.End.Y}
		least, max := LeastAndMax(numbers)
		for y := least; y <= max; y++ {
			aCoordinate := &Coordinate{
				X: x,
				Y: y,
			}
			coordinates = append(coordinates, aCoordinate)
		}

	} else if c.Start.Y == c.End.Y {

		y := c.Start.Y
		//diff := int(math.Abs(float64(c.Start.x - c.end.x)))
		numbers := []int{c.Start.X, c.End.X}
		least, max := LeastAndMax(numbers)
		for x := least; x <= max; x++ {
			aCoordinate := &Coordinate{
				X: x,
				Y: y,
			}
			coordinates = append(coordinates, aCoordinate)
		}

	} else {
		if includeDiag {
			xNumbers := NumbersBetween(c.Start.X, c.End.X)
			OrderNumbersStartingWithAndEndingWith(xNumbers, c.Start.X, c.End.X)
			yNumbers := NumbersBetween(c.Start.Y, c.End.Y)
			OrderNumbersStartingWithAndEndingWith(yNumbers, c.Start.Y, c.End.Y)

			for j := 0; j < len(xNumbers); j++ {
				jx := j
				jy := j
				if jx >= len(xNumbers) {
					jx = len(xNumbers) - 1
				}
				if jy >= len(yNumbers) {
					jy = len(yNumbers) - 1
				}
				aCoordinate := &Coordinate{
					X: xNumbers[jx],
					Y: yNumbers[jy],
				}
				coordinates = append(coordinates, aCoordinate)
			}
		}
	}

	return coordinates
}

func SimplePositionsRelativeTo(xstart, ystart int, endOfY int, endOfX int) []*Coordinate {
	coordinates := make([]*Coordinate, 0)
	for x := (xstart - 1); x <= (xstart + 1); x++ {

		if x != xstart {
			if x >= 0 && x < endOfX {
				coordinate := &Coordinate{
					X: x,
					Y: ystart,
				}
				coordinates = append(coordinates, coordinate)
			}
		}

	}

	for y := (ystart - 1); y <= (ystart + 1); y++ {

		if y != ystart {
			if y >= 0 && y < endOfY {
				coordinate := &Coordinate{
					X: xstart,
					Y: y,
				}
				coordinates = append(coordinates, coordinate)
			}
		}

	}

	return coordinates
}

func PositionsRelativeTo(ystart, xstart int, endOfY int, endOfX int) []*Coordinate {
	coordinates := make([]*Coordinate, 0)
	for y := (ystart - 1); y <= (ystart + 1); y++ {
		for x := (xstart - 1); x <= (xstart + 1); x++ {

			if y < endOfY && x < endOfX {
				if x >= 0 && y >= 0 {

					skip := (x == xstart && y == ystart)
					if !skip {
						coordinate := &Coordinate{
							X: x,
							Y: y,
						}
						coordinates = append(coordinates, coordinate)
					}

				}
			}

		}
	}
	return coordinates
}

func PrettyPrintInts(data [][]int) {
	for y, _ := range data {
		for x, _ := range data[y] {
			fmt.Printf("%d", data[y][x])
		}
		fmt.Printf("\n")
	}
}

// x = row
// y = column
type Grid struct {
	data [][]int
}

func (g *Grid) SetData(someData [][]int) *Grid {
	g.data = someData
	return g
}

func (g *Grid) GetData() [][]int {
	return g.data
}

func (g *Grid) Perimeter() int {
	outside := len(g.data)
	outside = (outside * 2) + ((outside - 2) * 2)
	return outside
}

func (g *Grid) IsPositionAtVisibleFromOutside(x int, y int) bool {
	visible := false
	if x == 0 || y == 0 {
		visible = true
	} else if x == len(g.data)-1 {
		visible = true
	} else if y == len(g.data[0])-1 {
		visible = true
	} else {
		targetValue := g.data[x][y]
		above, below, left, right := g.CollectElementsFromPosition(x, y)

		least, max := LeastAndMax(above)
		if least < targetValue && max < targetValue {
			visible = visible || true
		}

		least, max = LeastAndMax(below)
		if least < targetValue && max < targetValue {
			visible = visible || true
		}

		least, max = LeastAndMax(right)
		if least < targetValue && max < targetValue {
			visible = visible || true
		}

		least, max = LeastAndMax(left)
		if least < targetValue && max < targetValue {
			visible = visible || true

		}
	}

	return visible
}

func (g *Grid) CollectElementsFromPosition(x int, y int) (above []int, below []int, left []int, right []int) {

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

func PrettyPrint(coordinates []*Coordinate) {

	minX, minY, maxX, maxY := Boundaries(coordinates)

	xAdjustment := 0
	if minX < 0 {
		xAdjustment = AbsInt(minX)
	}

	yAdjustment := 0
	if minY < 0 {
		yAdjustment = AbsInt(minY)
	}

	data := make([][]string, yAdjustment+AbsInt(maxY)+2)
	for x, _ := range data {
		data[x] = make([]string, xAdjustment+AbsInt(maxX)+2)
	}

	for _, aCoordinate := range coordinates {
		if aCoordinate.Value == nil {
			aCoordinate.Value = ""
		}
		newX := aCoordinate.X + xAdjustment
		newY := aCoordinate.Y + yAdjustment

		value, err := strconv.Atoi(data[newY][newX])
		if err == nil {

			coordinateValue, _ := strconv.Atoi(fmt.Sprintf("%v", aCoordinate.Value))

			if value < coordinateValue {

			} else {
				data[newY][newX] = fmt.Sprintf("%s", aCoordinate.Value)
			}

		} else {
			data[newY][newX] = fmt.Sprintf("%s", aCoordinate.Value)
		}

	}

	for i := range data {
		i = len(data) - 1 - i
		row := data[i]
		for _, column := range row {
			if column == "" {
				column = "."
			}
			fmt.Printf("%s", column)
		}
		fmt.Printf("\n")
	}

}
