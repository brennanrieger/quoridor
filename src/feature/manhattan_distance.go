// WIP

package feature

import (
	"board"
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
	md.distances.Init(b.NRows, b.NCols)
	md.neighbors = []*board.Pos{b.Pos0}

	for head := 0; md.val == 0; head++ {
		md.bfs(head)
	}
	return md.val

}

func (md *ManhattanDistance) bfs(head int) {
	pos := md.neighbors[head]
	dist := md.distances.Get(pos).(int)
	b := md.board
	curWalker := true //TODO: make this real

	if !b.VertiWalls.Get(pos) && pos.Col != 0 && md.distances.Get(pos.L()) == 0 {
		md.distances.Set(pos.L(), dist+1)
		md.neighbors = append(md.neighbors, pos.L())
	}
	if !b.VertiWalls.Get(pos.R()) && pos.Col != b.NCols-1 && md.distances.Get(pos.R()) == 0 {
		md.distances.Set(pos.R(), dist+1)
		md.neighbors = append(md.neighbors, pos.R())
	}
	if !b.HorizWalls.Get(pos) && pos.Row == 0 && curWalker {
		md.val = dist + 1
	} else if !b.HorizWalls.Get(pos) && pos.Row != 0 && md.distances.Get(pos.D()) == 0 {
		md.distances.Set(pos.D(), dist+1)
		md.neighbors = append(md.neighbors, pos.D())
	}
	if !b.HorizWalls.Get(pos.U()) && pos.Row == b.NRows-1 && !curWalker {
		md.val = dist + 1
	} else if !b.HorizWalls.Get(pos.U()) && pos.Row != b.NRows-1 && md.distances.Get(pos.U()) == 0 {
		md.distances.Set(pos.U(), dist+1)
		md.neighbors = append(md.neighbors, pos.U())
	}

}
