package main

type Part2 struct {
	total int
}

func (alg *Part2) Process(data []string) (error, int) {

	var pairs []*Pair
	for _, aRow := range data {

		pair := &Pair{}
		pair.parse(aRow)
		if pair.doTheyIntersect() {
			alg.total++
		}

		pairs = append(pairs, pair)

	}
	return nil, alg.total

}
