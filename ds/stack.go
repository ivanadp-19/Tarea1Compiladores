package ds

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	n := len(s.items)
	v := s.items[n-1]
	s.items = s.items[:n-1]
	return v
}

func (s *Stack) Print() {
	fmt.Println(s.items)
}
