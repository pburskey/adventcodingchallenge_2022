package utility

type SimpleStringStack struct {
	stack []string
}

func NewSimpleStringStack() *SimpleStringStack {
	stack := &SimpleStringStack{stack: make([]string, 0)}
	return stack
}

func (s *SimpleStringStack) Reverse() []string {
	return s.stack
}

func (s *SimpleStringStack) Push(aString string) {
	s.stack = append(s.stack, aString)
}

func (s *SimpleStringStack) HasMore() bool {
	return len(s.stack) > 0
}

func (s *SimpleStringStack) Pop() string {
	index := len(s.stack) - 1
	aString := s.stack[index]
	s.stack = s.stack[:index]
	return aString
}

func (s *SimpleStringStack) Peek() string {
	index := len(s.stack) - 1
	aString := s.stack[index]
	return aString
}
