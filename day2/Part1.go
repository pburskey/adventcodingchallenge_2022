package main

import (
	"strings"
)

type Part1 struct {
	total int
}

type KindOfPiece int

const (
	Rock KindOfPiece = iota
	Paper
	Scizzors
)

func (alg *Part1) play(opponentPlay string, youPlay string) (opponentWin bool, youWin bool, youScore int) {
	var opponent KindOfPiece
	var you KindOfPiece

	youScore = 0

	if opponentPlay == "A" {
		opponent = Rock
	} else if opponentPlay == "B" {
		opponent = Paper
	} else if opponentPlay == "C" {
		opponent = Scizzors
	}

	if youPlay == "X" {
		you = Rock
		youScore += 1
	} else if youPlay == "Y" {
		you = Paper
		youScore += 2
	} else if youPlay == "Z" {
		you = Scizzors
		youScore += 3
	}

	if opponent == Rock && you == Scizzors {
		opponentWin = true
	} else if opponent == Scizzors && you == Paper {
		opponentWin = true
	} else if opponent == Paper && you == Rock {
		opponentWin = true
	} else if opponent == you {

	} else {
		youWin = true
	}

	if youWin {
		youScore += 6
	} else if opponentWin {
		youScore += 0
	} else if opponentWin == youWin {
		youScore += 3
	}

	return
}

func (alg *Part1) Process(data []string) (error, int) {

	for _, aRow := range data {

		words := strings.Fields(aRow)
		opponent := words[0]
		you := words[1]

		_, _, yourScore := alg.play(opponent, you)

		alg.total += yourScore

	}
	return nil, alg.total
}
