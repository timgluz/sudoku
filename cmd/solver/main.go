package main

import (
	"fmt"
	"github.com/timgluz/sudoku/pkg/sudoku"
)

func main() {
	// TODO: read from args
	//grid := "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	grid := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"

	solver := sudoku.NewSudoku()

	err := solver.Solve(grid)
	if err != nil {
		panic(fmt.Errorf("Failed to find solution: %v", err))
	}

	solver.Display()
}
