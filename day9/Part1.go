package main

import (
	"adventcodingchallenge_2022/utility"
	"strconv"
	"strings"
)

type Orientation int64

const (
	Right Orientation = iota
	Up
	Left
	Down
)

type Command struct {
	direction Orientation
	move      int
}

type Part1 struct {
	answer int
}

func parse(data []string) (commands []*Command, coordinates []*utility.Coordinate) {

	//commands := make([]*Command, 0)
	for _, aRow := range data {
		words := strings.Split(aRow, " ")
		var direction Orientation
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

	//var coordinates []*utility.Coordinate

	coordinates = append(coordinates, &utility.Coordinate{
		X: 0,
		Y: 0,
	})

	for _, aCommand := range commands {

		previousCoordinate := coordinates[len(coordinates)-1]
		coordinate := &utility.Coordinate{
			X: previousCoordinate.X,
			Y: previousCoordinate.Y,
		}
		if aCommand.direction == Up {
			coordinate.X += aCommand.move
		} else if aCommand.direction == Down {
			coordinate.X -= aCommand.move
		} else if aCommand.direction == Right {
			coordinate.Y += aCommand.move
		} else if aCommand.direction == Left {
			coordinate.Y -= aCommand.move
		}
		coordinates = append(coordinates, coordinate)

	}
	return
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	commands, coordinates := parse(data)
	grid := &utility.Grid{}

	return nil, alg.answer
}
