package visualizer

import "board"

// A visualizer is capable of displaying the board
type Visualizer interface {
	// display a visualization of the board
	Display(b *board.Board)
}
