package main

import (
	"fmt"
	"game"
	"player"
	"visualizer"
)

// TODO: read https://pdfs.semanticscholar.org/acad/6962a9bb3eb3fde4272f476d6625eb0a8182.pdf
// TODO: idea use ML on small scale board to find good heuristics
// TODO: put all move types in its own file
// TODO: make Matrix non-exported (matrix)
// TODO: curPlayer is probably a board concern, not game concern
// TODO: remove stutter wiht player package

func main() {
	av := &visualizer.AsciiVisualizer{}
	game := &game.Game{}
	p0 := &player.RandomPlayer{}
	p1 := &player.HumanPlayer{}
	game.Init(9, 11, p0, p1, av)
	bool := game.Play()
	if bool {
		fmt.Println("player 1 wins")
	} else {
		fmt.Println("player 0 wins")
	}
}
