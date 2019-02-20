package main

import (
	"fmt"
	"game"
	"player"
	"visualizer"
)

// TODO: read https://pdfs.semanticscholar.org/acad/6962a9bb3eb3fde4272f476d6625eb0a8182.pdf
// TODO: idea use ML on small scale board to find good heuristics
// TODO: make Matrix non-exported (matrix) - actually this might not work...
// TODO: remove stutter wiht player package
// TODO: add testing
// TODO: make vertiwall and horizwall functions rather than using get function
// TODO: remove unnecessary parentheses in if statements
// TODO: investigate why implementing jumping changed manhattan distance player behavior
// TODO: review capitalization on matrix.grid and matrix.NRows
// TODO: create randomized dfs feature (n samples)
// TODO: use better system than true/false for current player (enum based on boolean?)
// TODO: test StepMove method (research how to test a move with a panic)
// TODO: can I replace all custom Equal methods with reflect.DeepEqual?
// TODO: is there a better testing framework that doesn't require gc.equals true
// TODO: test make winning move separately
// TODO: capitalization for playerNum
// TODO: make edge numbers not look like part of board
// TODO: rerender board when displaying help instructions
// TODO: replace Init methods with methods to return the initialized object?

func main() {
	av := &visualizer.AsciiVisualizer{}
	game := &game.Game{}
	p0 := &player.HumanPlayer{
		PlayerNum: false,
	}
	p1 := &player.HumanPlayer{
		PlayerNum: true,
	}
	// p0.Init(false, &feature.ManhattanDistance{})
	game.Init(7, 7, p0, p1, av)
	bool, _ := game.Play()
	if bool {
		fmt.Println("player 1 wins")
	} else {
		fmt.Println("player 0 wins")
	}
}
