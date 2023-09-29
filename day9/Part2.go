package main

import (
	"fmt"
)

type Part2 struct {
	answer   int
	segments int
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	commands := parseCommands(data)

	rope := Rope{}

	rope.run(alg.segments, commands, false)

	unique := make(map[string]int)

	for _, aCoordinate := range rope.history {
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
