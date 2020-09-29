package main

import (
	"log"
)

func main() {
	myset := []int{7, 4, 2, 1, 10, 8, 5, 3, 6, 9}

	log.Print(myset)

	for i := 0; i < len(myset); i++ {
		for j := i; j > 0; j-- {
			if myset[j] > myset[j-1] {
				break
			}

			tmp := myset[j-1]
			myset[j-1] = myset[j]
			myset[j] = tmp
		}
	}

	log.Print(myset)
}
