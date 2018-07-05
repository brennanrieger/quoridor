package main

import "fmt"

type Game struct {
	board *Board

	p0         *Player
	p1         *Player
	visualizer *Visualizer

	win chan bool
}

func (g *Game) Init(n_rows int, n_cols int, p0 Player, p1 Player, v Visualizer) {
	g.p0 = &p0
	g.p1 = &p1

	win := make(chan bool)
	g.board = &Board{}
	g.board.Init(n_rows, n_cols, win)
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
			var moveType MoveType
			var wallPos *Pos
			if g.board.curPlayer {
				moveType, wallPos = (*g.p1).Move(g.board)
			} else {
				moveType, wallPos = (*g.p0).Move(g.board)
			}
			if err := g.board.Move(moveType, wallPos); err != nil {
				// if player makes invalid move other player wins
				fmt.Println(err)
				fmt.Println("bad move")
				return !g.board.curPlayer
			}
		}
	}
}

func (g *Game) Display() {
	v := *g.visualizer
	v.Display(g.board)
}
