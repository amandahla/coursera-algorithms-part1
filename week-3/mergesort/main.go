package main

// Reference:
// https://austingwalters.com/merge-sort-in-go-golang/

import "log"

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	mid := len(input)/2

	return Merge(MergeSort(input[:mid]),MergeSort(input[mid:]))
}

func Merge(left,right []int) []int {
	var i,j int
	size := len(left) + len(right)
	result := make([]int, size, size)

	endleft := len(left)-1
	endright := len(right)-1

	for k:=0; k<size; k++ {
		// left is over but we still have elements on the right side
		if i > endleft && j <= endright {
			result[k] = right[j]
			j++
			continue
		}

		// right is over but we still have elements on the left side
		if j > endright && i <=endleft {
			result[k] = left[i]
			i++
			continue
		}

		// compare elements
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		}else{
			result[k] = right[j]
			j++
		}
	}
	
	return result
}

func main(){

	myset := []int{7, 4, 2, 1, 10, 8, 5, 3, 6, 9}

        log.Print(myset)

	log.Print(MergeSort(myset))
}
