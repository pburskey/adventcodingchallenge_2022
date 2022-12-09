package utility

type SimpleStringQueue struct {
	queue []string
}

func NewSimpleStringQueue() *SimpleStringQueue {
	queue := &SimpleStringQueue{queue: make([]string, 0)}
	return queue
}

func (s *SimpleStringQueue) Enqueue(aString string) {
	s.queue = append(s.queue, aString)
}

func (s *SimpleStringQueue) Dequeue() string {
	element := s.queue[0] // The first element is the one to be dequeued.
	if len(s.queue) == 1 {
		s.queue = []string{}
	} else {
		s.queue = s.queue[1:]
	}

	return element
}
