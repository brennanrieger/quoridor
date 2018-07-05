package main

import (
	"fmt"
)

func main() {
	game := &Game{}
	p1 := &RandomPlayer{}
	game.Init(p1)
	fmt.Println(game)
}

type Player interface {
	Move() int
}

type RandomPlayer struct{}

func (rp *RandomPlayer) Move() int {
	return 0
}

type Game struct {
	p0 *Player
}

func (g *Game) Init(p0 Player) {
	g.p0 = &p0
}
