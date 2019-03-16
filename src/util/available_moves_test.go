package util

import (
	gc "github.com/go-check/check"
)

type AvailableMovesSuite struct {
	randomPlayer0 *player.RandomPlayer
	randomPlayer1 *player.RandomPlayer
}

func (s *AvailableMovesSuite) SetUpTest(c *gc.C) {
	s.randomPlayer0.Init(false)
	s.randomPlayer1.Init(true)
}

func (s *AvailableMovesSuite) TestAllAvailableMovesValid(c *gc.C) {
	// test available moves for 100 different random boards
	for i := 0; i < 100; i++ {
		testBoard = &Board{}
		testBoard.Init(6, 6)

		var move *Move

		// Make 10 random moves
		for j := 0; j < 10; j++ {
			for _, move = range AvailableMoves(testBoard) {
				err := testBoard.MakeMove(move)
				c.Check(err, gc.Equals, nil)
			}
		}
	}
}

var _ = gc.Suite(new(ValMatrixSuite))
