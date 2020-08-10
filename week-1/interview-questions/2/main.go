package main

import "log"

var myset = make([]int, 13)
var myweight = make([]int, 13)

func main() {
	for i := 0; i < len(myset); i++ {
		myset[i] = i
	}

	union(0, 1)
	union(1, 2)
	union(3, 4)
	union(0, 10)

	log.Println(find(0))
}

func union(a, b int) {
	roota := root(a)
	rootb := root(b)

	if roota == rootb {
		return
	}

	if myweight[roota] < myweight[rootb] {
		myset[roota] = rootb
		myweight[rootb] += myweight[roota]
		return
	}

	myset[rootb] = roota
	myweight[roota] += myweight[rootb]
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

func find(n int) int {
	biggest := n
	rootN := root(n)
	for i := 0; i < len(myset); i++ {
		if myset[i] == rootN && i > biggest {
			biggest = i
		}
	}

	return biggest
}
