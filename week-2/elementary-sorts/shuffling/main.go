package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	myset := []int{7, 4, 2, 1, 10, 8, 5, 3, 6, 9}

	log.Print(myset)

	for i := 0; i < len(myset); i++ {
		tmp := myset[i]
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(i+1)
		myset[i] = myset[r]
		myset[r] = tmp
	}

	log.Print(myset)
}
