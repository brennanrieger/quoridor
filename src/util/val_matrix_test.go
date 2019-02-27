package util

import (
	"board"
	"reflect"
	"testing"

	gc "github.com/go-check/check"
)

func Test(t *testing.T) { gc.TestingT(t) }

type ValMatrixSuite struct {
	srcMatrix *ValMatrix
}

func (s *ValMatrixSuite) SetUpTest(c *gc.C) {

	// Initialization so the ValMatrix is interesting
	//
	//  0  3  0
	//  1  0  2

	s.srcMatrix = &ValMatrix{}
	s.srcMatrix.Init(2, 3)
	s.srcMatrix.Set(
		&board.Pos{
			Row: 0,
			Col: 0,
		},
		1,
	)
	s.srcMatrix.Set(
		&board.Pos{
			Row: 0,
			Col: 2,
		},
		2,
	)
	s.srcMatrix.Set(
		&board.Pos{
			Row: 1,
			Col: 1,
		},
		3,
	)
}

func (s *ValMatrixSuite) TestInit(c *gc.C) {
	matrix := &ValMatrix{}
	matrix.Init(3, 4)

	c.Check(matrix.NRows, gc.Equals, 3)
	c.Check(matrix.NCols, gc.Equals, 4)
	c.Check(len(matrix.grid), gc.Equals, 12)
}

func (s *ValMatrixSuite) TestEqual(c *gc.C) {
	sameMatrix := &ValMatrix{}
	sameMatrix.Init(2, 3)
	sameMatrix.Set(
		&board.Pos{
			Row: 0,
			Col: 0,
		},
		1,
	)
	sameMatrix.Set(
		&board.Pos{
			Row: 0,
			Col: 2,
		},
		2,
	)
	sameMatrix.Set(
		&board.Pos{
			Row: 0,
			Col: 1,
		},
		3,
	)

	diffMatrix := &ValMatrix{}
	diffMatrix.Init(2, 3)
	diffMatrix.Set(
		&board.Pos{
			Row: 0,
			Col: 0,
		},
		7,
	)

	c.Check(s.srcMatrix.Equal(sameMatrix), gc.Equals, true)
	c.Check(s.srcMatrix.Equal(diffMatrix), gc.Equals, false)
}

func (s *ValMatrixSuite) TestCopy(c *gc.C) {
	// Check s.srcMatrix is copied by value but not reference
	c.Check(s.srcMatrix, gc.Not(gc.Equals), s.srcMatrix.Copy())
	c.Check(s.srcMatrix.NRows, gc.Equals, s.srcMatrix.Copy().NRows)
	c.Check(s.srcMatrix.NCols, gc.Equals, s.srcMatrix.Copy().NCols)

	// Check grid is copied by value but not reference
	c.Check(s.srcMatrix.grid, gc.Not(gc.Equals), s.srcMatrix.Copy().grid)
	c.Check(reflect.DeepEqual(s.srcMatrix.grid, s.srcMatrix.Copy().grid), gc.Equals, true)
}

func (s *ValMatrixSuite) TestFlip(c *gc.C) {
	destMatrix := &ValMatrix{}
	destMatrix.Init(2, 3)
	destMatrix.Set(
		&board.Pos{
			Row: 1,
			Col: 2,
		},
		1,
	)
	destMatrix.Set(
		&board.Pos{
			Row: 1,
			Col: 0,
		},
		2,
	)
	destMatrix.Set(
		&board.Pos{
			Row: 0,
			Col: 1,
		},
		3,
	)

	c.Check(reflect.DeepEqual(s.srcMatrix.Flip(), destMatrix), gc.Equals, true)
}

func (s *ValMatrixSuite) TestGet(c *gc.C) {
	pos := &board.Pos{
		Row: 0,
		Col: 2,
	}
	c.Check(s.srcMatrix.Get(pos), gc.Equals, 2)
}

func (s *ValMatrixSuite) TestSet(c *gc.C) {
	pos := &board.Pos{
		Row: 0,
		Col: 2,
	}
	s.srcMatrix.Set(pos, 7)
	c.Check(s.srcMatrix.Get(pos), gc.Equals, 7)
}

func (s *ValMatrixSuite) TestCharVal(c *gc.C) {
	stringMatrix := &ValMatrix{}
	stringMatrix.Init(2, 3)
	pos := &board.Pos{
		Row: 0,
		Col: 0,
	}
	stringMatrix.Set(pos, 'a')
	c.Check(stringMatrix.Get(pos), gc.Equals, 'a')
}

func (s *ValMatrixSuite) TestStringVal(c *gc.C) {
	stringMatrix := &ValMatrix{}
	stringMatrix.Init(2, 3)
	pos := &board.Pos{
		Row: 0,
		Col: 0,
	}
	stringMatrix.Set(pos, "hello")
	c.Check(stringMatrix.Get(pos), gc.Equals, "hello")
}

var _ = gc.Suite(new(ValMatrixSuite))
