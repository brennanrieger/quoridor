package util

import (
	"board"
)

func AvailableMoves(b *board.Board, playerNum bool) []*board.Move {
	var horizWallMoves = availableWalls(b, true)
	var vertiWallMoves = availableWalls(b, false)
	var allWallMoves = append(horizWallMoves, vertiWallMoves...)
	return append(allWallMoves, availableStepMoves(b, playerNum)...)
}

func availableWalls(b *board.Board, horizontal bool) []*board.Move {
	var availableWalls []*board.Move

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
			var move = &board.Move{
				Mt:  moveType,
				Pos: pos,
			}
			if err := boardCopy.MakeMove(move, true, dummyWinCh); err == nil && boardCopy.Validate() {
				availableWalls = append(availableWalls, move)
			}
		}
	}

	return availableWalls
}

func availableStepMoves(b *board.Board, playerNum bool) []*board.Move {
	var availableMoves []*board.Move

	var curPos *board.Pos
	if playerNum {
		curPos = b.Pos1
	} else {
		curPos = b.Pos0
	}

	var move = &board.Move{
		Mt:  0, // dummy value
		Pos: curPos,
	}

	if !b.HorizWalls.Get(curPos) && (curPos.Row != 0 || playerNum) {
		move.Mt = board.Down
		availableMoves = append(availableMoves, move.Copy())
	}
	if !b.HorizWalls.Get(curPos.U()) && (curPos.Row != b.NRows-1 || !playerNum) {
		move.Mt = board.Up
		availableMoves = append(availableMoves, move.Copy())
	}
	if !b.VertiWalls.Get(curPos) && curPos.Col != 0 {
		move.Mt = board.Left
		availableMoves = append(availableMoves, move.Copy())
	}
	if !b.VertiWalls.Get(curPos.R()) && curPos.Col != b.NCols-1 {
		move.Mt = board.Right
		availableMoves = append(availableMoves, move.Copy())
	}

	return availableMoves
}
