package board

import (
	"reflect"

	gc "github.com/go-check/check"
)

type MatrixSuite struct{}

func (s *MatrixSuite) TestCopy(c *gc.C) {
	matrix = &Matrix{
		NRows: 2,
		NCols: 2,
		grid:  make([]bool, 4, 4),
	}

	// Check matrix is copied by value but not reference
	c.Check(matrix, gc.Not(gc.Equals), matrix.Copy())
	c.Check(matrix.NRows, gc.Equals, matrix.Copy().NRows)
	c.Check(matrix.NCols, gc.Equals, matrix.Copy().NCols)

	// Check grid is copied by value but not reference
	c.Check(matrix.grid, gc.Not(gc.Equals), matrix.Copy().grid)
	c.Check(reflect.DeepEqual(matrix.grid, matrix.grid.Copy()), gc.Equals, true)
}

var _ = gc.Suite(new(MatrixSuite))
