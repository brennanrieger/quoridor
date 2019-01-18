package main

import (
	"feature"
	"fmt"
	"game"
	"player"
	"visualizer"
)

// TODO: read https://pdfs.semanticscholar.org/acad/6962a9bb3eb3fde4272f476d6625eb0a8182.pdf
// TODO: idea use ML on small scale board to find good heuristics
// TODO: make Matrix non-exported (matrix) - actually this might not work...
// TODO: curPlayer is probably a board concern, not game concern <== *this*
// TODO: remove stutter wiht player package
// TODO: add testing
// TODO: bug where you can't add wall to edge
// TODO: make vertiwall and horizwall functions rather than using get function
// TODO: remove unnecessary parentheses in if statements
// TODO: investigate why implementing jumping changed manhattan distance player behavior

func main() {
	av := &visualizer.AsciiVisualizer{}
	game := &game.Game{}
	p0 := &player.FeaturePlayer{}
	p1 := &player.HumanPlayer{}
	p0.Init(false, &feature.ManhattanDistance{})
	game.Init(9, 9, p0, p1, av)
	bool := game.Play()
	if bool {
		fmt.Println("player 1 wins")
	} else {
		fmt.Println("player 0 wins")
	}
}
