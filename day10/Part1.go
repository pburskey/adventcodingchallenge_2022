package main

import (
	"adventcodingchallenge_2022/utility"
	"strconv"
	"strings"
)

type InstructionType int64

const (
	addX InstructionType = iota
	noop
)

const (
	addxCycleCost = 2
	noopCycleCost = 1
)

type Instruction struct {
	instructionType InstructionType
	cycleCost       int
}

type ProgramStep struct {
	instruction *Instruction
	input       interface{}
	cyclesSpent int
	processed   bool
}

func (p *ProgramStep) execute() bool {
	return p.instruction.cycleCost == p.cyclesSpent
}

func (p *ProgramStep) HasBeenProcessed() bool {
	return p.processed
}

type Part1 struct {
	answer int
}

var INSTRUCTION_ADDX *Instruction
var INSTRUCTION_NOOP *Instruction

func parseCommands(data []string) (steps []*ProgramStep) {

	INSTRUCTION_ADDX = &Instruction{
		instructionType: addX,
		cycleCost:       addxCycleCost,
	}

	INSTRUCTION_NOOP = &Instruction{
		instructionType: noop,
		cycleCost:       noopCycleCost,
	}

	for _, aRow := range data {
		words := strings.Split(aRow, " ")
		var instruction *Instruction
		if words[0] == "addx" {
			instruction = INSTRUCTION_ADDX
		} else if words[0] == "noop" {
			instruction = INSTRUCTION_NOOP
		}
		var input interface{}
		if len(words) > 1 {
			input, _ = strconv.Atoi(words[1])
		}
		steps = append(steps, &ProgramStep{
			instruction: instruction,
			input:       input,
		})

	}
	return
}

type InstructionSet struct {
	instructionSteps *utility.SimpleQueue
	processedSteps   *utility.SimpleStack
	cycles           []int
	register         int
}

func (me *InstructionSet) signalStrength(startCycle int, variationAfterStart int) map[int]int {
	aCycleMap := make(map[int]int)

	for cycleIndex, aCycleRegisterValue := range me.cycles {
		oneBasedIndex := cycleIndex + 1
		interested := false

		if oneBasedIndex == startCycle {
			interested = true
		} else if oneBasedIndex > startCycle {
			remainder := (oneBasedIndex - startCycle) % variationAfterStart
			interested = (remainder == 0)
		}

		if interested {
			thisCycleValue := oneBasedIndex * aCycleRegisterValue
			aCycleMap[oneBasedIndex] = thisCycleValue
		}
	}
	return aCycleMap
}

func (me *InstructionSet) start() *ProgramStep {
	var currentInstruction *ProgramStep
	if me.processedSteps.HasMore() {
		currentInstruction = me.processedSteps.Peek().(*ProgramStep)
	}

	if currentInstruction == nil || currentInstruction.HasBeenProcessed() {
		currentInstruction = nil
		if ok, aQueueItem := me.instructionSteps.Dequeue(); ok {
			currentInstruction = aQueueItem.(*ProgramStep)
			me.processedSteps.Push(currentInstruction)
		}
	}
	return currentInstruction
}

func (me *InstructionSet) during(currentInstruction *ProgramStep) {
	me.cycles = append(me.cycles, me.register)
}

func (me *InstructionSet) finish(currentInstruction *ProgramStep) {

	if currentInstruction != nil {

		if currentInstruction.execute() {
			if currentInstruction.instruction.instructionType == noop {
				// nothing to see here, move along

			} else if currentInstruction.instruction.instructionType == addX {
				value := currentInstruction.input.(int)
				me.register += value
			}
			currentInstruction.processed = true
		}
		currentInstruction.cyclesSpent++
	}

}

func (me *InstructionSet) execute() {

	var currentInstruction *ProgramStep
	shutDown := false
	for !shutDown {

		currentInstruction = me.start()
		me.during(currentInstruction)
		me.finish(currentInstruction)

		shutDown = currentInstruction == nil
	}
}

func NewInstructionSet(programSteps []*ProgramStep) *InstructionSet {
	queue := utility.NewSimpleQueue()
	for _, aStep := range programSteps {
		queue.Enqueue(aStep)
	}
	instructionSet := &InstructionSet{
		instructionSteps: queue,
		processedSteps:   utility.NewSimpleStack(),
		cycles:           make([]int, 0),
		register:         1,
	}
	return instructionSet
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	programSteps := parseCommands(data)
	instructionSet := NewInstructionSet(programSteps)
	instructionSet.execute()

	aCycleMap := instructionSet.signalStrength(20, 40)
	for _, aValue := range aCycleMap {
		alg.answer += aValue
	}

	return nil, alg.answer
}
