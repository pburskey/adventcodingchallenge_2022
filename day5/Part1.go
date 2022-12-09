package main

import (
	"adventcodingchallenge_2022/utility"
	"strconv"
	"strings"
)

type Part1 struct {
	answer string
}

type StackOfCrates struct {
	crates utility.SimpleStringStack
}

type Instruction struct {
	move int
	from int
	to   int
}

func ParseStacks(data []string) (stacks []*StackOfCrates) {

	spacingRow := data[len(data)-1]
	spacingRowWidth := len(spacingRow)

	splits := strings.Split(spacingRow, " ")
	columns := 0
	if splits != nil {
		for _, aSplit := range splits {
			if aSplit != "" {
				columns++
			}
		}
	}
	columnWidth := 3

	for i := 0; i < columns; i++ {
		stacks = append(stacks, &StackOfCrates{})
	}

	for _, value := range data {

		start := 0
		end := 0
		spacingRowWidth = len(value)
		for aColumn := 0; start < spacingRowWidth; aColumn++ {
			end = start + columnWidth
			chunk := ""
			if end > spacingRowWidth {
				chunk = value[start:]
			} else {
				chunk = value[start:end]
			}

			if len(strings.TrimSpace(chunk)) == columnWidth {
				crate := chunk[1:2]
				stacks[aColumn].crates.Push(crate)

			}
			start = end + 1

		}
	}

	for _, aStack := range stacks {
		aStack.crates.Reverse()
	}

	return stacks
}

func ParseInstructions(data []string) []*Instruction {
	var instructions []*Instruction
	for _, aRow := range data {

		splits := strings.Split(aRow, " ")
		move, _ := strconv.Atoi(splits[1])
		from, _ := strconv.Atoi(splits[3])
		to, _ := strconv.Atoi(splits[5])
		instructions = append(instructions, &Instruction{move: move, from: from, to: to})

	}
	return instructions
}

func parseStacksAndInstructions(data []string) (stacks []*StackOfCrates, instructions []*Instruction) {

	var stackStrings []string
	var instructionStrings []string

	foundInstructions := false
	for _, aRow := range data {

		if aRow == "" {
			foundInstructions = true
		} else {
			if foundInstructions {
				instructionStrings = append(instructionStrings, aRow)
			} else {
				stackStrings = append(stackStrings, aRow)
			}
		}

	}

	stacks = ParseStacks(stackStrings)
	instructions = ParseInstructions(instructionStrings)
	return
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	stacks, instructions := parseStacksAndInstructions(data)

	for _, instruction := range instructions {
		move := instruction.move
		from := instruction.from - 1
		to := instruction.to - 1

		for i := 0; i < move; i++ {
			crate := stacks[from].crates.Pop()
			stacks[to].crates.Push(crate)
		}
	}

	var topCrates []string
	for i := 0; i < len(stacks); i++ {
		topCrates = append(topCrates, stacks[i].crates.Pop())
	}
	alg.answer = strings.Join(topCrates, "")
	return nil, alg.answer
}
