package board

import (
	"fmt"
)

type MoveType int

const (
	HorizWall      = 1
	VertiWall      = 2
	Up             = 3
	Down           = 4
	Right          = 5
	Left           = 6
	JumpUpUp       = 7
	JumpUpLeft     = 8
	JumpUpRight    = 9
	JumpDownDown   = 10
	JumpDownLeft   = 11
	JumpDownRight  = 12
	JumpRightRight = 13
	JumpRightUp    = 14
	JumpRightDown  = 15
	JumpLeftLeft   = 16
	JumpLeftUp     = 17
	JumpLeftDown   = 18
)

type Board struct {
	NCols int
	NRows int

	Pos0 *Pos
	Pos1 *Pos

	VertiWalls *Matrix
	HorizWalls *Matrix
}

func (b *Board) Init(NRows int, NCols int) {
	b.NRows = NRows
	b.NCols = NCols

	b.Pos0 = &Pos{
		r: 0,
		c: NCols / 2,
	}

	b.Pos1 = &Pos{
		r: NRows - 1,
		c: NCols / 2,
	}

	b.VertiWalls = &Matrix{}
	b.HorizWalls = &Matrix{}
	b.VertiWalls.Init(NRows, NCols+1)
	b.HorizWalls.Init(NRows+1, NCols)
}

func (b *Board) Move(moveType MoveType, wallPos *Pos, curPlayer bool, win chan bool) error {
	var boardCopy = b.Copy()
	dummyWinCh := make(chan bool, 2)
	if err := boardCopy.move(moveType, wallPos, curPlayer, dummyWinCh); err != nil {
		return err
	} else if !boardCopy.Validate() {
		return fmt.Errorf("New board is not valid")
	} else {
		b.move(moveType, wallPos, curPlayer, win)
	}
	return nil
}

func (b *Board) move(moveType MoveType, wallPos *Pos, curPlayer bool, win chan bool) error {

	var curPos *Pos
	if curPlayer {
		curPos = b.Pos1
	} else {
		curPos = b.Pos0
	}

	switch moveType {
	case HorizWall:
		if wallPos.r < 0 || wallPos.c < 0 || wallPos.r > b.NRows || wallPos.c > b.NCols-2 {
			return fmt.Errorf("wall out of bounds")
		} else if b.HorizWalls.Get(wallPos) || b.HorizWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall already exists")
		} else if wallPos.r != 0 && wallPos.r != b.NRows && b.VertiWalls.Get(wallPos.D().R()) && b.VertiWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall intersects")
		} else {
			b.HorizWalls.Set(wallPos)
			b.HorizWalls.Set(wallPos.R())
		}
	case VertiWall:
		if wallPos.r < 0 || wallPos.c < 1 || wallPos.r > b.NRows-2 || wallPos.c > b.NCols-1 { // do not allow columns on far edges
			return fmt.Errorf("wall out of bounds")
		} else if b.VertiWalls.Get(wallPos) || b.VertiWalls.Get(wallPos.U()) {
			return fmt.Errorf("wall already exists")
		} else if b.HorizWalls.Get(wallPos.U().L()) && b.HorizWalls.Get(wallPos.U()) {
			return fmt.Errorf("wall intersects")
		} else {
			b.VertiWalls.Set(wallPos)
			b.VertiWalls.Set(wallPos.U())
		}
	case Down:
		if b.HorizWalls.Get(curPos) {
			return fmt.Errorf("hit bottom wall")
		} else if curPos.r == 0 && curPlayer {
			win <- true
		} else if curPos.r == 0 && !curPlayer {
			return fmt.Errorf("hit floor")
		} else {
			if curPlayer {
				b.Pos1 = curPos.D()
			} else {
				b.Pos0 = curPos.D()
			}
		}
	case Up:
		if b.HorizWalls.Get(curPos.U()) {
			return fmt.Errorf("hit top wall")
		} else if curPos.r == b.NRows-1 && curPlayer {
			return fmt.Errorf("hit ceiling")
		} else if curPos.r == b.NRows-1 && !curPlayer {
			win <- false
		} else {
			if curPlayer {
				b.Pos1 = curPos.U()
			} else {
				b.Pos0 = curPos.U()
			}
		}
	case Left:
		if b.VertiWalls.Get(curPos) {
			return fmt.Errorf("hit left wall")
		} else if curPos.c == 0 {
			return fmt.Errorf("hit left border")
		} else {
			if curPlayer {
				b.Pos1 = curPos.L()
			} else {
				b.Pos0 = curPos.L()
			}
		}
	case Right:
		if b.VertiWalls.Get(curPos.R()) {
			return fmt.Errorf("hit right wall")
		} else if curPos.c == b.NCols-1 {
			return fmt.Errorf("hit right border")
		} else {
			if curPlayer {
				b.Pos1 = curPos.R()
			} else {
				b.Pos0 = curPos.R()
			}
		}
	default:
		return fmt.Errorf("Not a valid move type")
	}

	return nil
}

func (b *Board) Copy() *Board {
	newBoard := &Board{}
	newBoard.Init(b.NRows, b.NCols)
	newBoard.Pos1 = b.Pos1.Copy()
	newBoard.Pos0 = b.Pos0.Copy()
	newBoard.VertiWalls = b.VertiWalls.Copy()
	newBoard.HorizWalls = b.HorizWalls.Copy()
	return newBoard
}

func (b *Board) Flip() *Board {
	newBoard := &Board{}
	newBoard.Init(b.NRows, b.NCols)
	newBoard.Pos1 = b.flipPos(b.Pos0)
	newBoard.Pos0 = b.flipPos(b.Pos1)
	newBoard.VertiWalls = b.VertiWalls.Flip()
	newBoard.HorizWalls = b.HorizWalls.Flip()
	return newBoard
}

func (b *Board) flipPos(pos *Pos) *Pos {
	return &Pos{
		r: b.NRows - pos.r - 1,
		c: b.NCols - pos.c - 1,
	}
}

func (b *Board) Validate() bool {
	var visited0 = &Matrix{}
	var visited1 = &Matrix{}
	visited0.Init(b.NRows, b.NCols)
	visited1.Init(b.NRows, b.NCols)
	return b.walk(b.Pos0, visited0, false) && b.walk(b.Pos1, visited1, true)
}

func (b *Board) walk(pos *Pos, visited *Matrix, curWalker bool) bool {
	visited.Set(pos)

	var neighbors []*Pos
	if !b.VertiWalls.Get(pos) && pos.c != 0 {
		neighbors = append(neighbors, pos.L())
	}
	if !b.VertiWalls.Get(pos.R()) && pos.c != b.NCols-1 {
		neighbors = append(neighbors, pos.R())
	}
	if !b.HorizWalls.Get(pos) && pos.r == 0 && curWalker {
		return true
	} else if !b.HorizWalls.Get(pos) && pos.r != 0 {
		neighbors = append(neighbors, pos.D())
	}
	if !b.HorizWalls.Get(pos.U()) && pos.r == b.NRows-1 && !curWalker {
		return true
	} else if !b.HorizWalls.Get(pos.U()) && pos.r != b.NRows-1 {
		neighbors = append(neighbors, pos.U())
	}

	for _, neighborPos := range neighbors {
		if !visited.Get(neighborPos) && b.walk(neighborPos, visited, curWalker) {
			return true
		}
	}
	return false
}
