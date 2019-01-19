package board

import (
	gc "github.com/go-check/check"
)

type PosSuite struct{}

func (s *PosSuite) TestEqual(c *gc.C) {
	srcPos := &Pos{
		Row: 10,
		Col: 10,
	}
	samePos := &Pos{
		Row: 10,
		Col: 10,
	}
	diffPos := &Pos{
		Row: 10,
		Col: 11,
	}
	c.Check(srcPos.Equal(samePos), gc.Equals, true)
	c.Check(srcPos.Equal(diffPos), gc.Equals, false)
}

func (s *PosSuite) TestU(c *gc.C) {
	srcPos := &Pos{
		Row: 10,
		Col: 10,
	}
	destPos := &Pos{
		Row: 11,
		Col: 10,
	}
	c.Check(srcPos.U().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestD(c *gc.C) {
	srcPos := &Pos{
		Row: 10,
		Col: 10,
	}
	destPos := &Pos{
		Row: 9,
		Col: 10,
	}
	c.Check(srcPos.U().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestR(c *gc.C) {
	srcPos := &Pos{
		Row: 10,
		Col: 10,
	}
	destPos := &Pos{
		Row: 10,
		Col: 11,
	}
	c.Check(srcPos.U().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestL(c *gc.C) {
	srcPos := &Pos{
		Row: 10,
		Col: 10,
	}
	destPos := &Pos{
		Row: 10,
		Col: 9,
	}
	c.Check(srcPos.U().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestCopy(c *gc.C) {
	srcPos := &Pos{
		Row: 10,
		Col: 10,
	}
	c.Check(srcPos.Equal(srcPos.Copy()), gc.Equals, true)
	c.Check(srcPos, gc.Not(gc.Equals), srcPos.Copy())
}

var _ = gc.Suite(new(PosSuite))
