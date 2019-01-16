package board

import "fmt"

type Board struct {
	NCols int
	NRows int

	Pos0 *Pos
	Pos1 *Pos

	VertiWalls *Matrix
	HorizWalls *Matrix
}

func (b *Board) Init(nRows int, nCols int) {
	b.NRows = nRows
	b.NCols = nCols

	b.Pos0 = &Pos{
		Row: 0,
		Col: nCols / 2,
	}

	b.Pos1 = &Pos{
		Row: nRows - 1,
		Col: nCols / 2,
	}

	b.VertiWalls = &Matrix{}
	b.HorizWalls = &Matrix{}
	b.VertiWalls.Init(nRows, nCols+1)
	b.HorizWalls.Init(nRows+1, nCols)
}

func (b *Board) MakeMove(move *Move, curPlayer bool, win chan bool) error {
	var boardCopy = b.Copy()
	dummyWinCh := make(chan bool, 2)
	if err := boardCopy.makeMove(move, curPlayer, dummyWinCh); err != nil {
		return err
	} else if !boardCopy.Validate() {
		return fmt.Errorf("New board is not valid")
	} else {
		b.makeMove(move, curPlayer, win)
	}
	return nil
}

func (b *Board) makeMove(move *Move, curPlayer bool, win chan bool) error {
	var moveType = move.Mt
	var wallPos = move.Pos

	var curPos *Pos
	if curPlayer {
		curPos = b.Pos1
	} else {
		curPos = b.Pos0
	}

	switch moveType {
	case HorizWall:
		if wallPos.Row < 0 || wallPos.Col < 0 || wallPos.Row > b.NRows || wallPos.Col > b.NCols-2 {
			return fmt.Errorf("wall out of bounds")
		} else if b.HorizWalls.Get(wallPos) || b.HorizWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall already exists")
		} else if wallPos.Row != 0 && wallPos.Row != b.NRows && b.VertiWalls.Get(wallPos.D().R()) && b.VertiWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall intersects")
		} else {
			b.HorizWalls.Set(wallPos)
			b.HorizWalls.Set(wallPos.R())
		}
	case VertiWall:
		if wallPos.Row < 0 || wallPos.Col < 1 || wallPos.Row > b.NRows-2 || wallPos.Col > b.NCols-1 { // do not allow columns on far edges
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
		} else if curPos.Row == 0 && curPlayer {
			win <- true
		} else if curPos.Row == 0 && !curPlayer {
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
		} else if curPos.Row == b.NRows-1 && curPlayer {
			return fmt.Errorf("hit ceiling")
		} else if curPos.Row == b.NRows-1 && !curPlayer {
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
		} else if curPos.Col == 0 {
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
		} else if curPos.Col == b.NCols-1 {
			return fmt.Errorf("hit right border")
		} else {
			if curPlayer {
				b.Pos1 = curPos.R()
			} else {
				b.Pos0 = curPos.R()
			}
		}
	case Jump:
		futurePos = move.Pos
		if curPlayer {
			enemyPos = b.Pos0
		} else {
			enemyPos = b.Pos1
		}
		if futurePos.Equal(curPos) {
			return fmt.Errorf("cannot jump to current position")
		} else if (!b.neighbors(curPos, enemyPos)) {
			return fmt.Errorf("the two players must be touching to perform a jump")
		} else if (!b.neighbors(enemyPos, futurePos)) {
			return fmt.Errorf("the destination space must be next to the opponent")
		} else if () {

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
		Row: b.NRows - pos.Row - 1,
		Col: b.NCols - pos.Col - 1,
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
	if !b.VertiWalls.Get(pos) && pos.Col != 0 {
		neighbors = append(neighbors, pos.L())
	}
	if !b.VertiWalls.Get(pos.R()) && pos.Col != b.NCols-1 {
		neighbors = append(neighbors, pos.R())
	}
	if !b.HorizWalls.Get(pos) && pos.Row == 0 && curWalker {
		return true
	} else if !b.HorizWalls.Get(pos) && pos.Row != 0 {
		neighbors = append(neighbors, pos.D())
	}
	if !b.HorizWalls.Get(pos.U()) && pos.Row == b.NRows-1 && !curWalker {
		return true
	} else if !b.HorizWalls.Get(pos.U()) && pos.Row != b.NRows-1 {
		neighbors = append(neighbors, pos.U())
	}

	for _, neighborPos := range neighbors {
		if !visited.Get(neighborPos) && b.walk(neighborPos, visited, curWalker) {
			return true
		}
	}
	return false
}

// Checks if pos2 can be reached from pos1 in one move
func (b *Board) neighbors(pos1 *Pos, pos2 *Pos) bool {
	if pos1.Row == pos2.Row {
		if pos1.Col+1 == pos2.Col {
			return board.VertiWalls(pos2)
		} else if pos2.Col+1 == pos1.Col {
			return board.VertiWalls(pos1)
		}
	} else if pos2.Col == pos2.Col {
		if pos1.Row+1 == pos2.Row {
			return board.HorizWalls(pos2)
		} else if pos2.Row+1 == pos1.Row {
			return board.HorizWalls(pos1)
		}
	}
	// Not on adjacent square
	return false
}

