package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	av := &AsciiVisualizer{}
	game := &Game{}
	p0 := &RandomPlayer{}
	p1 := &RandomPlayer{}
	game.Init(3, 3, p0, p1, av)
	bool := game.Play()
	if bool {
		fmt.Println("player 1 wins")
	} else {
		fmt.Println("player 0 wins")
	}
}
