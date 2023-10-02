package main

type Part2 struct {
	answer   int
	segments int
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	programSteps := parseCommands(data)
	instructionSet := NewInstructionSet(programSteps)
	instructionSet.execute()

	crt := &CRT{cycles: instructionSet.cycles}
	crt.render()

	return nil, alg.answer

}
