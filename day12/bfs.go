package main

import (
	"adventcodingchallenge_2022/utility"
	"fmt"
	"sort"
)

type BFS struct {
	matrix      *AdjacencyMatrix
	isNavigable func(*Cell, *Cell) bool
	sortNodes   func([]*Cell) []*Cell
	tree        *Node
}

type Paths struct {
	paths []*utility.SimpleStack
}

func (me *Paths) add(path *utility.SimpleStack) *Paths {
	me.paths = append(me.paths, path)
	me.prettyPrint(path)
	return me
}

func (me *Paths) prettyPrint(stack *utility.SimpleStack) {
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

func (me *Paths) shortest() *utility.SimpleStack {
	var path *utility.SimpleStack
	for _, candidate := range me.paths {
		if path == nil {
			path = candidate
		} else if path != nil && candidate != nil && candidate.Size() > path.Size() {
			path = candidate
		}
	}
	return path
}

func (me *BFS) build(currentNode *Node, visited map[string]*Node, edgeMap map[string]*Node, queue *utility.SimpleQueue) *Node {

	//println(fmt.Sprintf("Visiting Node: %s", currentNode.node.id))
	visited[currentNode.node.id] = currentNode

	nodeList := me.matrix.nodeListMap[currentNode.node.id]
	if nodeList != nil {
		nodes := me.sortNodes(nodeList.gather())
		for _, aPotentialVisitableChildNode := range nodes {
			if aPotentialVisitableChildNode != nil && me.isNavigable(currentNode.node, aPotentialVisitableChildNode) {
				var aVisitableChildNode *Node
				if _, ok := edgeMap[aPotentialVisitableChildNode.id]; !ok {
					aVisitableChildNode = &Node{
						nodes: make([]*Node, 0),
						node:  aPotentialVisitableChildNode,
					}
					edgeMap[aVisitableChildNode.node.id] = aVisitableChildNode
				} else {
					aVisitableChildNode = edgeMap[aPotentialVisitableChildNode.id]
					//println(fmt.Sprintf("Previously visited node: %s", aPotentialVisitableChildNode.id))
				}

				if currentNode.containsChildNode(aVisitableChildNode) {
					//println(fmt.Sprintf("Node: %s is already a child of: %s", aVisitableChildNode.node.id, currentNode.node.id))
				} else {
					//println(fmt.Sprintf("Node: %s adding child: %s", currentNode.node.id, aVisitableChildNode.node.id))
					currentNode.nodes = append(currentNode.nodes, aVisitableChildNode)
				}
			}
		}

		for _, aVisitableChildNode := range currentNode.nodes {
			if _, ok := visited[aVisitableChildNode.node.id]; !ok {
				me.build(aVisitableChildNode, visited, edgeMap, queue)
			}

		}

	}
	//queue.Dequeue()

	return currentNode

}

func (me *BFS) prepare(startingCell *Cell) *Node {

	visited := make(map[string]*Node)
	edgeMap := make(map[string]*Node)
	queue := utility.NewSimpleQueue()

	var currentNode *Node
	currentNode = &Node{
		nodes: make([]*Node, 0),
		node:  startingCell,
	}

	me.build(currentNode, visited, edgeMap, queue)
	me.tree = currentNode
	return currentNode
}

func isAppropriate(sourceWeight interface{}, targetItem *utility.WeightedItem) bool {

	appropriate := false
	sourceID := sourceWeight.(string)
	if targetItem.Weight == nil {
		appropriate = false
	} else {
		targetSourceID := targetItem.Weight.(string)
		appropriate = sourceID == targetSourceID
	}

	return appropriate
}

func sortForDequeue(items []interface{}) []interface{} {
	sort.Slice(items, func(i, j int) bool {
		return items[i].(*utility.WeightedItem).Weight.(string) > items[j].(*utility.WeightedItem).Weight.(string)
	})
	return items
}

func (me *BFS) shortest(end *Cell) *utility.SimpleStack {

	paths := Paths{paths: make([]*utility.SimpleStack, 0)}
	currentPath := utility.NewSimpleStack()

	visited := make(map[string]*Node)
	queue := utility.NewWeightedQueue()

	var weightedIDOfInterest string
	{
		node := me.tree
		visited[node.node.id] = node
		weightedIDOfInterest = node.node.id
		queue.Enqueue(node, weightedIDOfInterest)
	}

	for queue.HasMore() {
		var node *Node
		/*
			if dequeue is nil then we know that we have reached the furthest edge of the tree.
		*/

		if ok, temp := queue.DequeueUsing(weightedIDOfInterest, isAppropriate, sortForDequeue); ok {
			node = temp.(*utility.WeightedItem).Item.(*Node)
		}

		//if node == nil {
		//	if ok, temp := queue.Dequeue(); ok {
		//		node = temp.(*utility.WeightedItem).Item.(*Node)
		//	}
		//
		//} else {
		//	if ok, temp := queue.DequeueUsing(node.node.id, isAppropriate); ok {
		//		node = temp.(*utility.WeightedItem).Item.(*Node)
		//	}
		//
		//}
		if node == nil {
			println("queue has items but was unable to draw")
			currentPath.Pop()
		} else {
			currentPath.Push(node.node)

			if node.node == end {
				//- if at the end, copy off the stack.
				paths.add(currentPath.Copy())

				//- pop an item off of current stack and search queue based upon that weight
				currentPath.Pop()

			} else {
				//nodeList := me.matrix.nodeListMap[node.id]
				nodeList := node.nodes
				if nodeList != nil {
					//nodes := sortNodesDescending(nodeList)
					nodes := nodeList
					for _, aChildNode := range nodes {
						if aChildNode != nil && me.isNavigable(node.node, aChildNode.node) && !currentPath.Contains(aChildNode.node) {

							if _, ok := visited[aChildNode.node.id]; !ok {
								visited[aChildNode.node.id] = aChildNode
								queue.Enqueue(aChildNode, node.node.id)

							}
						}
					}
				}
			}
		}

		weightedIDOfInterest = currentPath.Peek().(*Cell).id

	}

	return paths.shortest()
}

func sortCellsAscending(cells []*Cell) []*Cell {
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].z < cells[j].z
	})
	return cells
}

func sortCellsDescending(cells []*Cell) []*Cell {
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].z > cells[j].z
	})
	return cells
}

func sortNodesDescending(nodes []*Node) []*Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].node.z > nodes[j].node.z
	})
	return nodes
}

func sortNodesAscending(nodes []*Node) []*Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].node.z < nodes[j].node.z
	})
	return nodes
}
