package main

import "fmt"

type Stack struct {
	Items []Flight
}

func (s *Stack) Pop() Flight {
	if !s.IsEmpty() {
		last := s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
		return last
	}
	return Flight{}
}

func (s *Stack) Push(f Flight) {
	s.Items = append(s.Items, f)
}

func (s *Stack) Peek() Flight {
	if !s.IsEmpty() {
		return s.Items[len(s.Items)-1]
	}
	return Flight{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func goChall5() {
	fmt.Println("Go Stack Implementation")
}
