package board

import (
	"reflect"

	gc "github.com/go-check/check"
)

type BoardSuite struct {
	srcBoard *Board
}

func (s *BoardSuite) SetUpTest(c *gc.C) {
	s.srcBoard = &Board{}
	s.srcBoard.Init(4, 4)
}

func (s *BoardSuite) TestInit(c *gc.C) {
	// Player 0 moves first
	board1 := &Board{}
	board1.Init(3, 3)
	c.Check(board1.CurPlayer, gc.Equals, false)

	// If odd size board, players start in middle on top and bottom
	// Player 0 is on bottom; player 1 is on top
	board2 := &Board{}
	board2.Init(9, 9)
	board2_pos0 := &Pos{
		Row: 0,
		Col: 4,
	}
	board2_pos1 := &Pos{
		Row: 8,
		Col: 4,
	}
	c.Check(board2.Pos0.Equal(board2_pos0), gc.Equals, true)
	c.Check(board2.Pos1.Equal(board2_pos1), gc.Equals, true)

	// If even size board players start slightly to right on top and bottom
	// Player 0 is on bottom; player 1 is on top
	board3 := &Board{}
	board3.Init(10, 10)
	board3_pos0 := &Pos{
		Row: 0,
		Col: 5,
	}
	board3_pos1 := &Pos{
		Row: 9,
		Col: 5,
	}
	c.Check(board3.Pos0.Equal(board3_pos0), gc.Equals, true)
	c.Check(board3.Pos1.Equal(board3_pos1), gc.Equals, true)
}

func (s *BoardSuite) TestCopy(c *gc.C) {

	// Random initialization so board is interesting
	s.srcBoard.MakeMove(StepMove(Up))
	s.srcBoard.MakeMove(StepMove(Left))
	pos := &Pos{
		Row: 0,
		Col: 0,
	}
	s.srcBoard.MakeMove(&Move{
		Mt:  HorizWall,
		Pos: pos,
	})
	s.srcBoard.MakeMove(&Move{
		Mt:  VertiWall,
		Pos: pos,
	})

	// Check s.srcBoard is copied by value but not reference
	c.Check(s.srcBoard, gc.Not(gc.Equals), s.srcBoard.Copy())
	c.Check(s.srcBoard.NRows, gc.Equals, s.srcBoard.Copy().NRows)
	c.Check(s.srcBoard.NCols, gc.Equals, s.srcBoard.Copy().NCols)
	c.Check(s.srcBoard.CurPlayer, gc.Equals, s.srcBoard.Copy().CurPlayer)

	// Check Pos0 is copied by value but not reference
	c.Check(s.srcBoard.Pos0, gc.Not(gc.Equals), s.srcBoard.Copy().Pos0)
	c.Check(reflect.DeepEqual(s.srcBoard.Pos0.Row, s.srcBoard.Copy().Pos0.Row), gc.Equals, true)
	c.Check(reflect.DeepEqual(s.srcBoard.Pos0.Col, s.srcBoard.Copy().Pos0.Col), gc.Equals, true)

	// Check Pos1 is copied by value but not reference
	c.Check(s.srcBoard.Pos1, gc.Not(gc.Equals), s.srcBoard.Copy().Pos1)
	c.Check(reflect.DeepEqual(s.srcBoard.Pos1.Row, s.srcBoard.Copy().Pos1.Row), gc.Equals, true)
	c.Check(reflect.DeepEqual(s.srcBoard.Pos1.Col, s.srcBoard.Copy().Pos1.Col), gc.Equals, true)

	// Check VertiWalls is copied by value but not reference
	c.Check(s.srcBoard.VertiWalls, gc.Not(gc.Equals), s.srcBoard.Copy().VertiWalls)
	c.Check(reflect.DeepEqual(s.srcBoard.VertiWalls.grid, s.srcBoard.Copy().VertiWalls.grid), gc.Equals, true)

	// Check HorizWalls is copied by value but not reference
	c.Check(s.srcBoard.HorizWalls, gc.Not(gc.Equals), s.srcBoard.Copy().HorizWalls)
	c.Check(reflect.DeepEqual(s.srcBoard.HorizWalls.grid, s.srcBoard.Copy().HorizWalls.grid), gc.Equals, true)
}

// func (s *MatrixSuite) TestFlip(c *gc.C) {
// 	destMatrix := &Matrix{
// 		NRows: 2,
// 		NCols: 3,
// 		grid:  []bool{false, true, false, false, true, true},
// 	}
// 	c.Check(reflect.DeepEqual(s.srcMatrix.Flip(), destMatrix), gc.Equals, true)
// }

// func (s *MatrixSuite) TestSet(c *gc.C) {
// 	destMatrix := &Matrix{
// 		NRows: 2,
// 		NCols: 3,
// 		grid:  []bool{true, true, true, false, true, false},
// 	}
// 	pos := &Pos{
// 		Row: 0,
// 		Col: 2,
// 	}
// 	s.srcMatrix.Set(pos)
// 	c.Check(reflect.DeepEqual(s.srcMatrix, destMatrix), gc.Equals, true)
// }

// func (s *MatrixSuite) TestGet(c *gc.C) {
// 	pos := &Pos{
// 		Row: 0,
// 		Col: 2,
// 	}
// 	c.Check(s.srcMatrix.Get(pos), gc.Equals, false)
// }

var _ = gc.Suite(new(MatrixSuite))
