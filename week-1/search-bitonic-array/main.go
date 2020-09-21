package main

import (
	"log"
	"sort"
)

func biggest(myset []int) int {
        first := 0
        last := len(myset) - 1
	var middle,biggest int

        for first <= last {
                middle = first + (last-first)/2

                //fmt.Printf("first: %d last: %d middle: %d biggest: %d \n",first,last,middle, biggest)

		if myset[middle+1] < myset[middle] {
			return middle
		}

		if myset[middle+1] > myset[middle] && myset[middle+1] > myset[biggest] {
			first = middle + 1
			biggest = middle + 1
			continue
		} 

                return biggest
        }

	return -1
}

func binarysearch(myset []int, numberToFind int) bool {
	first := 0
        last := len(myset) - 1
	var middle int

	sort.Ints(myset)

        //fmt.Println(myset)

        for first <= last {
                middle = first + (last-first)/2

                //fmt.Printf("first: %d last: %d middle: %d \n",first,last,middle)

                if numberToFind > myset[middle] {
                        first = middle + 1
                        continue
                }

                if numberToFind < myset[middle] {
                        last = middle - 1
                        continue
                }

                log.Printf("Found %d in key %d \n",numberToFind,middle)
                return true
        }
	
	return false

}

func main() {
	myset := []int{-3, 9, 8, 20, 17, 5, 1}

	numberToFind := -3

	big := biggest(myset)

	firstSearch := binarysearch(myset[0:big],numberToFind)

	if !firstSearch {
		last := len(myset)-1
		secondSearch := binarysearch(myset[big:last],numberToFind)
		if !secondSearch {
			log.Printf("Number %d not found \n",numberToFind)
		}
	}

}

