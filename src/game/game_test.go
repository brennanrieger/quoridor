package game

import (
	"player"
	"testing"
	"visualizer"

	gc "github.com/go-check/check"
)

func Test(t *testing.T) { gc.TestingT(t) }

type GameSuite struct{}

func (s *GameSuite) TestSmoke(c *gc.C) {
	game := &Game{}
	p0 := &player.RandomPlayer{}
	p0.Init(false)
	p1 := &player.RandomPlayer{}
	p1.Init(true)
	visualizer := &visualizer.NilVisualizer{}
	game.Init(3, 3, p0, p1, visualizer)
	game.Play()
}

var _ = gc.Suite(new(GameSuite))
