package board

import (
	"reflect"

	gc "github.com/go-check/check"
)

type MatrixSuite struct {
	srcMatrix *Matrix
}

func (s *MatrixSuite) SetUpTest(c *gc.C) {
	s.srcMatrix = &Matrix{
		NRows: 2,
		NCols: 3,
		grid:  []bool{true, true, false, false, true, false},
	}
}

func (s *MatrixSuite) TestInit(c *gc.C) {
	matrix := &Matrix{}
	matrix.Init(3, 4)

	c.Check(matrix.NRows, gc.Equals, 3)
	c.Check(matrix.NCols, gc.Equals, 4)
	c.Check(len(matrix.grid), gc.Equals, 12)
}

func (s *MatrixSuite) TestEqual(c *gc.C) {
	sameMatrix := &Matrix{
		NRows: 2,
		NCols: 3,
		grid:  []bool{true, true, false, false, true, false},
	}
	diffMatrix := &Matrix{
		NRows: 2,
		NCols: 3,
		grid:  []bool{true, true, true, false, true, false},
	}
	c.Check(s.srcMatrix.Equal(sameMatrix), gc.Equals, true)
	c.Check(s.srcMatrix.Equal(diffMatrix), gc.Equals, false)
}

func (s *MatrixSuite) TestCopy(c *gc.C) {
	// Check s.srcMatrix is copied by value but not reference
	c.Check(s.srcMatrix, gc.Not(gc.Equals), s.srcMatrix.Copy())
	c.Check(s.srcMatrix.NRows, gc.Equals, s.srcMatrix.Copy().NRows)
	c.Check(s.srcMatrix.NCols, gc.Equals, s.srcMatrix.Copy().NCols)

	// Check grid is copied by value but not reference
	c.Check(s.srcMatrix.grid, gc.Not(gc.Equals), s.srcMatrix.Copy().grid)
	c.Check(reflect.DeepEqual(s.srcMatrix.grid, s.srcMatrix.Copy().grid), gc.Equals, true)
}

func (s *MatrixSuite) TestFlip(c *gc.C) {
	destMatrix := &Matrix{
		NRows: 2,
		NCols: 3,
		grid:  []bool{false, true, false, false, true, true},
	}
	c.Check(reflect.DeepEqual(s.srcMatrix.Flip(), destMatrix), gc.Equals, true)
}

func (s *MatrixSuite) TestSet(c *gc.C) {
	destMatrix := &Matrix{
		NRows: 2,
		NCols: 3,
		grid:  []bool{true, true, true, false, true, false},
	}
	pos := &Pos{
		Row: 0,
		Col: 2,
	}
	s.srcMatrix.Set(pos)
	c.Check(reflect.DeepEqual(s.srcMatrix, destMatrix), gc.Equals, true)
}

func (s *MatrixSuite) TestGet(c *gc.C) {
	pos := &Pos{
		Row: 0,
		Col: 2,
	}
	c.Check(s.srcMatrix.Get(pos), gc.Equals, false)
}

var _ = gc.Suite(new(MatrixSuite))
