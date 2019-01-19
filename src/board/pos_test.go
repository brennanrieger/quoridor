package board

import (
	"testing"

	gc "github.com/go-check/check"
)

func Test(t *testing.T) { gc.TestingT(t) }

type PosSuite struct {
	srcPos *Pos
}

func (s *PosSuite) SetUpSuite(c *gc.C) {
	s.srcPos = &Pos{
		Row: 10,
		Col: 10,
	}
}

func (s *PosSuite) TestEqual(c *gc.C) {
	samePos := &Pos{
		Row: 10,
		Col: 10,
	}
	diffPos := &Pos{
		Row: 10,
		Col: 11,
	}
	c.Check(s.srcPos.Equal(samePos), gc.Equals, true)
	c.Check(s.srcPos.Equal(diffPos), gc.Equals, false)
}

func (s *PosSuite) TestU(c *gc.C) {
	destPos := &Pos{
		Row: 11,
		Col: 10,
	}
	c.Check(s.srcPos.U().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestD(c *gc.C) {
	destPos := &Pos{
		Row: 9,
		Col: 10,
	}
	c.Check(s.srcPos.D().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestR(c *gc.C) {
	destPos := &Pos{
		Row: 10,
		Col: 11,
	}
	c.Check(s.srcPos.R().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestL(c *gc.C) {
	destPos := &Pos{
		Row: 10,
		Col: 9,
	}
	c.Check(s.srcPos.L().Equal(destPos), gc.Equals, true)
}

func (s *PosSuite) TestCopy(c *gc.C) {
	c.Check(s.srcPos.Equal(s.srcPos.Copy()), gc.Equals, true)
	c.Check(s.srcPos, gc.Not(gc.Equals), s.srcPos.Copy())
}

var _ = gc.Suite(new(PosSuite))
