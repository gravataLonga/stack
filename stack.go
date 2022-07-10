package stack

// Stack save every element as stack
type Stack struct {
	stack []interface{}
}

// New create a new stack
func New() *Stack {
	return &Stack{}
}

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.stack) <= 0
}

// Push new item to top of stack
func (s *Stack) Push(e interface{}) {
	s.stack = append(s.stack, e)
}

// Pop first element of stack
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return e, true
}
