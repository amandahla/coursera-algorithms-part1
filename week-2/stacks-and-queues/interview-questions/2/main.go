package main

import (
	"fmt"
	"log"
)

type node struct {
	value int
	next  *node
}

type stack struct {
	first *node
	total int
}

var max stack

func (s *stack) push(v int) {
	n := node{value: v}
	n.next = s.first
	s.first = &n

	maxNumber, _ := max.top()
	if v > maxNumber {
		max.push(v)
	}

	s.total += 1
}

func (s *stack) pop() (node, error) {
	if s.first != nil {
		old := s.first
		s.first = old.next

		maxNumber, _ := max.top()
		if maxNumber == old.value {
			max.pop()
		}

		s.total -= 1

		return *old, nil
	}

	return node{}, fmt.Errorf("Stack is empty")
}

func (s *stack) top() (int, error) {
	if s.first != nil {
		return s.first.value, nil
	}

	return 0, fmt.Errorf("Stack is empty")
}

func main() {
	mystack := stack{}
	mystack.push(1)
	mystack.push(25)
	mystack.push(100)
	mystack.push(30)
	mystack.push(2)
	mystack.push(400)

	maxNumber, err := max.top()
	if err != nil {
		log.Print(err)
	}
	log.Printf("My stack has %d items and the maximum is %d\n", mystack.total, maxNumber)

	log.Printf("Pop 2 items")

	_, err = mystack.pop()
	if err != nil {
		log.Print(err)
	}
	_, err = mystack.pop()
	if err != nil {
		log.Print(err)
	}

	maxNumber, err = max.top()
	if err != nil {
		log.Print(err)
	}
	log.Printf("My stack has %d items and the maximum is %d\n", mystack.total, maxNumber)
}
