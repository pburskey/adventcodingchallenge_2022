package main

type Part2 struct {
	answer int
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	grid := parse(data)
	if grid != nil {
		//alg.answer += grid.perimeter()
		for x := 0; x < len(grid.data); x++ {
			for y := 0; y < len(grid.data[x]); y++ {
				multipliers := make([]int, 0)
				targetValue := grid.data[x][y]
				above, below, left, right := grid.collectElementsFromPosition(x, y)

				multipliers = append(multipliers, len(ValuesUntilEqualToGreaterThan(targetValue, above)))
				multipliers = append(multipliers, len(ValuesUntilEqualToGreaterThan(targetValue, below)))
				multipliers = append(multipliers, len(ValuesUntilEqualToGreaterThan(targetValue, left)))
				multipliers = append(multipliers, len(ValuesUntilEqualToGreaterThan(targetValue, right)))
				scenicScore := 1
				for _, score := range multipliers {
					if score == 0 {
						score = 1
					}
					scenicScore = scenicScore * score
				}
				if scenicScore > alg.answer {
					alg.answer = scenicScore
				}
			}

		}

	}
	return nil, alg.answer
}

func ValuesUntilEqualToGreaterThan(target int, source []int) (results []int) {
	if source != nil {
		for _, aNumber := range source {
			results = append(results, aNumber)

			if aNumber >= target {
				break
			}
		}
	}
	return
}
