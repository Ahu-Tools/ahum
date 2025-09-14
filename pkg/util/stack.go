package util

// Stack is a generic stack implementation.
// It is a last-in, first-out (LIFO) data structure.
type Stack[T any] struct {
	elements []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

// Pop removes and returns the top element of the stack.
// It returns the element and a boolean indicating if the operation was successful.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T // Return the zero value for the type if stack is empty
		return zero, false
	}

	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index] // Re-slice to remove the last element
	return element, true
}

// Peek returns the top element of the stack without removing it.
// It returns the element and a boolean indicating if the operation was successful.
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) MustPeek() T {
	if s.IsEmpty() {
		panic("MustPeek: attempted to peek at an empty stack")
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
