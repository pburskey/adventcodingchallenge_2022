package utility

import "fmt"

type Coordinate struct {
	X int
	Y int
}

type CoordinateRange struct {
	Start *Coordinate
	End   *Coordinate
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
				aCoordinate := &Coordinate{
					X: xNumbers[j],
					Y: yNumbers[j],
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
