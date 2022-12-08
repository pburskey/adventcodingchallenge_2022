package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
	"unicode"
)

type Part1 struct {
	total int
}

type Sack struct {
	a          string
	b          string
	priorities map[rune]int
}

func initializePriorities() map[rune]int {

	priorities := make(map[rune]int)
	index := 1
	for aLowerCaseLetter := 'a'; aLowerCaseLetter <= 'z'; aLowerCaseLetter++ {
		aCapitalLetter := unicode.ToUpper(aLowerCaseLetter)
		priorities[aLowerCaseLetter] = index
		priorities[aCapitalLetter] = 26 + index
		index++
	}

	return priorities

}

func (sack *Sack) priority(aRune rune) int {
	if sack.priorities == nil {
		sack.priorities = initializePriorities()
	}

	aPriority := sack.priorities[aRune]
	fmt.Println(fmt.Sprintf("Rune %U Priority: %d", aRune, aPriority))
	return aPriority

}

func (sack *Sack) commonItems() []rune {

	commonItems := utility.IntersectionOf(sack.a, sack.b)

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
		sack := &Sack{a: firstHalf, b: secondHalf, priorities: priorities}

		commonItems := sack.commonItems()
		if commonItems != nil {
			for _, aRune := range commonItems {

				aPriority := sack.priority(aRune)
				alg.total += aPriority
			}
		}

	}
	return nil, alg.total
}
