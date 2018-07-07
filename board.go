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

	curPlayer bool

	vertiWalls *Matrix
	horizWalls *Matrix

	win chan bool
}

func (b *Board) Init(n_rows int, n_cols int, win chan bool) {
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

	b.win = win
}

func (b *Board) Move(moveType MoveType, wallPos *Pos) error {
	var boardCopy = b.Copy()
	if err := boardCopy.move(moveType, wallPos); err != nil {
		return err
	} else if !boardCopy.Validate() {
		return fmt.Errorf("New board is not valid")
	} else {
		b.move(moveType, wallPos)
		b.curPlayer = !b.curPlayer
	}
	return nil
}

func (b *Board) move(moveType MoveType, wallPos *Pos) error {

	var curPos *Pos
	if b.curPlayer {
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
		} else if curPos.r == 0 && b.curPlayer {
			b.win <- true
		} else if curPos.r == 0 && !b.curPlayer {
			return fmt.Errorf("hit floor")
		} else {
			if b.curPlayer {
				b.pos1 = curPos.D()
			} else {
				b.pos0 = curPos.D()
			}
		}
	case Up:
		if b.horizWalls.Get(curPos.U()) {
			return fmt.Errorf("hit top wall")
		} else if curPos.r == b.n_rows-1 && b.curPlayer {
			return fmt.Errorf("hit ceiling")
		} else if curPos.r == b.n_rows-1 && !b.curPlayer {
			b.win <- false
		} else {
			if b.curPlayer {
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
			if b.curPlayer {
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
			if b.curPlayer {
				b.pos1 = curPos.R()
			} else {
				b.pos0 = curPos.R()
			}
		}
	default:
		return fmt.Errorf("Not a valid move type")
	}

	// only switch player if move was successful
	b.curPlayer = !b.curPlayer
	return nil
}

func (b *Board) Copy() *Board {
	newBoard := &Board{}
	newBoard.Init(b.n_rows, b.n_cols, b.win)
	newBoard.curPlayer = b.curPlayer
	newBoard.pos1 = b.pos1.Copy()
	newBoard.pos0 = b.pos0.Copy()
	newBoard.vertiWalls = b.vertiWalls.Copy()
	newBoard.horizWalls = b.horizWalls.Copy()
	return newBoard
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
