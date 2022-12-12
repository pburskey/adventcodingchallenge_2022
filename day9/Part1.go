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

func parse(data []string) []*Command {

	commands := make([]*Command, 0)
	gridData := utility.ParseRowsToInts(data)
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
	return commands
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	commands := parse(data)

	var coordinates []*utility.Coordinate

	coordinate := &utility.Coordinate{
		X: 0,
		Y: 0,
	}

	coordinates = append(coordinates, coordinate)

	for _, aCommand := range commands {

		previousCoordinate := coordinates[len(coordinates)-1]
		currentCoordinate := &utility.Coordinate{
			X: previousCoordinate.X,
			Y: previousCoordinate.Y,
		}
		if aCommand.direction == Up {
			currentCoordinate.X += aCommand.move
		} else if aCommand.direction == Down {
			currentCoordinate.X -= aCommand.move
		} else if aCommand.direction == Right {
			currentCoordinate.Y += aCommand.move
		} else if aCommand.direction == Left {
			currentCoordinate.Y -= aCommand.move
		}

	}

	return nil, alg.answer
}
