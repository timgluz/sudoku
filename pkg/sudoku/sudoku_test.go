package sudoku

import (
	"fmt"
	"sort"
	"testing"
)

func assertStringSlices(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCrossProductwithHappyCases(t *testing.T) {
	var examples = []struct {
		a, b     string
		expected []string
	}{
		{"", "12", nil},
		{"A", "1", []string{"A1"}},
		{"A", "12", []string{"A1", "A2"}},
		{"AB", "1", []string{"A1", "B1"}},
		{"AB", "12", []string{"A1", "A2", "B1", "B2"}},
	}

	for _, tt := range examples {
		testname := fmt.Sprintf("crossProduct(%s,%s)", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			res := crossProduct(tt.a, tt.b)
			if !assertStringSlices(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestBuildUnitlistWithSudokuBoard(t *testing.T) {
	rows := "ABCDEFGHI"
	columns := "123456789"

	res := build_unitlist(rows, columns)
	if len(res) != 27 {
		t.Errorf("got %v, expected %v for unitlist length", len(res), 27)
	}

	// expects length of first row to be 9
	if len(res[0]) != 9 {
		t.Errorf("got %v, expected %v for unitlist item length", len(res[0]), 9)
	}

	expected_unit := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1"}
	if !assertStringSlices(res[0], expected_unit) {
		t.Errorf("got %v, expected %v to be first unitlist item", res[0], expected_unit)
	}
}

func TestBuildUnits(t *testing.T) {
	rows := "ABCDEFGHI"
	columns := "123456789"

	units := build_units(rows, columns)

	if len(units) != 81 {
		t.Errorf("got %v, expected %v units", len(units), 81)
	}

	c2_units, ok := units["C2"]
	if !ok {
		t.Errorf("found no C2 in units")
	}

	if len(c2_units) != 3 {
		t.Errorf("got %v, expected 3 units for C2", len(c2_units))
	}

	expected_unit := []string{"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2", "I2"}
	if !assertStringSlices(c2_units[0], expected_unit) {
		t.Errorf("got %v, expected %v as first unit", c2_units[0], expected_unit)
	}

}

func TestBuildPeers(t *testing.T) {
	rows := "ABCDEFGHI"
	columns := "123456789"

	peers := build_peers(rows, columns)

	if len(peers) != 81 {
		t.Errorf("got %v, expected 81 squares in peer tables", len(peers))
	}

	c2_peers, ok := peers["C2"]
	if !ok {
		t.Errorf("found no C2 in peers")
	}

	expected_peers := []string{
		"A2", "B2", "D2", "E2", "F2", "G2", "H2", "I2",
		"C1", "C3", "C4", "C5", "C6", "C7", "C8", "C9",
		"A1", "A3", "B1", "B3",
	}

	if !assertStringSlices(c2_peers, expected_peers) {
		t.Errorf("got %v, excpected %v as C2 peers", c2_peers, expected_peers)
	}
}

func TestParseGridHappyCase(t *testing.T) {
	example_grid := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"

	state, err := parseGrid(example_grid)
	if err != nil {
		t.Errorf("got error, expected to parse the example grid without error")
	}

	// first row
	if state["A1"] != "4" {
		t.Errorf("got %v, expected 4 for A1", state["A1"])
	}

	if state["A2"] != DIGITS {
		t.Errorf("got %v, expected %v for A2", state["A2"], DIGITS)
	}

	if state["A7"] != "8" {
		t.Errorf("got %v, expected 8 for A7", state["A7"])
	}

	if state["A9"] != "5" {
		t.Errorf("got %v, expected 5 for A9", state["A9"])
	}

	// second row
	if state["B1"] != DIGITS {
		t.Errorf("got %v, expected %v for B1", state["B1"], DIGITS)
	}

	if state["B2"] != "3" {
		t.Errorf("got %v, expected 3 for B2", state["B2"])
	}
}

func TestParseGridWithInvalidExamples(t *testing.T) {
	_, err := parseGrid("")
	if err == nil {
		t.Errorf("got no error, expected to see error with empty string")
	}

	_, err = parseGrid("...0")
	if err == nil {
		t.Errorf("got no error, expected error when string shorter than 81 characters")
	}

	too_long := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......+"
	_, err = parseGrid(too_long)
	if err == nil {
		t.Errorf("got no error, expected error when string longer than 81 characters")
	}
}
