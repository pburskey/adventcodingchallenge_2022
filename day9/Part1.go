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
	segments []*utility.Coordinate
	history  []*utility.Coordinate
}

func (r *Rope) getCurrentPositionOfSegment(segment int) *utility.Coordinate {
	var coordinate *utility.Coordinate
	if r.segments != nil && len(r.segments) > 0 {
		coordinate = r.segments[segment]
	}
	return coordinate
}

func (r *Rope) getCurrentPositionOfAllSegments() []*utility.Coordinate {
	//coordinates := make([]*utility.Coordinate, 0)
	//for _, aSegmentHistory := range r.segments {
	//	aCoordinate := aSegmentHistory[len(aSegmentHistory)-1]
	//	coordinates = append(coordinates, aCoordinate)
	//}

	return r.segments
}

//
//func (r *Rope) getPreviousPositionOfSegment(segment int) *utility.Coordinate {
//	var coordinate *utility.Coordinate
//
//	history := r.getHistoryOfSegment(segment)
//	if history != nil && len(history) > 1 {
//		coordinate = history[len(history)-2]
//	}
//
//	return coordinate
//}
//
//func (r *Rope) getHistoryOfSegment(anIndex int) []*utility.Coordinate {
//	var coordinates []*utility.Coordinate
//	if r.segments != nil && len(r.segments) > 0 {
//		history := r.segments[anIndex]
//		coordinates = history
//	}
//	return coordinates
//}

func (r *Rope) getHead() *utility.Coordinate {
	var coordinate *utility.Coordinate
	if r.segments != nil && len(r.segments) > 0 {
		coordinate = r.getCurrentPositionOfSegment(0)
	}
	return coordinate
}

func (r *Rope) advanceHead(aDirection Direction, step int) {
	r.advanceSegment(0, aDirection, step)
}

func (r *Rope) advanceSegment(segmentIndex int, aDirection Direction, step int) {

	segment := r.segments[segmentIndex]

	if aDirection == Up {
		segment.Y += step
	} else if aDirection == Down {
		segment.Y -= step
	} else if aDirection == Right {
		segment.X += step
	} else if aDirection == Left {
		segment.X -= step
	}
	//r.segments[0] = append(r.segments[0], nextHeadPosition)

}

//
//func (r *Rope) advanceSegmentsStartingAt(start int) {
//
//	for i := start; i < len(r.segments); i++ {
//		r.advanceSegment(i)
//
//		head := r.getCurrentPositionOfSegment(i - 1)
//		tail := r.getCurrentPositionOfSegment(i)
//
//		adjactent := tail.IsAdjacent(head)
//
//		if !adjactent {
//			panic("Tail is not adjacent or in same position as head after advancing")
//		}
//	}
//}

func (r *Rope) newCoordinateForSegment(x int, y int, segment int) *utility.Coordinate {
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

	r.segments = make([]*utility.Coordinate, segmentsCount)
	r.history = make([]*utility.Coordinate, 0)
	for i := 0; i < segmentsCount; i++ {

		r.segments[i] = r.newCoordinateForSegment(0, 0, i)
		//r.segments = append(r.segments, aCoordinate)

	}

	for _, aCommand := range commands {
		fmt.Println(fmt.Sprintf("== %v %d ==", aCommand.direction, aCommand.move))
		for i := 0; i < aCommand.move; i++ {
			r.advanceSegments(aCommand.direction, 1)
		}
		utility.PrettyPrint(r.segments)
	}
	if debug {
		//utility.PrettyPrint(r.getCurrentPositionOfAllSegments())
		fmt.Println("")
	}
}

func (r *Rope) advanceSegments(direction Direction, distance int) {
	r.advanceHead(direction, distance)

	/*
		If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step in that direction so it remains close enough:
	*/
	coordinates := r.getCurrentPositionOfAllSegments()
	leader := coordinates[0]
	currentSegment := coordinates[1]
	distanceCoordinate := leader.DistanceBetween(currentSegment)

	if !currentSegment.IsAdjacent(leader) {

		aCoordinateRange := utility.CoordinateRange{Start: currentSegment, End: leader}

		if currentSegment.IsInSameXAs(leader) && distanceCoordinate.Y == 2 {
			coordinatesBetween := aCoordinateRange.DetermineCoordinatesInRange(false)
			coordinatesBetween = aCoordinateRange.FilterStartAndEnd(coordinatesBetween)
			for _, aCoordinate := range coordinatesBetween {
				currentSegment.MoveTo(aCoordinate)
			}

		} else if currentSegment.IsInSameYAs(leader) && distanceCoordinate.X == 2 {
			coordinatesBetween := aCoordinateRange.DetermineCoordinatesInRange(false)
			coordinatesBetween = aCoordinateRange.FilterStartAndEnd(coordinatesBetween)
			for _, aCoordinate := range coordinatesBetween {
				currentSegment.MoveTo(aCoordinate)
			}
		} else {
			/*
				Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally to keep up:
			*/
			coordinatesBetween := aCoordinateRange.DetermineCoordinatesInRange(true)
			coordinatesBetween = aCoordinateRange.FilterStartAndEnd(coordinatesBetween)
			for _, aCoordinate := range coordinatesBetween {
				currentSegment.MoveTo(aCoordinate)
			}
		}
		r.history = append(r.history, &utility.Coordinate{
			X:     currentSegment.X,
			Y:     currentSegment.Y,
			Value: currentSegment.Value,
		})

		if !currentSegment.IsAdjacent(leader) {
			panic("Tail is not adjacent or in same position as head after advancing")
		}
	}

	//if debug {
	//	utility.PrettyPrint(r.getCurrentPositionOfAllSegments())
	//	fmt.Println("")
	//}
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	commands := parseCommands(data)

	rope := Rope{}

	rope.run(alg.segments, commands, false)

	unique := make(map[string]int)

	for _, aCoordinate := range rope.history {
		key := fmt.Sprintf("x:%d y:%d", aCoordinate.X, aCoordinate.Y)
		unique[key]++
	}

	alg.answer = len(unique)
	//for key, value := range unique {
	//	log.Println(fmt.Sprintf("Visited: %s :%d times...", key, value))
	//	alg.answer++
	//}

	return nil, alg.answer
}
