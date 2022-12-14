package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
)

type Part2 struct {
	answer   int
	segments int
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	commands := parseCommands(data)
	segments := make([][]*utility.Coordinate, 0)
	rope := &Rope{segments: segments}

	rope.run(alg.segments, commands, true)

	unique := make(map[string]int)

	history := rope.getHistoryOfSegment(len(rope.segments) - 1)
	for _, aCoordinate := range history {
		key := fmt.Sprintf("x:%d y:%d", aCoordinate.X, aCoordinate.Y)
		unique[key]++
	}

	alg.answer = len(unique)
	//for key, value := range unique {
	//	log.Println(fmt.Sprintf("Visited: %s :%d times...", key, value))
	//	alg.answer++
	//}

	return nil, alg.answer
}
