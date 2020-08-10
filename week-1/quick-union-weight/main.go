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
	union(0, 3)
	union(4, 5)
	union(6, 7)
	union(8, 9)
	union(9, 10)
	union(5, 7)
	union(7, 10)
	union(11, 12)
	log.Println(isConnected(2, 4))
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
		i = myset[i]
	}

	return i
}
