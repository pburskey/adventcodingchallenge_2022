package utility

type WeightedItem struct {
	Item   interface{}
	Weight interface{}
}

type WeightedQueue struct {
	queue []*WeightedItem
}

func NewWeightedQueue() *WeightedQueue {
	queue := &WeightedQueue{queue: make([]*WeightedItem, 0)}
	return queue
}

func (me *WeightedQueue) Enqueue(item interface{}, weight interface{}) {
	me.queue = append(me.queue, &WeightedItem{
		Item:   item,
		Weight: weight,
	})
}

func (me *WeightedQueue) Dequeue() (ok bool, element interface{}) {
	if len(me.queue) > 0 {
		element = me.queue[0] // The first element is the one to be dequeued.
		if len(me.queue) == 1 {
			me.queue = []*WeightedItem{}
		} else {
			me.queue = me.queue[1:]
		}
		return true, element
	}
	return false, nil
}

func (me *WeightedQueue) DequeueUsing(weigthToConsider interface{}, isAppropriate func(targetWeight interface{}, queuedItem *WeightedItem) bool, sortNodes func([]interface{}) []interface{}) (ok bool, element interface{}) {
	if len(me.queue) > 0 {

		values := make([]interface{}, 0)
		index := make(map[int]*WeightedItem)
		/*
			sort
		*/
		for x, element := range me.queue {
			appropriate := isAppropriate(weigthToConsider, element)
			if appropriate {
				values = append(values, element)
				index[x] = element
			}
		}

		var element interface{}
		sortNodes(values)
		if values != nil && len(values) > 0 {
			element = values[0]
		}
		if element == nil {
			return false, nil
		}

		x := 0
		for key, value := range index {
			if value == element {
				x = key
				break
			}
		}

		if element != nil {
			if len(me.queue) == 1 {
				me.queue = []*WeightedItem{}
			} else if len(me.queue) == x+1 {
				me.queue = append(me.queue[:x])
			} else {
				me.queue = append(me.queue[:x], me.queue[x+1])
			}
			return true, element
		}

		//
		//for x, element := range me.queue {
		//	appropriate := isAppropriate(weigthToConsider, element)
		//	if appropriate {
		//		if len(me.queue) == 1 {
		//			me.queue = []*WeightedItem{}
		//		} else if len(me.queue) == x+1 {
		//			me.queue = append(me.queue[:x])
		//		} else {
		//			me.queue = append(me.queue[:x], me.queue[x+1])
		//		}
		//		return true, element
		//	}
		//}

	}
	return false, nil
}

func (me *WeightedQueue) IsEmpty() bool {
	return len(me.queue) == 0
}

func (me *WeightedQueue) HasMore() bool {
	return !me.IsEmpty()
}
