package main

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
	n_cols int
	n_rows int

	pos0 *Pos
	pos1 *Pos

	vertiWalls *Matrix
	horizWalls *Matrix
}

func (b *Board) Init(n_rows int, n_cols int) {
	b.n_rows = n_rows
	b.n_cols = n_cols

	b.pos0 = &Pos{
		r: 0,
		c: n_cols / 2,
	}

	b.pos1 = &Pos{
		r: n_rows - 1,
		c: n_cols / 2,
	}

	b.vertiWalls = &Matrix{}
	b.horizWalls = &Matrix{}
	b.vertiWalls.Init(n_rows, n_cols+1)
	b.horizWalls.Init(n_rows+1, n_cols)
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
		curPos = b.pos1
	} else {
		curPos = b.pos0
	}

	switch moveType {
	case HorizWall:
		if wallPos.r < 0 || wallPos.c < 0 || wallPos.r > b.n_rows || wallPos.c > b.n_cols-2 {
			return fmt.Errorf("wall out of bounds")
		} else if b.horizWalls.Get(wallPos) || b.horizWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall already exists")
		} else if wallPos.r != 0 && wallPos.r != b.n_rows && b.vertiWalls.Get(wallPos.D().R()) && b.vertiWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall intersects")
		} else {
			b.horizWalls.Set(wallPos)
			b.horizWalls.Set(wallPos.R())
		}
	case VertiWall:
		if wallPos.r < 0 || wallPos.c < 1 || wallPos.r > b.n_rows-2 || wallPos.c > b.n_cols-1 { // do not allow columns on far edges
			return fmt.Errorf("wall out of bounds")
		} else if b.vertiWalls.Get(wallPos) || b.vertiWalls.Get(wallPos.U()) {
			return fmt.Errorf("wall already exists")
		} else if b.horizWalls.Get(wallPos.U().L()) && b.horizWalls.Get(wallPos.U()) {
			return fmt.Errorf("wall intersects")
		} else {
			b.vertiWalls.Set(wallPos)
			b.vertiWalls.Set(wallPos.U())
		}
	case Down:
		if b.horizWalls.Get(curPos) {
			return fmt.Errorf("hit bottom wall")
		} else if curPos.r == 0 && curPlayer {
			win <- true
		} else if curPos.r == 0 && !curPlayer {
			return fmt.Errorf("hit floor")
		} else {
			if curPlayer {
				b.pos1 = curPos.D()
			} else {
				b.pos0 = curPos.D()
			}
		}
	case Up:
		if b.horizWalls.Get(curPos.U()) {
			return fmt.Errorf("hit top wall")
		} else if curPos.r == b.n_rows-1 && curPlayer {
			return fmt.Errorf("hit ceiling")
		} else if curPos.r == b.n_rows-1 && !curPlayer {
			win <- false
		} else {
			if curPlayer {
				b.pos1 = curPos.U()
			} else {
				b.pos0 = curPos.U()
			}
		}
	case Left:
		if b.vertiWalls.Get(curPos) {
			return fmt.Errorf("hit left wall")
		} else if curPos.c == 0 {
			return fmt.Errorf("hit left border")
		} else {
			if curPlayer {
				b.pos1 = curPos.L()
			} else {
				b.pos0 = curPos.L()
			}
		}
	case Right:
		if b.vertiWalls.Get(curPos.R()) {
			return fmt.Errorf("hit right wall")
		} else if curPos.c == b.n_cols-1 {
			return fmt.Errorf("hit right border")
		} else {
			if curPlayer {
				b.pos1 = curPos.R()
			} else {
				b.pos0 = curPos.R()
			}
		}
	default:
		return fmt.Errorf("Not a valid move type")
	}

	return nil
}

func (b *Board) Copy() *Board {
	newBoard := &Board{}
	newBoard.Init(b.n_rows, b.n_cols)
	newBoard.pos1 = b.pos1.Copy()
	newBoard.pos0 = b.pos0.Copy()
	newBoard.vertiWalls = b.vertiWalls.Copy()
	newBoard.horizWalls = b.horizWalls.Copy()
	return newBoard
}

func (b *Board) Flip() *Board {
	newBoard := &Board{}
	newBoard.Init(b.n_rows, b.n_cols)
	newBoard.pos1 = b.flipPos(b.pos0)
	newBoard.pos0 = b.flipPos(b.pos1)
	newBoard.vertiWalls = b.vertiWalls.Flip()
	newBoard.horizWalls = b.horizWalls.Flip()
	return newBoard
}

func (b *Board) flipPos(pos *Pos) *Pos {
	return &Pos{
		r: b.n_rows - pos.r - 1,
		c: b.n_cols - pos.c - 1,
	}
}

func (b *Board) Validate() bool {
	var visited0 = &Matrix{}
	var visited1 = &Matrix{}
	visited0.Init(b.n_rows, b.n_cols)
	visited1.Init(b.n_rows, b.n_cols)
	return b.walk(b.pos0, visited0, false) && b.walk(b.pos1, visited1, true)
}

func (b *Board) walk(pos *Pos, visited *Matrix, curWalker bool) bool {
	visited.Set(pos)

	var neighbors []*Pos
	if !b.vertiWalls.Get(pos) && pos.c != 0 {
		neighbors = append(neighbors, pos.L())
	}
	if !b.vertiWalls.Get(pos.R()) && pos.c != b.n_cols-1 {
		neighbors = append(neighbors, pos.R())
	}
	if !b.horizWalls.Get(pos) && pos.r == 0 && curWalker {
		return true
	} else if !b.horizWalls.Get(pos) && pos.r != 0 {
		neighbors = append(neighbors, pos.D())
	}
	if !b.horizWalls.Get(pos.U()) && pos.r == b.n_rows-1 && !curWalker {
		return true
	} else if !b.horizWalls.Get(pos.U()) && pos.r != b.n_rows-1 {
		neighbors = append(neighbors, pos.U())
	}

	for _, neighborPos := range neighbors {
		if !visited.Get(neighborPos) && b.walk(neighborPos, visited, curWalker) {
			return true
		}
	}
	return false
}
