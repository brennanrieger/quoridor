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
	val0      int
	val1      int
}

func (md *ManhattanDistance) Val(b *board.Board) (float64, float64) {
	md.board = b
	md.distances = &util.ValMatrix{}

	md.distances.Init(b.NRows, b.NCols)
	md.neighbors = []*board.Pos{b.Pos0}
	md.distances.Set(b.Pos0, 0)
	for head := 0; md.val0 == 0; head++ {
		md.bfs(head, false)
	}
	md.distances.Show()

	md.distances.Init(b.NRows, b.NCols)
	md.neighbors = []*board.Pos{b.Pos1}
	md.distances.Set(b.Pos1, 0)
	for head := 0; md.val1 == 0; head++ {
		md.bfs(head, true)
	}

	return float64(md.val0), float64(md.val1)
}

func (md *ManhattanDistance) bfs(head int, curWalker bool) {
	pos := md.neighbors[head]
	dist := md.distances.Get(pos).(int)
	b := md.board

	if !b.VertiWalls.Get(pos) && pos.Col != 0 && md.distances.Get(pos.L()) == nil {
		md.distances.Set(pos.L(), dist+1)
		md.neighbors = append(md.neighbors, pos.L())
	}
	if !b.VertiWalls.Get(pos.R()) && pos.Col != b.NCols-1 && md.distances.Get(pos.R()) == nil {
		md.distances.Set(pos.R(), dist+1)
		md.neighbors = append(md.neighbors, pos.R())
	}
	if !b.HorizWalls.Get(pos) && pos.Row == 0 && curWalker {
		md.val1 = dist + 1
	} else if !b.HorizWalls.Get(pos) && pos.Row != 0 && md.distances.Get(pos.D()) == nil {
		md.distances.Set(pos.D(), dist+1)
		md.neighbors = append(md.neighbors, pos.D())
	}
	if !b.HorizWalls.Get(pos.U()) && pos.Row == b.NRows-1 && !curWalker {
		md.val0 = dist + 1
	} else if !b.HorizWalls.Get(pos.U()) && pos.Row != b.NRows-1 && md.distances.Get(pos.U()) == nil {
		md.distances.Set(pos.U(), dist+1)
		md.neighbors = append(md.neighbors, pos.U())
	}

}
