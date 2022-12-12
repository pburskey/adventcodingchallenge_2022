package main

import "adventcodingchallenge_2022/utility"

type Part1 struct {
	answer int
}

func parse(data []string) [][]int {

	var grid [][]int
	grid = utility.ParseRowsToInts(data)
	return grid

}

func (alg *Part1) Process(data []string) (error, interface{}) {
	grid := parse(data)
	if grid != nil {
		outside
	}
	return nil, alg.answer
}
