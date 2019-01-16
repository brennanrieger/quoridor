package util

import (
	"board"
)

func AvailableMoves(b *board.Board, playerNum bool) []*board.Move {
	var wallMoves = availableWalls(b)
	return append(wallMoves, availableStepMoves(b, playerNum)...)
}

func availableWalls(b *board.Board) []*board.Move {
	var availableWalls []*board.Move

	for r := 0; r < b.NRows-1; r++ {
		for c := 0; c < b.NCols-1; c++ {
			pos := &board.Pos{
				Row: r,
				Col: c,
			}

			// horizontal wall
			var move = &board.Move{
				Mt:  board.HorizWall,
				Pos: pos,
			}
			if testMove(b, move) {
				availableWalls = append(availableWalls, move)
			}

			// vertical wall
			move.Mt = board.VertiWall
			if testMove(b, move) {
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

	move.Mt = board.Down
	if testMove(b, move) {
		availableMoves = append(availableMoves, move.Copy())
	}
	move.Mt = board.Left
	if testMove(b, move) {
		availableMoves = append(availableMoves, move.Copy())
	}
	move.Mt = board.Up
	if testMove(b, move) {
		availableMoves = append(availableMoves, move.Copy())
	}
	move.Mt = board.Right
	if testMove(b, move) {
		availableMoves = append(availableMoves, move.Copy())
	}

	return availableMoves
}

func availableJumpMoves(b *board.Board, playerNum bool) []*board.Move {
	var availableMoves []*board.Move

	for

	return availableMoves
}

func testMove(b *board.Board, move *board.Move) bool {
	var boardCopy = b.Copy()
	dummyWinCh := make(chan bool, 2)
	err := boardCopy.MakeMove(move, true, dummyWinCh)
	return err == nil && boardCopy.Validate()
}
