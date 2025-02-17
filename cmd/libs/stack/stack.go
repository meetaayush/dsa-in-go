package libs

import (
	"errors"
)

type Stack struct {
	elements []int
	top      int
	size     int
}

func (st *Stack) IsEmpty() bool {
	if st.top == -1 {
		return true
	}
	return false
}

func (st *Stack) isFull() bool {
	return st.top >= st.size
}

func NewStack(size int) *Stack {
	elements := make([]int, size)
	for i := range elements {
		elements[i] = -1
	}
	return &Stack{
		elements: elements,
		top:      -1,
		size:     size,
	}
}

func (st *Stack) Push(x int) error {
	if st.isFull() {
		return errors.New("stack is full")
	}
	st.top = st.top + 1
	st.elements[st.top] = x

	return nil
}

func (st *Stack) Pop() error {
	if st.IsEmpty() == true {
		return errors.New("stack is empty")
	}

	st.top = st.top - 1

	return nil
}

func (st *Stack) Peek() (int, error) {
	if st.IsEmpty() == true {
		return 0, errors.New("stack is empty")
	}

	return st.elements[st.top], nil
}
