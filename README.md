# SUDOKU SOLVER ala Norvig

Go adoption of Norvig's blog post: ["Solving Every Sudoku Puzzle"](http://norvig.com/sudoku.html)

It small Sudoku solver that demonstrates power of Constraint Propagation and Search;

### Usage

By default the solver cmd returns a line with solution, which you could pass along to another CLI tool or send to frontend.



`-human` flag prints out human friendly Sudoko table



```
$> go run cmd/solver/main.go -human 003020600900305001001806400008102900700000008006708200002609500800203009005010300
 4 8 3| 9 2 1| 6 5 7
 9 6 7| 3 4 5| 8 2 1
 2 5 1| 8 7 6| 4 9 3
------+------+------
 5 4 8| 1 3 2| 9 7 6
 7 2 9| 5 6 4| 1 3 8
 1 3 6| 7 9 8| 2 4 5
------+------+------
 3 7 2| 6 8 9| 5 1 4
 8 1 4| 2 5 3| 7 6 9
 6 9 5| 4 1 7| 3 8 2

```

### Testing

```
# run all integration test against hard problems
go test -v -run TestSudokuSolve
```
