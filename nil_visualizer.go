package main

import "board"

type NilVisualizer struct{}

func (nv *NilVisualizer) Display(b *board.Board) {
	return
}
