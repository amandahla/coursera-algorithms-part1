package main

import (
	"fmt"
	"sort"
)


func main() {
	myset := []int{10,10,-20,40,-80,40,10,10,10,10,-10,-10}
	sort.Ints(myset)
	result := [][]int{}

	fmt.Println(myset)

	for i := 0; i < len(myset); i++ {
		mymap := make(map[int]int)

		if i != 0 && myset[i] == myset[i - 1] {
			 continue
		}

		for j := i+1; j < len(myset); j++ {
			c := -myset[i] - myset[j]
			if _, ok := mymap[c]; ok {
				result = append(result,[]int{myset[i],myset[j],c})
				for j+1 < len(myset) && myset[j+1] == myset[j] {
					j+=1
				}
			}
			mymap[myset[j]] = 0
		}
	}
	fmt.Println(result)
}
