package utility

type SimpleQueue struct {
	queue []interface{}
}

func NewSimpleQueue() *SimpleQueue {
	queue := &SimpleQueue{queue: make([]interface{}, 0)}
	return queue
}

func (s *SimpleQueue) Enqueue(item interface{}) {
	s.queue = append(s.queue, item)
}

func (s *SimpleQueue) Dequeue() (ok bool, element interface{}) {
	if len(s.queue) > 0 {
		element = s.queue[0] // The first element is the one to be dequeued.
		if len(s.queue) == 1 {
			s.queue = []interface{}{}
		} else {
			s.queue = s.queue[1:]
		}
		return true, element
	}
	return false, nil

}

func (s *SimpleQueue) IsEmpty() bool {
	return len(s.queue) == 0
}
