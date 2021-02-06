package sudoku

import "errors"

type sudoku_state map[string]string

func DFSSearch(sudoku Sudoku) (Sudoku, error) {
	if sudoku.IsSolved() {
		return sudoku, nil
	}

	// pick the unfilled square s with the fewest possibilities
	next_square, ok := leastRemainingSquare(sudoku.State)
	if !ok {
		return Sudoku{}, errors.New("No remaing Squares")
	}

	for _, val := range sudoku.State[next_square] {
		tmp := CloneSudoku(sudoku)
		tmp.assign(next_square, val)
		solved_sudoku, err := DFSSearch(tmp)
		if err == nil {
			return solved_sudoku, nil
		}
	}

	return Sudoku{}, errors.New("No solution")
}

// returns square with smallest options or the first with 2options
func leastRemainingSquare(state map[string]string) (string, bool) {
	var minSquare string
	minLen := 9
	found := false

	for k, v := range state {
		if len(v) < minLen {
			minLen = len(v)
			minSquare = k
			found = true
		}

		// early stop
		if minLen == 2 {
			break
		}
	}

	return minSquare, found
}
