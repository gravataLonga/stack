package stack

import "testing"

func TestStack_IsEmpty(t *testing.T) {
	s := New()

	if !s.IsEmpty() {
		t.Fatalf("Stack is not empty")
	}
}

func TestStack_Push(t *testing.T) {
	s := New()

	s.Push(true)

	if s.IsEmpty() {
		t.Fatalf("Stack is empty")
	}
}

func TestStack_Pop(t *testing.T) {
	s := New()
	s.Push("hello")
	s.Pop()

	if !s.IsEmpty() {
		t.Fatalf("Stack is not empty")
	}
}

func TestStack_PopGetSameElement(t *testing.T) {
	s := New()
	a := "hello"
	s.Push(a)
	p, _ := s.Pop()

	if !s.IsEmpty() {
		t.Fatalf("Stack is not empty")
	}

	if p != a {
		t.Fatalf("element popped isn't same as created")
	}
}

func TestStack_PopOnEmptyStack(t *testing.T) {
	s := New()
	p, ok := s.Pop()

	if !s.IsEmpty() {
		t.Fatalf("Stack is not empty")
	}

	if p != nil {
		t.Errorf("Popped element isn't nil")
	}

	if ok {
		t.Errorf("Boolean isn't true")
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New()
		s.Push("hello")
		s.Push(20000)
		s.Push([]struct{ i int }{{10}})
		s.Push(false)
		s.IsEmpty()
		s.Pop()
		s.Pop()
		s.Pop()
		s.Pop()
	}
}

var numElement = 100000

func BenchmarkNew2(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {

		for i <= numElement {
			s.Push("hello world")
			i++
		}

		for i <= numElement {
			s.Pop()
			i++
		}

	}
}
