package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"unicode"

	"github.com/timgluz/sudoku/pkg/sudoku"
)

func main() {
	var showTable = flag.Bool("human", false, "prints soulution as grid, more human friendly")
	flag.Parse()

	grid := cleanGrid(flag.Arg(0))
	err := validateGrid(grid)
	if err != nil {
		fmt.Println("Invalid input:", err)
		os.Exit(1)
	}

	solver := sudoku.NewSudoku()

	err = solver.Solve(grid)
	if err != nil {
		panic(fmt.Errorf("Failed to find solution: %v", err))
	}

	if *showTable {
		solver.Display()
	} else {
		fmt.Println(solver.String())
	}
}

func cleanGrid(grid string) string {
	var res string

	// remove all non digits
	for _, ch := range grid {
		if unicode.IsSpace(ch) || unicode.IsControl(ch) {
			continue
		}

		res += string(ch)
	}

	return res
}

func validateGrid(grid string) error {
	if len(grid) != 81 {
		return errors.New("grid must have 81 elements")
	}

	return nil
}
