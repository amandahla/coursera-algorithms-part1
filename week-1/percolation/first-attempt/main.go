package main

import "fmt"

var table [][]bool
var unions []SetUnion

func percolation(size int) {
	table = make([][]bool, size)
	for i := 0; i < size; i++ {
		table[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			table[i][j] = false
		}
	}
}

func open(row, col int) {
	table[row][col] = true
}

func isOpen(row, col int) bool {
	return table[row][col]
}

func isFull(row, col int) bool {
	if row == 0 {
		return true
	}

	if isOpen(row, col) {
		if isOpen(row-1, col) {
			return true
		}

		for i := 0; i < len(table[i]); i++ {
			if isOpen(row, i) {
				if isOpen(row-1, i) {
					return true
				}
			}
		}
	}
	return false
}

func numberOfOpenSites() int {
	total := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] {
				total++
			}
		}
	}

	return total
}

func percolates() bool {
	lastrow := len(table) - 1

	for i := 0; i < len(table[lastrow]); i++ {
		if isFull(lastrow, i) {
			return true
		}
	}

	return false
}

func main() {
	percolation(4)
	open(0, 1)
	open(1, 1)
	open(1, 2)
	open(2, 2)
	open(2, 3)
	open(3, 3)
	fmt.Println(numberOfOpenSites())
	fmt.Println(percolates())
}
