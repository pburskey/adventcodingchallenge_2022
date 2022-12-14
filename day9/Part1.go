package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
	"strconv"
	"strings"
)

type Direction int64

const (
	Right Direction = iota
	Up
	Left
	Down
)

type Command struct {
	direction Direction
	move      int
}

type Part1 struct {
	answer   int
	segments int
}

func parseCommands(data []string) (commands []*Command) {

	//commands := make([]*Command, 0)
	for _, aRow := range data {
		words := strings.Split(aRow, " ")
		var direction Direction
		if words[0] == "R" {
			direction = Right
		} else if words[0] == "U" {
			direction = Up
		} else if words[0] == "D" {
			direction = Down
		} else if words[0] == "L" {
			direction = Left
		}
		move, _ := strconv.Atoi(words[1])
		commands = append(commands, &Command{
			direction: direction,
			move:      move,
		})

	}
	return
}

type Rope struct {
	segments [][]*utility.Coordinate
	//headCoordinates []*utility.Coordinate
	//tailCoordinates []*utility.Coordinate
	//knots []*utility.Coordinate
}

func (r *Rope) getCurrentPositionOfSegment(segment int) *utility.Coordinate {
	var coordinate *utility.Coordinate
	if r.segments != nil && len(r.segments) > 0 {
		history := r.segments[segment]
		coordinate = history[len(history)-1]
	}
	return coordinate
}

func (r *Rope) getCurrentPositionOfAllSegments() []*utility.Coordinate {
	coordinates := make([]*utility.Coordinate, 0)
	for _, aSegmentHistory := range r.segments {
		aCoordinate := aSegmentHistory[len(aSegmentHistory)-1]
		coordinates = append(coordinates, aCoordinate)
	}

	return coordinates
}

func (r *Rope) getPreviousPositionOfSegment(segment int) *utility.Coordinate {
	var coordinate *utility.Coordinate

	history := r.getHistoryOfSegment(segment)
	if history != nil && len(history) > 1 {
		coordinate = history[len(history)-2]
	}

	return coordinate
}

func (r *Rope) getHistoryOfSegment(anIndex int) []*utility.Coordinate {
	var coordinates []*utility.Coordinate
	if r.segments != nil && len(r.segments) > 0 {
		history := r.segments[anIndex]
		coordinates = history
	}
	return coordinates
}

func (r *Rope) getHead() *utility.Coordinate {
	var coordinate *utility.Coordinate
	if r.segments != nil && len(r.segments) > 0 {
		coordinate = r.getCurrentPositionOfSegment(0)
	}
	return coordinate
}

func (r *Rope) advanceHead(aDirection Direction, step int) {
	head := r.getHead()

	nextHeadPosition := r.newCoordinateForSegment(head.X, head.Y, 0, len(r.segments))
	//nextHeadPosition := &utility.Coordinate{
	//	X: head.X,
	//	Y: head.Y,
	//}

	if aDirection == Up {
		nextHeadPosition.X += step
	} else if aDirection == Down {
		nextHeadPosition.X -= step
	} else if aDirection == Right {
		nextHeadPosition.Y += step
	} else if aDirection == Left {
		nextHeadPosition.Y -= step
	}
	r.segments[0] = append(r.segments[0], nextHeadPosition)
}

func (r *Rope) advanceSegment(segment int) {

	headIndex := segment - 1

	head := r.getCurrentPositionOfSegment(headIndex)
	tail := r.getCurrentPositionOfSegment(segment)

	if tail.IsInSamePositionAs(head) || tail.IsAdjacent(head) {
		//all good
	} else {

		previousHead := r.getPreviousPositionOfSegment(headIndex)
		nextTailPosition := r.newCoordinateForSegment(previousHead.X, previousHead.Y, segment, len(r.segments))
		//nextTailPosition := &utility.Coordinate{
		//	X: previousHead.X,
		//	Y: previousHead.Y,
		//}

		r.segments[segment] = append(r.segments[segment], nextTailPosition)

	}

}

func (r *Rope) advanceSegmentsStartingAt(start int) {

	for i := start; i < len(r.segments); i++ {
		r.advanceSegment(i)

		head := r.getCurrentPositionOfSegment(i - 1)
		tail := r.getCurrentPositionOfSegment(i)

		adjactent := tail.IsAdjacent(head)

		if !adjactent {
			panic("Tail is not adjacent or in same position as head after advancing")
		}
	}
}

func (r *Rope) newCoordinateForSegment(x int, y int, segment int, segmentCount int) *utility.Coordinate {
	value := ""
	//if segment == 0{
	//	value = "H"
	//} else{
	value = fmt.Sprintf("%d", segment)
	//}
	aCoordinate := &utility.Coordinate{
		X:     x,
		Y:     y,
		Value: value,
	}
	return aCoordinate
}

func (r *Rope) run(segmentsCount int, commands []*Command, debug bool) {

	for i := 0; i < segmentsCount; i++ {

		aCoordinate := r.newCoordinateForSegment(0, 0, i, segmentsCount)
		segments := make([]*utility.Coordinate, 0)
		r.segments = append(r.segments, segments)
		r.segments[i] = append(r.segments[i], aCoordinate)

	}

	for _, aCommand := range commands {
		fmt.Println(fmt.Sprintf("== %v %d ==", aCommand.direction, aCommand.move))
		for i := 0; i < aCommand.move; i++ {
			r.advanceHead(aCommand.direction, 1)
			r.advanceSegmentsStartingAt(1)
			if debug {
				utility.PrettyPrint(r.getCurrentPositionOfAllSegments())
				fmt.Println("")
			}

		}

	}
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	commands := parseCommands(data)
	segments := make([][]*utility.Coordinate, 0)
	rope := &Rope{segments: segments}

	rope.run(alg.segments, commands, false)

	unique := make(map[string]int)

	for _, aCoordinate := range rope.getHistoryOfSegment(len(rope.segments) - 1) {
		key := fmt.Sprintf("x:%d y:%d", aCoordinate.X, aCoordinate.Y)
		unique[key]++
	}

	utility.PrettyPrint(rope.getHistoryOfSegment(len(rope.segments) - 1))

	alg.answer = len(unique)
	//for key, value := range unique {
	//	log.Println(fmt.Sprintf("Visited: %s :%d times...", key, value))
	//	alg.answer++
	//}

	return nil, alg.answer
}
