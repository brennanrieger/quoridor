package main

// A feature is an evaluate of the static board
type Feature interface {
	// return the value of the feature for each player
	Val(b *Board) float32
}
