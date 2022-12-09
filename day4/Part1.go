package main

import (
	"adventcodingchallenge_2022/utility"
	"strconv"
	"strings"
)

type Part1 struct {
	total int
}

type SectionRange struct {
	sections []int
}

func (sectionRange *SectionRange) parse(data string) {

	parts := strings.Split(data, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	for i := start; i <= end; i++ {
		sectionRange.sections = append(sectionRange.sections, i)

	}

}

type Pair struct {
	a SectionRange
	b SectionRange
}

func (pair *Pair) isOneASubSetOfTheOTher() bool {

	intersection := utility.SubSetPresent(pair.a.sections, pair.b.sections)

	return intersection
}

func (pair *Pair) doTheyIntersect() bool {

	var numberSets [][]int
	numberSets = append(numberSets, pair.a.sections)
	numberSets = append(numberSets, pair.b.sections)

	intersection := utility.IntersectionOfNumbers(numberSets)

	return intersection != nil && len(intersection) > 0
}

func (pair *Pair) parse(data string) {
	parts := strings.Split(data, ",")
	pair.a.parse(parts[0])
	pair.b.parse(parts[1])

}

func (alg *Part1) Process(data []string) (error, int) {

	var pairs []*Pair
	for _, aRow := range data {

		pair := &Pair{}
		pair.parse(aRow)
		if pair.isOneASubSetOfTheOTher() {
			alg.total++
		}

		pairs = append(pairs, pair)

	}
	return nil, alg.total
}
