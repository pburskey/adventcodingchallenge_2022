package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
)

const day = "7"

func main() {

	data, err := utility.ParseDayForInputIntoStringRows(day, "input.txt")
	if err != nil {
		panic(err)
	}

	data = utility.ParseInputStringIntoArray(data[0])

	_, solution_a := solution_part_a(data)
	fmt.Println(fmt.Sprintf("Result Part 1: %d", solution_a))

	_, solution_b := solution_part_b(data)
	fmt.Println(fmt.Sprintf("Result Part 2: %d", solution_b))

}

func solution_part_a(data []string) (error, interface{}) {
	algorithm := &Part1{}
	_, solution := algorithm.Process(data)
	return nil, solution

}

func solution_part_b(data []string) (error, interface{}) {
	algorithm := &Part2{}
	_, solution := algorithm.Process(data)
	return nil, solution

}
