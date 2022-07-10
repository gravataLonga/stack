package stack

// stack save every element as stack
type stack struct {
	stack []interface{}
}

// New create a new stack
func New() *stack {
	return &stack{}
}

// IsEmpty check if stack is empty
func (s *stack) IsEmpty() bool {
	return len(s.stack) <= 0
}

// Push new item to top of stack
func (s *stack) Push(e interface{}) {
	s.stack = append(s.stack, e)
}

// Pop first element of stack
func (s *stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return e, true
}

func (s *stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	return s.stack[len(s.stack)-1], true
}
