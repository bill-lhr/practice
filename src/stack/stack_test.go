package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	top := stack.Top()
	if top != 4 {
		t.Errorf("stack.Top() err, got: %v", top)
	}

	pop := stack.Pop()
	if pop != 4 {
		t.Errorf("stack.Pop() err, got: %v", pop)
	}

	pop = stack.Pop()
	if pop != 3 {
		t.Errorf("stack.Pop() err, got: %v", pop)
	}

	stack.Push(5)
	top = stack.Top()
	if top != 5 {
		t.Errorf("stack.Top() err, got: %v", top)
	}

	stack.Pop()
	stack.Pop()
	stack.Pop()
	if !stack.Empty() {
		t.Errorf("stack.Empty() err")
	}
}
