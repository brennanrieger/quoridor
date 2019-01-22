package board

import (
	"reflect"

	gc "github.com/go-check/check"
)

type MatrixSuite struct {
	matrix *Matrix
}

func (s *MatrixSuite) SetUpSuite(c *gc.C) {
	s.matrix = &Matrix{
		NRows: 2,
		NCols: 3,
		grid:  []bool{true, true, false, false, true, false},
	}
}

func (s *MatrixSuite) TestCopy(c *gc.C) {
	// Check s.matrix is copied by value but not reference
	c.Check(s.matrix, gc.Not(gc.Equals), s.matrix.Copy())
	c.Check(s.matrix.NRows, gc.Equals, s.matrix.Copy().NRows)
	c.Check(s.matrix.NCols, gc.Equals, s.matrix.Copy().NCols)

	// Check grid is copied by value but not reference
	c.Check(s.matrix.grid, gc.Not(gc.Equals), s.matrix.Copy().grid)
	c.Check(reflect.DeepEqual(s.matrix.grid, s.matrix.Copy().grid), gc.Equals, true)
}

var _ = gc.Suite(new(MatrixSuite))
