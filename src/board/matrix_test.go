package board

import (
	"reflect"

	gc "github.com/go-check/check"
)

type MatrixSuite struct {
	srcMatrix *Matrix
}

func (s *MatrixSuite) SetUpSuite(c *gc.C) {
	s.srcMatrix = &Matrix{
		NRows: 2,
		NCols: 3,
		grid:  []bool{true, true, false, false, true, false},
	}
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

var _ = gc.Suite(new(MatrixSuite))
