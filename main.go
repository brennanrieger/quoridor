package main

import (
	"fmt"
)

// TODO: read https://pdfs.semanticscholar.org/acad/6962a9bb3eb3fde4272f476d6625eb0a8182.pdf
// TODO: idea use ML on small scale board to find good heuristics
// TODO: put Move into game.go?

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
