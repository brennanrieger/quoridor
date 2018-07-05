package main

// A player is capable of playing quoridor by specifying a move given the board setup
type Player interface {
	// return a move given the board
	Move(b *Board) (MoveType, *Pos)
}
