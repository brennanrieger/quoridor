package feature

import "board"

// A feature is an evaluation of the static board
type Feature interface {
	// return the value for player0, player1
	Val(b *board.Board) (float64, float64)
}
