package main

import (
	"strconv"
)

type Part1 struct {
	total int
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	currentTotal := 0

	for _, aRow := range data {

		if aRow == "" {

			if currentTotal > alg.total {
				alg.total = currentTotal
			}
			currentTotal = 0

		} else {
			aNumber, _ := strconv.Atoi(aRow)
			currentTotal += aNumber
		}

	}
	return nil, alg.total
}
