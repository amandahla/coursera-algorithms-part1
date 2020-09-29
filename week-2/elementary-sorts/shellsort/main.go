package main

import (
	"log"
)

func main() {
	myset := []int{7, 4, 2, 1, 10, 8, 5, 3, 6, 9}

	log.Print(myset)

	h := 1
	for h < len(myset)/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < len(myset); i++ {
			for j := i; j >= h && myset[j] < myset[j-h]; j = j - h {
				tmp := myset[j]
				myset[j] = myset[j-h]
				myset[j-h] = tmp
			}
		}

		h = h / 3
	}

	log.Print(myset)
}
