package board

import (
	gc "github.com/go-check/check"
)

func TestEqual(c *gc.C) {
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

func TestU(c *gc.C) {
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

func TestD(c *gc.C) {
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

func TestR(c *gc.C) {
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

func TestL(c *gc.C) {
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
