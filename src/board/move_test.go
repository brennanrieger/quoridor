package board

import (
	gc "github.com/go-check/check"
)

type MoveSuite struct{}

func (s *MoveSuite) TestCopy(c *gc.C) {
	pos = &Pos{
		Row: 10,
		Col: 10,
	}
	move = &Move{
		Mt:  1,
		Pos: pos,
	}

	// Check move is copied by value but not reference
	c.Check(move, gc.Not(gc.Equals), move.Copy())
	c.Check(move.Mt, gc.Equals, move.Copy().Mt)

	// Check pos is copied by value but not reference
	c.Check(move.Pos, gc.Not(gc.Equals), move.Copy().Pos)
	c.Check(move.Pos.Row, gc.Equals, move.Copy().Pos.Row)
	c.Check(move.Pos.Col, gc.Equals, move.Copy().Pos.Col)
}

var _ = gc.Suite(new(MoveSuite))
