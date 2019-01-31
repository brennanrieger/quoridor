package game

import (
	"board"
	"fmt"
	"player"
	"visualizer"
)

type Game struct {
	board *board.Board

	p0         *player.Player
	p1         *player.Player
	visualizer *visualizer.Visualizer

	win chan bool
}

func (g *Game) Init(NRows int, NCols int, p0 player.Player, p1 player.Player, v visualizer.Visualizer) {
	g.p0 = &p0
	g.p1 = &p1

	g.win = make(chan bool, 2)

	g.board = &board.Board{}
	g.board.Init(NRows, NCols, g.win)

	g.visualizer = &v
}

func (g *Game) Play() bool {
	for {
		select {
		case winner := <-g.win:
			fmt.Println("won")
			return winner
		default:
			g.Display()
			var move *board.Move
			if g.board.CurPlayer {
				move = (*g.p1).Move(g.board)
			} else {
				move = (*g.p0).Move(g.board)
			}
			if err := g.board.MakeMove(move, g.win); err != nil {
				// if player makes invalid move other player wins
				fmt.Println(err)
				fmt.Println("bad move")
				return !g.board.CurPlayer
			}
		}
	}
}

func (g *Game) Display() {
	v := *g.visualizer
	v.Display(g.board)
}
