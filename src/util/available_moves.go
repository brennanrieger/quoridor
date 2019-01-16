package util

import (
	"board"
)

func AvailableMoves(b *board.Board, playerNum bool) []*board.Move {
	var wallMoves = availableWallMoves(b)
	var stepMoves = availableStepMoves(b, playerNum)
	var jumpMoves = availableJumpMoves(b, playerNum)
	return append(append(wallMoves, stepMoves...), jumpMoves...)
}

func availableWallMoves(b *board.Board) []*board.Move {
	var availableMoves []*board.Move

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
			addMoveIfValid(b, move, availableMoves)

			// vertical wall
			move.Mt = board.VertiWall
			addMoveIfValid(b, move, availableMoves)
		}
	}

	return availableMoves
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
	addMoveIfValid(b, move, availableMoves)

	move.Mt = board.Left
	addMoveIfValid(b, move, availableMoves)

	move.Mt = board.Up
	addMoveIfValid(b, move, availableMoves)

	move.Mt = board.Right
	addMoveIfValid(b, move, availableMoves)

	return availableMoves
}

func availableJumpMoves(b *board.Board, playerNum bool) []*board.Move {
	var availableMoves []*board.Move

	var enemyPos *Pos
	if playerNum {
		enemyPos = b.Pos0
	} else {
		enemyPos = b.pos1
	}

	for _, futurePos := range b.Neighbors(enemyPos) {
		var move = &board.Move{
			Mt:  board.Jump,
			Pos: futurePos,
		}
		addMoveIfValid(b, move, availableMoves)
	}

	return availableMoves
}

func addMoveIfValid(b *board.Board, move *board.Move, availableMoves []*board.Move) bool {
	var boardCopy = b.Copy()
	dummyWinCh := make(chan bool, 2)
	if err := boardCopy.MakeMove(move, true, dummyWinCh); err == nil && boardCopy.Validate() {
		availableMoves = append(availableMoves, move.Copy())
	}
}
