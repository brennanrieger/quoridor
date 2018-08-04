package main

import (
	"fmt"
)

// TODO: read https://pdfs.semanticscholar.org/acad/6962a9bb3eb3fde4272f476d6625eb0a8182.pdf
// TODO: idea use ML on small scale board to find good heuristics
// TODO: put all move types in its own file
// TODO: make Matrix non-exported (matrix)
// TODO: make NRows and NCols args to Init not capitalized
// TODO: curPlayer is probably a board concern, not game concern

func main() {
	av := &AsciiVisualizer{}
	game := &Game{}
	p0 := &RandomPlayer{}
	p1 := &HumanPlayer{}
	game.Init(9, 11, p0, p1, av)
	bool := game.Play()
	if bool {
		fmt.Println("player 1 wins")
	} else {
		fmt.Println("player 0 wins")
	}
}
