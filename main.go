package main

import (
	"fmt"
)

func main() {
	av := &AsciiVisualizer{}
	game := &Game{}
	p0 := &RandomPlayer{}
	p1 := &RandomPlayer{}
	game.Init(19, 9, p0, p1, av)
	bool := game.Play()
	if bool {
		fmt.Println("player 1 wins")
	} else {
		fmt.Println("player 0 wins")
	}
}
