package main

import (
	"log"
)

var grid [][]bool
var s SetUnion
var idxs [][]int
var size int

// Creates a grid with false for closed and true for opened
// Creates a idxs with a number for each cell
// Creates a s to support union operations where we link idxs as each cell opens
func percolation(n int) {
	size = n
	grid = make([][]bool, size)
	idxs = make([][]int, size)
	idx := 0

	s = SetUnion{}
	s.create(size * size)

	for i := 0; i < size; i++ {
		grid[i] = make([]bool, size)
		idxs[i] = make([]int, size)
		for j := 0; j < size; j++ {
			grid[i][j] = false
			idxs[i][j] = idx
			idx++
		}
	}
}

func open(row, col int) {
	log.Printf("opening cell: %d,%d \n", row, col)
	grid[row][col] = true
	cellidx := idxs[row][col]
	// top
	if row > 0 && isOpen(row-1, col) {
		cellidxTop := idxs[row-1][col]
		s.union(cellidx, cellidxTop)
		log.Printf("top is open: %d,%d \n", row-1, col)
	}
	// down
	limit := size - 1
	if row < limit && isOpen(row+1, col) {
		cellidxDown := idxs[row+1][col]
		s.union(cellidx, cellidxDown)
		log.Printf("down is open: %d,%d \n", row+1, col)
	}
	// right
	if col < limit && isOpen(row, col+1) {
		cellidxRight := idxs[row][col+1]
		s.union(cellidx, cellidxRight)
		log.Printf("right is open: %d,%d \n", row, col+1)
	}
	// left
	if col > 0 && isOpen(row, col-1) {
		cellidxLeft := idxs[row][col-1]
		s.union(cellidx, cellidxLeft)
		log.Printf("left is open: %d,%d \n", row, col-1)
	}
}

func isOpen(row, col int) bool {
	return grid[row][col]
}

/*
A full site is an open site that can be connected to an open site in the top row
via a chain of neighboring (left, right, up, down) open sites. If there is a full
site in the bottom row, then we say that the system percolates.
*/
func isFull(row, col int) bool {
	if row == 0 { // top
		return true
	}

	if isOpen(row, col) {
		idx := idxs[row][col]
		// for each item on top row, check if its connected
		for i := 0; i < size; i++ {
			idxtop := idxs[0][i]
			log.Printf("testing if item %d is connected with %d \n", idx, idxtop)
			if s.isConnected(idx, idxtop) {
				return true
			}
		}
	}
	return false
}

func numberOfOpenSites() int {
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				total++
			}
		}
	}

	return total
}

func percolates() bool {
	lastrow := len(grid) - 1

	// for each item on bottom row, check if isFull (connected with a top row cell)
	for i := 0; i < len(grid[lastrow]); i++ {
		if isFull(lastrow, i) {
			return true
		}
	}

	return false
}

func main() {
	percolation(1000)
	open(0, 1)
	open(1, 1)
	open(1, 2)
	open(2, 2)
	open(2, 3)
	open(3, 3)
	for i := 0; i < 1000; i++ {
		open(i, 0)
	}
	log.Println(numberOfOpenSites())
	if percolates() {
		log.Println("Percolates!")
	}
}
