package feature

import (
	"board"
	"testing"

	gc "github.com/go-check/check"
)

func Test(t *testing.T) { gc.TestingT(t) }

type ManhattanDistanceSuite struct{}

func (s *ManhattanDistanceSuite) TestInitialDistance(c *gc.C) {
	for i := 2; i < 10; i++ {
		testBoard := &board.Board{}
		testBoard.Init(i, i)
		md := &ManhattanDistance{}
		val0, val1 := md.Val(testBoard)
		c.Check(int(val0), gc.Equals, i)
		c.Check(int(val1), gc.Equals, i)
	}
}

func (s *ManhattanDistanceSuite) TestReusageGeneratesDiffValues(c *gc.C) {
	testBoard := &board.Board{}
	testBoard.Init(5, 5)
	md := &ManhattanDistance{}

	val0_initial, val1_initial := md.Val(testBoard)
	testBoard.MakeMove(false, &board.Move{
		Mt: board.HorizWall,
		Pos: &board.Pos{
			Row: 2,
			Col: 2,
		},
	})
	val0_final, val1_final := md.Val(testBoard)

	c.Check(val0_initial, gc.Not(gc.Equals), val0_final)
	c.Check(val1_initial, gc.Not(gc.Equals), val1_final)
}

func (s *ManhattanDistanceSuite) TestCanChooseShortestPath(c *gc.C) {
	testBoard := &board.Board{}
	testBoard.Init(4, 4)

	// Create non-trivial position
	//
	//   ·   ·   ·   ·   ·
	//         1
	//   ╶───────╴   ╷   ·
	//               │
	//   ·   ╶───────┤   ·
	//             0 │
	//   ·   ·   ·   ╵   ·
	//
	//   ·   ·   ╶───────╴

	testBoard.MakeMove(false, board.StepMove(board.Up))
	testBoard.MakeMove(true, board.StepMove(board.Left))
	testBoard.MakeMove(false, &board.Move{
		Mt: board.HorizWall,
		Pos: &board.Pos{
			Row: 3,
			Col: 0,
		},
	})
	testBoard.MakeMove(true, &board.Move{
		Mt: board.VertiWall,
		Pos: &board.Pos{
			Row: 1,
			Col: 3,
		},
	})
	testBoard.MakeMove(false, &board.Move{
		Mt: board.HorizWall,
		Pos: &board.Pos{
			Row: 2,
			Col: 1,
		},
	})
	testBoard.MakeMove(true, &board.Move{
		Mt: board.HorizWall,
		Pos: &board.Pos{
			Row: 0,
			Col: 2,
		},
	})

	md := &ManhattanDistance{}
	val0, val1 := md.Val(testBoard)
	c.Check(int(val0), gc.Equals, 6)
	c.Check(int(val1), gc.Equals, 7)
}

var _ = gc.Suite(new(ManhattanDistanceSuite))
