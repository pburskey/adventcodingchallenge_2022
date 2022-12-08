package main

import (
	"strings"
)

type Part2 struct {
	total int
}

func (alg *Part2) play(opponentPlay string, youPlay string) (youScore int) {
	var opponent KindOfPiece
	var you KindOfPiece

	var youShouldWin, youShouldDraw, youShouldLose bool
	youScore = 0

	if opponentPlay == "A" {
		opponent = Rock
	} else if opponentPlay == "B" {
		opponent = Paper
	} else if opponentPlay == "C" {
		opponent = Scizzors
	}

	if youPlay == "X" {
		youShouldLose = true
	} else if youPlay == "Y" {
		youShouldDraw = true
	} else if youPlay == "Z" {
		youShouldWin = true
	}

	if youShouldDraw {
		you = opponent
	} else if youShouldLose {
		if opponent == Rock {
			you = Scizzors
		} else if opponent == Scizzors {
			you = Paper
		} else if opponent == Paper {
			you = Rock
		}
	} else if youShouldWin {
		if opponent == Rock {
			you = Paper
		} else if opponent == Scizzors {
			you = Rock
		} else if opponent == Paper {
			you = Scizzors
		}
	}

	if you == Rock {
		youScore += 1
	} else if you == Paper {
		youScore += 2
	} else if you == Scizzors {
		youScore += 3
	}

	if youShouldWin {
		youScore += 6
	} else if youShouldLose {
		youScore += 0
	} else if youShouldDraw {
		youScore += 3
	}

	return
}
func (alg *Part2) Process(data []string) (error, interface{}) {

	for _, aRow := range data {

		words := strings.Fields(aRow)
		opponent := words[0]
		you := words[1]

		yourScore := alg.play(opponent, you)

		alg.total += yourScore

	}
	return nil, alg.total

}
