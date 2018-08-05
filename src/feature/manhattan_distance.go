// WIP

package feature

import (
	"board"
	"util"
)

type ManhattanDistance struct {
	board     *board.Board
	distances *util.ValMatrix{}
	neighbors []*board.Pos
	val       int
}

func (md *ManhattanDistance) Val(b *board.Board) int {
	if md.val != 0 {
		return val
	} else {
		md.distances.Init(b.NRows, b.NCols)
		md.neighbors = []*board.Pos{b.Pos0}

		for head := 0; md.val == 0; head++ {
			bfs(head)
		}
		return md.val
	}

}

func (md *ManhattanDistance) bfs(pos *board.Pos) {
	pos := neighbors[head]
	dist := distances.Get(pos)
	b := md.board

	if !b.VertiWalls.Get(pos) && pos.Col != 0 && md.distances.Get(pos.L()) == 0 {
		md.Set(pos.L(), dist+1)
		md.neighbors = append(md.neighbors, pos.L())
	}
	if !b.VertiWalls.Get(pos.R()) && pos.Col != b.NCols-1 && md.distances.Get(pos.R()) == 0 {
		md.Set(pos.R(), dist+1)
		md.neighbors = append(md.neighbors, pos.R())
	}
	if !b.HorizWalls.Get(pos) && pos.Row == 0 && curWalker {
		md.val = dist+1
	} else if !b.HorizWalls.Get(pos) && pos.Row != 0 && md.distances.Get(pos.D()) == 0 {
		md.Set(pos.D(), dist+1)
		md.neighbors = append(md.neighbors, pos.D())
	}
	if !b.HorizWalls.Get(pos.U()) && pos.Row == b.NRows-1 && !curWalker {
		md.val = dist+1
	} else if !b.HorizWalls.Get(pos.U()) && pos.Row != b.NRows-1 && md.distances.Get(pos.U()) == 0 {
		md.Set(pos.U(), dist+1)
		md.neighbors = append(md.neighbors, pos.U())
	}

}
