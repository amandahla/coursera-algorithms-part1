package main

import "log"

var myset = make([]int, 10)
var myweight = make([]int, 10)
var first = 0
var last = len(myset) - 1
var printMessage = true

func main() {
	for i := 0; i < len(myset); i++ {
		myset[i] = i
	}

	union(0, 1)
	union(1, 2)
	union(2, 3)
	union(3, 4)
	union(4, 5)
	union(5, 6)
	union(6, 7)
	union(7, 8)
	union(8, 9)
	union(0, 1)
	union(1, 2)
	union(2, 3)
	union(3, 4)
}

func isEveryoneTogether() {
	if isConnected(first, last) && printMessage {
		printMessage = false
		log.Printf("All friends here!")
	}
}

func union(a, b int) {
	log.Printf("%d and %d will be friends", a, b)

	roota := root(a)
	rootb := root(b)

	if roota == rootb {
		isEveryoneTogether()
		return
	}

	if myweight[roota] < myweight[rootb] {
		myset[roota] = rootb
		myweight[rootb] += myweight[roota]
	} else {
		myset[rootb] = roota
		myweight[roota] += myweight[rootb]
	}

	isEveryoneTogether()
}

func isConnected(a, b int) bool {
	return root(a) == root(b)
}

func root(i int) int {
	for i != myset[i] {
		// flattening the tree
		// change my actual root to be the root of my root
		myroot := myset[i]
		myset[i] = myset[myroot]
		i = myset[i]
	}

	return i
}
