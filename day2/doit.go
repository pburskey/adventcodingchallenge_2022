package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
)

const day = "1"

func main() {

	data, err := utility.ParseDayForInputIntoStringRows(day, "data_test.txt")
	if err != nil {
		panic(err)
	}

	sanity_part_a(data, 15)

	data, err = utility.ParseDayForInputIntoStringRows(day, "input.txt")
	if err != nil {
		panic(err)
	}

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

func sanity_part_a(data []string, answer int) (error, interface{}) {
	algorithm := &Part1{}
	_, solution := algorithm.Process(data)

	if answer != algorithm.total {
		panic("Test results failing")
	} else {
		fmt.Println(fmt.Sprintf("Tested Successful with result: %d", answer))

	}
	return nil, solution

}
