package player

import (
	"board"
	"feature"
	"math/rand"
)

type FeaturePlayer struct {
	playerNum bool
}

func (fp *FeaturePlayer) Init(playerNum bool) {
	fp.playerNum = playerNum
}

func (fp *FeaturePlayer) Move(b *board.Board) (board.MoveType, *board.Pos) {
	var boardCopy = b.Copy()
	var option = rand.Intn(3)
	var mt board.MoveType
	var pos *board.Pos
	if option == 1 {
		mt, pos = fp.makeWall(boardCopy, true)
	} else if option == 2 {
		mt, pos = fp.makeWall(boardCopy, false)
	} else {
		mt, pos = fp.movePiece(b), &board.Pos{}
	}
	return mt, pos
}

func (fp *FeaturePlayer) makeWall(b *board.Board, horizontal bool) (board.MoveType, *board.Pos) {
	var availablePositions []*board.Pos

	var moveType board.MoveType
	if horizontal {
		moveType = board.HorizWall
	} else {
		moveType = board.VertiWall
	}

	for r := 0; r < b.NRows-1; r++ {
		for c := 0; c < b.NCols-1; c++ {
			var boardCopy = b.Copy()
			pos := &board.Pos{
				Row: r,
				Col: c,
			}
			dummyWinCh := make(chan bool, 2)
			if err := boardCopy.Move(moveType, pos, fp.playerNum, dummyWinCh); err == nil && boardCopy.Validate() {
				availablePositions = append(availablePositions, pos)
			}
		}
	}

	if len(availablePositions) > 0 {
		return moveType, availablePositions[rand.Intn(len(availablePositions))]
	} else {
		return fp.movePiece(b), nil
	}
}

func (fp *FeaturePlayer) movePiece(b *board.Board) board.MoveType {
	var availableMoves []board.MoveType

	var curPos *board.Pos
	if fp.playerNum {
		curPos = b.Pos1
	} else {
		curPos = b.Pos0
	}

	if !b.HorizWalls.Get(curPos) && (curPos.Row != 0 || fp.playerNum == true) {
		availableMoves = append(availableMoves, board.Down)
	}
	if !b.HorizWalls.Get(curPos.U()) && (curPos.Row != b.NRows-1 || fp.playerNum == false) {
		availableMoves = append(availableMoves, board.Up)
	}
	if !b.VertiWalls.Get(curPos) && curPos.Col != 0 {
		availableMoves = append(availableMoves, board.Left)
	}
	if !b.VertiWalls.Get(curPos.R()) && curPos.Col != b.NCols-1 {
		availableMoves = append(availableMoves, board.Right)
	}

	bestVal := 0
	bestI := 0
	for i, _ := range availableMoves {
		md := &feature.ManhattanDistance{}
		bNew := b.Copy()
		bNew.Pos0 = bNew.Pos0.R() // TODO: make this real
		val := md.Val(bNew)
		if val > bestVal {
			bestVal = val
			bestI = i
		}
	}
	return availableMoves[bestI]
}
