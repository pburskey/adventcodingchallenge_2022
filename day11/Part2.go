package main

import "adventcodingchallenge_2022/utility"

type Part2 struct {
	answer   int
	segments int
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	container := parseCommands(data, false)
	if container != nil {
		for round := 0; round < 10000; round++ {
			for _, monkey := range container.monkeys {
				monkey.conductBusiness(container)
			}
		}

		/*
			find two most active monkeys
		*/
		numbers := make([]int, 0)
		for _, monkey := range container.monkeys {
			aValue := monkey.inspectionCount
			numbers = append(numbers, aValue)
		}
		numbers = utility.OrderNumbersSortReversed(numbers)
		alg.answer = numbers[0] * numbers[1]

	}
	return nil, alg.answer

}
