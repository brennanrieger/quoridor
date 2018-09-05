// WIP

package feature

import (
	"board"
	"fmt"
	"util"
)

type ManhattanDistance struct {
	board     *board.Board
	distances *util.ValMatrix
	neighbors []*board.Pos
	val       int
}

func (md *ManhattanDistance) Val(b *board.Board) int {
	md.board = b
	md.distances = &util.ValMatrix{}
	md.distances.Init(b.NRows, b.NCols)
	md.neighbors = []*board.Pos{b.Pos0}
	md.distances.Set(b.Pos0, 0)

	for head := 0; md.val == 0; head++ {
		md.bfs(head)
	}
	return md.val

}

func (md *ManhattanDistance) bfs(head int) {
	fmt.Println(head)
	pos := md.neighbors[head]
	dist := md.distances.Get(pos).(int)
	b := md.board
	curWalker := false //TODO: put this in implementation

	if !b.VertiWalls.Get(pos) && pos.Col != 0 && md.distances.Get(pos.L()) == nil {
		md.distances.Set(pos.L(), dist+1)
		md.neighbors = append(md.neighbors, pos.L())
	}
	if !b.VertiWalls.Get(pos.R()) && pos.Col != b.NCols-1 && md.distances.Get(pos.R()) == nil {
		md.distances.Set(pos.R(), dist+1)
		md.neighbors = append(md.neighbors, pos.R())
	}
	if !b.HorizWalls.Get(pos) && pos.Row == 0 && curWalker {
		md.val = dist + 1
	} else if !b.HorizWalls.Get(pos) && pos.Row != 0 && md.distances.Get(pos.D()) == nil {
		md.distances.Set(pos.D(), dist+1)
		md.neighbors = append(md.neighbors, pos.D())
	}
	if !b.HorizWalls.Get(pos.U()) && pos.Row == b.NRows-1 && !curWalker {
		md.val = dist + 1
	} else if !b.HorizWalls.Get(pos.U()) && pos.Row != b.NRows-1 && md.distances.Get(pos.U()) == nil {
		md.distances.Set(pos.U(), dist+1)
		md.neighbors = append(md.neighbors, pos.U())
	}

}
