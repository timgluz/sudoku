package sudoku

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	DIGITS       string = "123456789"
	COLUMNS      string = "123456789"
	ROWS         string = "ABCDEFGHI"
	SQUARE_COUNT int    = 81
)

// Peers & Units
var (
	Peers map[string][]string   = build_peers(ROWS, COLUMNS)
	Units map[string][][]string = build_units(ROWS, COLUMNS)
)

type Sudoku struct {
	State map[string]string
}

func NewSudoku() Sudoku {
	coords := crossProduct(ROWS, COLUMNS)
	state := make(map[string]string)
	for _, coord := range coords {
		state[coord] = DIGITS
	}

	return Sudoku{state}
}

func CloneSudoku(sudoku Sudoku) Sudoku {
	tmp := NewSudoku()

	for k, v := range sudoku.State {
		tmp.State[k] = v
	}

	return tmp
}

func (s *Sudoku) Solve(grid string) error {
	// constraint propagate
	squares := crossProduct(ROWS, COLUMNS)

	// FIX: we are visiting same squares many times
	for i, item := range grid {
		if strings.ContainsRune(DIGITS, item) {
			err := s.assign(squares[i], item)
			if err != nil {
				return fmt.Errorf("Failure during CP at coord: %v, error: %v:", squares[i], err)
			}
		}
	}

	solution, err := DFSSearch(*s)
	if err != nil {
		return fmt.Errorf("No solution: %v", err)
	}

	s.State = solution.State
	return nil
}

func (s *Sudoku) IsSolved() bool {
	for _, v := range s.State {
		if len(v) != 1 {
			return false
		}
	}

	return true
}

func (s Sudoku) Display() {
	width := 1
	for _, values := range s.State {
		if len(values) > width {
			width = len(values)
		}

		if width == 9 {
			break
		}
	}

	cell_format := "%" + strconv.Itoa(width+1) + "s"
	divider := strings.Repeat("-", 3*(width+1))
	horizontal_divider := fmt.Sprintf("%v+%v+%v", divider, divider, divider)

	for x, row := range ROWS {
		if x == 3 || x == 6 {
			fmt.Println(horizontal_divider)
		}

		line := ""
		for y, col := range COLUMNS {
			if y == 3 || y == 6 {
				line += "|"
			}

			coord := string(row) + string(col)
			line += fmt.Sprintf(cell_format, s.State[coord])
		}
		fmt.Println(line)
	}
}

// returns 1line string presentation of Sudoku board
func (s Sudoku) String() string {
	var line string

	for _, row := range ROWS {
		for _, col := range COLUMNS {
			coord := string(row) + string(col)

			if len(s.State[coord]) == 1 {
				line += s.State[coord]
			} else {
				line += "0"
			}
		}
	}

	return line
}

// -- Constraint Propagation
// eliminate all other values (except digit) from state[coord] and propagate
func (s *Sudoku) assign(coord string, digit rune) error {
	other_values := filterRune(s.State[coord], digit)
	for _, other_digit := range other_values {
		err := s.eliminate(coord, other_digit)
		if err != nil {
			return fmt.Errorf("Failed to remove other digit `%v` from peer %v: %v", string(other_digit), coord, err)
		}
	}

	return nil
}

// eliminates digit from state[coord]; propagates when places <= 2
// returns error if contradiction is detected
func (s *Sudoku) eliminate(coord string, digit rune) error {
	// skip if it is already eliminated
	if !strings.ContainsRune(s.State[coord], digit) {
		return nil
	}

	// eliminate the digit
	s.State[coord] = filterRune(s.State[coord], digit)
	square := s.State[coord]

	// rule 1: if square s is reduced to single value, then eliminate it from peers
	if len(square) == 0 {
		return errors.New("Contradiction: eliminated too many values")
	}

	if len(square) == 1 {
		for _, peer := range Peers[coord] {
			err := s.eliminate(peer, rune(square[0]))
			if err != nil {
				return errors.New("Failed eliminate value from peers")
			}
		}
	}

	// rule 2: if a unit u is reduced to only place for a value d, then put it there
	for _, unit := range Units[coord] {
		var dplaces []string
		for _, unit_coord := range unit {
			if strings.ContainsRune(s.State[unit_coord], digit) {
				dplaces = append(dplaces, unit_coord)
			}
		}

		if len(dplaces) == 0 {
			// contradiction: no place at units
			return errors.New("eliminate: cant assign number to any place in unit")
		}

		if len(dplaces) == 1 {
			return s.assign(dplaces[0], digit)
		}
	}

	return nil
}

// helpers

func filterRune(txt string, needle rune) string {
	return strings.Map(func(r rune) rune {
		if r == needle {
			return -1
		} else {
			return r
		}
	}, txt)
}

/// cross products of 2 strings
/// it is used to generate coordinate pairs
func crossProduct(A, B string) []string {
	var products []string

	for _, a := range A {
		for _, b := range B {
			products = append(products, string(a)+string(b))
		}
	}

	return products
}

// returns column unit
func build_unitlist(rows string, columns string) [][]string {
	var units [][]string

	// units by column
	for _, c := range columns {
		units = append(units, crossProduct(rows, string(c)))
	}

	// units by row
	for _, r := range rows {
		units = append(units, crossProduct(string(r), columns))
	}

	// units by box
	row_boxes := []string{"ABC", "DEF", "GHI"}
	column_boxes := []string{"123", "456", "789"}
	for _, rs := range row_boxes {
		for _, cs := range column_boxes {
			units = append(units, crossProduct(rs, cs))
		}
	}

	return units
}

type square_table = map[string][]string

func hasSquare(unit []string, square string) bool {
	if len(unit) == 0 {
		return false
	}

	for _, item := range unit {
		if item == square {
			return true
		}
	}

	return false
}

func build_units(rows, columns string) map[string][][]string {
	squares := crossProduct(rows, columns)
	unitlist := build_unitlist(rows, columns)

	idx := make(map[string][][]string)

	for _, square := range squares {
		var square_units [][]string

		for _, unit := range unitlist {
			if hasSquare(unit, square) == true {
				square_units = append(square_units, unit)
			}
		}

		idx[square] = square_units
	}

	return idx
}

func build_peer_list(square_units [][]string, square string) []string {
	uniq_peers := make(map[string]bool)

	for _, unit := range square_units {
		for _, unit_square := range unit {
			if unit_square != square {
				uniq_peers[unit_square] = true
			}
		}
	}

	var peer_list []string
	for k := range uniq_peers {
		peer_list = append(peer_list, k)
	}

	return peer_list
}

func build_peers(rows, columns string) map[string][]string {
	squares := crossProduct(rows, columns)
	units := build_units(rows, columns)

	peers := make(map[string][]string)
	for _, square := range squares {
		peers[square] = build_peer_list(units[square], square)
	}

	return peers
}
