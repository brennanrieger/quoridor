package board

import (
	"reflect"
	"testing"
  "board"

	gc "github.com/go-check/check"
)

func Test(t *testing.T) { gc.TestingT(t) }

type ManhattanDistanceSuite struct {
	board *Board
}

func (s *ManhattanDistanceSuite) SetUpTest(c *gc.C) {
	s.board = &board.Board{}
	s.board.Init(4, 4)
}

func (s *BoardSuite) TestInitialDistance(c *gc.C) {
	for i := range(3,10) {
    board := &board.Board{}
    board.Init(i,i)
    md := &ManhattanDistance{}
    val0, val1 := md.Val(board)
    c.Check(val0, gc.Equals, i)
    c.Check(val1, gc.Equals, i)
  }
}

func (s *BoardSuite) TestReuseManhattanDistanceGenerator(c *gc.C) {

var _ = gc.Suite(new(BoardSuite))
