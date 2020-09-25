package main

import "log"

var myset []int

func resize(size int) {
	tmp := make([]int,size)
	limit := len(myset)
	if size < limit {
		limit = size
	}
	myset = myset[:limit]	
	copy(tmp,myset)
	myset = tmp
}

func describe() {
	log.Printf("myset: %v len: %d cap: %d\n", myset, len(myset), cap(myset))
}

func main() {
	myset = append(myset,[]int{1,2,3,4,5}...)
	describe()
	resize(10)
	describe()
	resize(3)
	describe()
}
