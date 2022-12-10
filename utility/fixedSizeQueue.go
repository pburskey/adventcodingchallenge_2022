package utility

type FixedSizeStringQueue struct {
	queue []string
	size  int
}

func NewFixedSizeStringQueue(aSize int) *FixedSizeStringQueue {
	queue := &FixedSizeStringQueue{queue: make([]string, aSize), size: aSize}
	return queue
}

func (s *FixedSizeStringQueue) Enqueue(aString string) {
	for i := 0; i < s.size-1; i++ {
		s.queue[i] = s.queue[i+1]
	}
	s.queue[s.size-1] = aString
}

func (s *FixedSizeStringQueue) PeakAll() []string {
	return s.queue
}
