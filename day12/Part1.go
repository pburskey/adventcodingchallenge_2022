package main

type Part1 struct {
	answer int
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	grid := &Grid{
		cells: make([][]*Cell, 0),
	}

	parseCommands(data, grid)
	matrix := &AdjacencyMatrix{}
	matrix.from(grid)

	configuration := &Configuration{
		start: int('S' - '0'),
		end:   int('E' - '0'),
	}
	configuration.LoadStartAndEnd(grid)
	travel(configuration.startCell, configuration.endCell, configuration, matrix, make(map[string]int))
	return nil, alg.answer
}

func travel(start *Cell, end *Cell, configuration *Configuration, matrix *AdjacencyMatrix, visited map[string]int) {
	if visited[start.id] > 0 {
		return
	}
	visited[start.id]++
	current := start
	nodeList := matrix.nodeListMap[current.id]
	if nodeList != nil {
		for _, aNeighborCell := range nodeList.nodes {
			startingHeight := configuration.nextHeightFrom(current)
			if aNeighborCell != nil {
				ok := false

				targetHeight := aNeighborCell.z

				if startingHeight == targetHeight {
					ok = true
				} else if targetHeight < startingHeight {
					ok = true
				} else if (targetHeight - 1) == startingHeight {
					ok = true
				}
				if ok {
					travel(aNeighborCell, end, configuration, matrix, visited)
				}
			}
		}
	}

}
