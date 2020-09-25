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

func (s *stack) pop() (*node, error) {
        if s.first != nil {
                old := s.first
                s.first = old.next

                maxNumber, _ := max.top()
                if maxNumber == old.value {
                        max.pop()
                }

                s.total -= 1

                return old, nil
        }

        return nil, fmt.Errorf("Stack is empty")
}

func (s *stack) top() (int, error) {
        if s.first != nil {
                return s.first.value, nil
        }

        return 0, fmt.Errorf("Stack is empty")
}

type queue struct {
	stackin *stack
	stackout *stack
}

func (q *queue) add(v int){
	q.stackin.push(v)
}

func (q *queue) remove() *node {
	if q.stackout.total == 0 {
		n,_ := q.stackin.pop()
		for n != nil {
			q.stackout.push(n.value)
			n,_ = q.stackin.pop()
		}
	}

	n,_ := q.stackout.pop()
	return n
}


func main() {
	myqueue := queue{stackin: &stack{}, stackout: &stack{}}
	myqueue.add(1)
	myqueue.add(2)
	myqueue.add(3)
	log.Printf("First in line: %d",myqueue.remove().value)
}

