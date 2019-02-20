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
}

func (g *Game) Init(NRows int, NCols int, p0 player.Player, p1 player.Player, v visualizer.Visualizer) error {
	if NCols < 2 {
		return fmt.Errorf("must have at least two columns")
	}

	g.p0 = &p0
	g.p1 = &p1

	g.board = &board.Board{}
	g.board.Init(NRows, NCols)

	g.visualizer = &v

	return nil
}

func (g *Game) Play() (bool, error) {
	for {
		select {
		case winner := <-g.board.Win:
			return winner, nil
		default:
			g.Display()
			var move *board.Move
			var err error
			if g.board.CurPlayer {
				move = (*g.p1).Move(g.board)
				err = g.board.MakeMove(true, move)
			} else {
				move = (*g.p0).Move(g.board)
				err = g.board.MakeMove(false, move)
			}
			if err != nil {
				return !g.board.CurPlayer, fmt.Errorf("Illegal move: %s", err)
			}
		}
	}
}

func (g *Game) Display() {
	v := *g.visualizer
	v.Display(g.board)
}
