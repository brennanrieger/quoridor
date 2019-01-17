package util

import (
	"board"
)

func AvailableMoves(b *board.Board, playerNum bool) []*board.Move {
	av := make([]*board.Move, 0)
	availableMoves := &av

	addAvailableWallMoves(b, playerNum, availableMoves)
	addAvailableStepMoves(b, playerNum, availableMoves)
	addAvailableJumpMoves(b, playerNum, availableMoves)
	return *availableMoves
}

func addAvailableWallMoves(b *board.Board, playerNum bool, availableMoves *[]*board.Move) {
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
			addMoveIfValid(b, move, playerNum, availableMoves)

			// vertical wall
			move.Mt = board.VertiWall
			addMoveIfValid(b, move, playerNum, availableMoves)
		}
	}
}

func addAvailableStepMoves(b *board.Board, playerNum bool, availableMoves *[]*board.Move) {
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
	addMoveIfValid(b, move, playerNum, availableMoves)

	move.Mt = board.Left
	addMoveIfValid(b, move, playerNum, availableMoves)

	move.Mt = board.Up
	addMoveIfValid(b, move, playerNum, availableMoves)

	move.Mt = board.Right
	addMoveIfValid(b, move, playerNum, availableMoves)
}

func addAvailableJumpMoves(b *board.Board, playerNum bool, availableMoves *[]*board.Move) {
	var enemyPos *board.Pos
	if playerNum {
		enemyPos = b.Pos0
	} else {
		enemyPos = b.Pos1
	}

	for _, futurePos := range b.Neighbors(enemyPos) {
		var move = &board.Move{
			Mt:  board.Jump,
			Pos: futurePos,
		}
		addMoveIfValid(b, move, playerNum, availableMoves)
	}
}

func addMoveIfValid(b *board.Board, move *board.Move, playerNum bool, availableMoves *[]*board.Move) {
	var boardCopy = b.Copy()
	dummyWinCh := make(chan bool, 2)
	if err := boardCopy.MakeMove(move, playerNum, dummyWinCh); err == nil && boardCopy.Validate() {
		*availableMoves = append(*availableMoves, move.Copy())
	}
}
