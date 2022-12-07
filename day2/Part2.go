package main

import (
	"adventcodingchallenge_2022/utility"
	"strconv"
)

type Part2 struct {
	answer int
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	currentTotal := 0
	var totals []int
	for _, aRow := range data {

		if aRow == "" {
			totals = append(totals, currentTotal)

			currentTotal = 0

		} else {
			aNumber, _ := strconv.Atoi(aRow)
			currentTotal += aNumber
		}

	}

	totals = utility.OrderNumbersSortReversed(totals)
	alg.answer += totals[0]
	alg.answer += totals[1]
	alg.answer += totals[2]
	return nil, alg

}
