package util

import (
	"board"
)

func AvailableMoves(b *board.Board) []*board.Move {
	av := make([]*board.Move, 0)
	availableMoves := &av

	addAvailableWallMoves(b, availableMoves)
	addAvailableStepMoves(b, availableMoves)
	addAvailableJumpMoves(b, availableMoves)
	return *availableMoves
}

func addAvailableWallMoves(b *board.Board, availableMoves *[]*board.Move) {
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
}

func addAvailableStepMoves(b *board.Board, availableMoves *[]*board.Move) {
	var move *board.Move

	move = board.StepMove(board.Down)
	addMoveIfValid(b, move, availableMoves)

	move = board.StepMove(board.Up)
	addMoveIfValid(b, move, availableMoves)

	move = board.StepMove(board.Left)
	addMoveIfValid(b, move, availableMoves)

	move = board.StepMove(board.Right)
	addMoveIfValid(b, move, availableMoves)
}

func addAvailableJumpMoves(b *board.Board, availableMoves *[]*board.Move) {
	var enemyPos *board.Pos
	if b.CurPlayer {
		enemyPos = b.Pos0
	} else {
		enemyPos = b.Pos1
	}

	for _, futurePos := range b.Neighbors(enemyPos) {
		var move = &board.Move{
			Mt:  board.Jump,
			Pos: futurePos,
		}
		addMoveIfValid(b, move, availableMoves)
	}
}

func addMoveIfValid(b *board.Board, move *board.Move, availableMoves *[]*board.Move) {
	var boardCopy = b.Copy()
	if err := boardCopy.MakeMove(move); err == nil && boardCopy.Validate() {
		*availableMoves = append(*availableMoves, move.Copy())
	}
}
