package stack

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	length := len(s.items)
	if length == 0 {
		return nil
	}
	item := s.items[length-1]
	s.items = s.items[0 : length-1]
	return item
}

func (s *Stack) Top() interface{} {
	length := len(s.items)
	if length == 0 {
		return nil
	}
	return s.items[length-1]
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}
