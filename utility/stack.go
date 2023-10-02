package utility

type SimpleStack struct {
	stack []interface{}
}

func NewSimpleStack() *SimpleStack {
	stack := &SimpleStack{stack: make([]interface{}, 0)}
	return stack
}

func (s *SimpleStack) Reverse() []interface{} {
	inputLen := len(s.stack)
	inputMid := inputLen / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		s.stack[i], s.stack[j] = s.stack[j], s.stack[i]
	}
	return s.stack
}

func (s *SimpleStack) Push(item interface{}) {
	s.stack = append(s.stack, item)
}

func (s *SimpleStack) HasMore() bool {
	return len(s.stack) > 0
}

func (s *SimpleStack) Pop() interface{} {
	index := len(s.stack) - 1
	item := s.stack[index]
	s.stack = s.stack[:index]
	return item
}

func (s *SimpleStack) Peek() interface{} {
	index := len(s.stack) - 1
	item := s.stack[index]
	return item
}
