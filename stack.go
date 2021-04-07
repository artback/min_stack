package main

import (
	"math"
)

func main() {
	s := Stack{}
	s.push(5)
	s.push(2)
	for {
		value := s.pop()
		if value == nil {
			break
		}
	}

}

type element struct {
	value int
	min   int
	next *element
}
type Stack struct {
	head *element
}

func (s *Stack) push(val int) {
	currMin := math.MaxInt32
	if s.head != nil {
		currMin = s.head.min
	}
	element := &element{value: val, min: currMin}
	element.next = s.head
	if currMin > val {
		element.min = val
	}
	s.head = element
}
func (s *Stack) pop() *int {
	if s.head == nil {
		return nil
	}
	val := s.head
	s.head = s.head.next
	return &val.value
}

func (s *Stack) peek() *int {
	if s.head == nil {
		return nil
	}
	return &s.head.value
}
func (s *Stack) min() *int {
	if s.head == nil {
		return nil
	}
	return &s.head.min
}
