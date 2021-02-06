package sudoku_test

import (
	"fmt"
	"testing"

	"github.com/timgluz/sudoku/pkg/sudoku"
)

func TestSudokuSolve(t *testing.T) {
	var examples = []struct {
		grid     string
		expected string
	}{
		{
			"003020600900305001001806400008102900700000008006708200002609500800203009005010300",
			"483921657967345821251876493548132976729564138136798245372689514814253769695417382",
		},
		{
			"52...6.........7.13...........4..8..6......5...........418.........3..2...87.....",
			"527316489896542731314987562172453896689271354453698217941825673765134928238769145",
		},
		{
			"6.....8.3.4.7.................5.4.7.3..2.....1.6.......2.....5.....8.6......1....",
			"617459823248736915539128467982564371374291586156873294823647159791385642465912738",
		},
	}

	for i, tt := range examples {
		testname := fmt.Sprintf("example.%v", i+1)
		t.Run(testname, func(t *testing.T) {
			solver := sudoku.NewSudoku()
			err := solver.Solve(tt.grid)
			if err != nil {
				t.Errorf("got error, but expected solution for %v", tt.grid)
			}

			if tt.expected != solver.String() {
				t.Errorf("got %v, expected %v", solver.String(), tt.expected)
			}
		})
	}
}
