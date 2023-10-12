package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
)

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
		start:  "S",
		end:    "z",
		matrix: matrix,
	}
	configuration.LoadStartAndEnd(grid)

	matrix.prepareForBFS(configuration.startCell, next, sortCellsAscending)

	configuration.bfs2(configuration.startCell)
	//shortestStack := configuration.shortestBFS()
	//alg.answer = shortestStack.Size()
	return nil, alg.answer
}

func (me *Configuration) dfs(current *Cell, end *Cell, workingStack *utility.SimpleStack, visitedStack *utility.SimpleStack) {
	workingStack.Push(current)
	visitedStack.Push(current)
	if current == end {
		return
	}
	nodeList := me.matrix.nodeListMap[current.id]
	if nodeList != nil {
		for _, aNode := range nodeList.nodes {
			if aNode != nil && me.isNavigable(current, aNode, visitedStack) {
				me.dfs(aNode, end, workingStack, visitedStack)
				if workingStack.Peek() != aNode {
					panic("should be the same....")
				}
				workingStack.Pop()
			}
		}
	}
}

func prettyPrintStack(stack *utility.SimpleStack) {
	items := stack.Items()
	size := len(items)
	fmt.Print(fmt.Sprintf("Size: %10d [", size))

	for x, item := range items {
		fmt.Print(fmt.Sprintf("%s:%s", item.(*Cell).id, item.(*Cell).z))
		if x < size-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("]")
	fmt.Println("")
}

func (me *Configuration) bfs(current *Cell, end *Cell, visitedStack *utility.SimpleStack, queue *utility.SimpleQueue) (shortestStack *utility.SimpleStack) {

	/*
		visit
	*/
	visitedStack.Push(current)

	if current == end {
		prettyPrintStack(visitedStack)
	} else {
		nodeList := me.matrix.nodeListMap[current.id]
		if nodeList != nil {
			for _, aNode := range nodeList.nodes {
				if aNode != nil && me.isNavigable(current, aNode, visitedStack) {
					/*
						queue it
					*/
					queue.Enqueue(aNode)

					/*
						want to continue with a copy of the visited stack
					*/
					//visitorCopy := me.bfs(aNode, end, visitedStack.Copy(), queue)
					aTempStack := me.bfs(aNode, end, visitedStack, queue).Copy()
					if aTempStack != nil && aTempStack.Peek() == end {

						if shortestStack == nil {
							shortestStack = aTempStack
						} else if aTempStack.Size() < shortestStack.Size() {
							shortestStack = aTempStack
						}
					}
				}
			}
		}
	}

	/*
		dequeue current node
	*/
	if ok, deQueueElement := queue.Dequeue(); !ok {
		if deQueueElement != nil && deQueueElement != current {
			panic("dequeued element does not match current element")
		}
	}

	if shortestStack == nil {
		visitedStack.Pop()
		shortestStack = visitedStack
	}
	return shortestStack
}

func (me *Configuration) bfs2(node *Cell) (shortestStack *utility.SimpleStack) {

	queue := utility.NewSimpleQueue()
	visited := make(map[string]string)

	visited[node.id] = node.id
	queue.Enqueue(node)

	for !queue.IsEmpty() {
		if ok, temp := queue.Dequeue(); ok {
			node = temp.(*Cell)
		}

		nodeList := me.matrix.nodeListMap[node.id]
		if nodeList != nil {
			for _, aChildNode := range nodeList.nodes {
				if aChildNode != nil && me.isNavigable(node, aChildNode, nil) {

					if _, ok := visited[aChildNode.id]; !ok {
						visited[aChildNode.id] = aChildNode.id
						queue.Enqueue(aChildNode)
					}

				}
			}
		}

	}
	return nil
}
