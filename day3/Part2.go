package main

import "adventcodingchallenge_2022/utility"

type Part2 struct {
	total int
}

type Group struct {
	sacks []*Sack
}

func (group *Group) commonItems() []rune {

	var compartments []string
	var commonItems []rune
	if group.sacks != nil {
		for _, item := range group.sacks {
			if item.a != "" && len(item.a) > 0 {
				compartments = append(compartments, item.a)
			}

			if item.b != "" && len(item.b) > 0 {
				compartments = append(compartments, item.b)
			}
		}

		commonItems = utility.IntersectionOf(compartments)
	}

	return commonItems
}

func (group *Group) priority(priorities map[rune]int, aRune rune) int {

	aPriority := priorities[aRune]
	//fmt.Println(fmt.Sprintf("Rune %U Priority: %d", aRune, aPriority))
	return aPriority

}

func (alg *Part2) Process(data []string) (error, int) {

	priorities := initializePriorities()

	if priorities['a'] != 1 {
		panic("priorities failure")
	}

	if priorities['A'] != 27 {
		panic("priorities failure")
	}
	if priorities['z'] != 26 {
		panic("priorities failure")
	}
	if priorities['Z'] != 52 {
		panic("priorities failure")
	}

	group := &Group{}
	for i, aRow := range data {

		if i%3 == 0 {

			commonItems := group.commonItems()
			if commonItems != nil {
				for _, aRune := range commonItems {
					aPriority := group.priority(priorities, aRune)
					alg.total += aPriority
				}
			}

			group = &Group{}

		}
		//middle := len(aRow) / 2
		//firstHalf := aRow[0:middle]
		//secondHalf := aRow[middle:len(aRow)]
		sack := &Sack{a: aRow, b: ""}
		group.sacks = append(group.sacks, sack)

	}
	commonItems := group.commonItems()
	if commonItems != nil {
		for _, aRune := range commonItems {
			aPriority := group.priority(priorities, aRune)
			alg.total += aPriority
		}
	}

	return nil, alg.total

}
