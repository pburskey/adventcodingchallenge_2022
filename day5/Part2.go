package main

import (
	"adventcodingchallenge_2022/utility"
	"strings"
)

type Part2 struct {
	answer string
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	stacks, instructions := parseStacksAndInstructions(data)

	for _, instruction := range instructions {
		move := instruction.move
		from := instruction.from - 1
		to := instruction.to - 1

		simpleStack := &utility.SimpleStringStack{}
		for i := 0; i < move; i++ {
			crate := stacks[from].crates.Pop()
			simpleStack.Push(crate)
		}

		for simpleStack.HasMore() {
			stacks[to].crates.Push(simpleStack.Pop())
		}

	}

	var topCrates []string
	for i := 0; i < len(stacks); i++ {
		topCrates = append(topCrates, stacks[i].crates.Pop())
	}
	alg.answer = strings.Join(topCrates, "")
	return nil, alg.answer

}
