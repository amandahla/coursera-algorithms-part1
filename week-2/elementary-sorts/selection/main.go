package main

import (
	"log"
)

func main() {
	myset := []int{7,4,2,1,10,8,5,3,6,9}

	log.Print(myset)

	for i := 0; i < len(myset); i++ {
		min := i
		for j := i + 1; j < len(myset); j++ {
			if myset[j] < myset[min] {
				min = j
			}
		}
		tmp := myset[i]
		myset[i] = myset[min]
		myset[min] = tmp
	}

	log.Print(myset)
}
