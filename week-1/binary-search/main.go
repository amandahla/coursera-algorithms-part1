package main

import "fmt"

func main() {
	var middle int

	myset := []int{-3, 8, 9, 20, 30, 45, 104}

	numberToFind := 8

	first := 0
	last := len(myset) - 1
	
	fmt.Println(myset)

	for first <= last {
		middle = first + (last-first)/2

		fmt.Printf("first: %d last: %d middle: %d \n",first,last,middle)

		if numberToFind > myset[middle] {
			first = middle + 1
			continue
		}

		if numberToFind < myset[middle] {
			last = middle - 1
			continue
		}

		fmt.Printf("Found %d in key %d \n",numberToFind,middle)
		return
	}
}
