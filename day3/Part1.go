package main

import (
	"adventcodingchallenge_2022/utility"
)

type Part1 struct {
	total int
}

type Sack struct {
	a string
	b string
}

func initializePriorities() map[rune]int {

	priorities := utility.AssignNumbersToLetters(1, 26)
	return priorities

}

func (sack *Sack) priority(priorities map[rune]int, aRune rune) int {

	aPriority := priorities[aRune]
	//fmt.Println(fmt.Sprintf("Rune %U Priority: %d", aRune, aPriority))
	return aPriority

}

func (sack *Sack) commonItems() []rune {

	commonItems := utility.IntersectionOfTwoStrings(sack.a, sack.b)

	return commonItems

}

func (alg *Part1) Process(data []string) (error, int) {

	priorities := initializePriorities()

	if priorities['a'] != 1 {
		panic("priorities failure")
	}

	if priorities['A'] != 27 {
		panic("priorities failure")
	}
	if priorities['z'] != 26 {
		panic("priorities failure")
	}
	if priorities['Z'] != 52 {
		panic("priorities failure")
	}

	for _, aRow := range data {

		middle := len(aRow) / 2
		firstHalf := aRow[0:middle]
		secondHalf := aRow[middle:len(aRow)]
		sack := &Sack{a: firstHalf, b: secondHalf}

		commonItems := sack.commonItems()
		if commonItems != nil {
			for _, aRune := range commonItems {

				aPriority := sack.priority(priorities, aRune)
				alg.total += aPriority
			}
		}

	}
	return nil, alg.total
}
