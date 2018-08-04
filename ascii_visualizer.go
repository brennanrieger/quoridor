package main

import (
	"board"
	"fmt"
	"strings"
)

type AsciiVisualizer struct{}

var (
	boxChars = [16]string{"·", "╵", "╶", "└", "╷", "│", "┌", "├", "╴", "┘", "─", "┴", "┐", "┤", "┬", "┼"}
)

func (av *AsciiVisualizer) Display(b *board.Board) {
	var disp string
	for r := b.NRows; r >= 0; r-- {
		// no gapRow before first lineRow
		if r != b.NRows {
			disp += av.gapRow(b, r) + "\n"
		}
		disp += av.lineRow(b, r) + "\n"
	}
	disp += "\n\n  "
	disp += strings.Repeat("=", b.NCols*2-3)
	disp += "\n\n"
	fmt.Println(disp)
}

func (av *AsciiVisualizer) lineRow(b *board.Board, r int) string {
	var lineRow string
	for c := 0; c < b.NCols+1; c++ {
		pos := &board.Pos{
			r: r,
			c: c,
		}
		lineRow += av.intersectionChar(b, pos)

		// no horizChar following last intersectionChar
		if c != b.NCols {
			lineRow += av.horizChar(b, pos)
		}
	}
	return lineRow
}

func (av *AsciiVisualizer) gapRow(b *board.Board, r int) string {
	var gapRow string
	for c := 0; c < b.NCols+1; c++ {
		pos := &board.Pos{
			r: r,
			c: c,
		}
		gapRow += av.vertiChar(b, pos)

		// no midChar following last vertiChar
		if c != b.NCols {
			gapRow += av.midChar(b, pos)
		}
	}
	return gapRow
}

func (av *AsciiVisualizer) intersectionChar(b *board.Board, pos *board.Pos) string {
	var up bool
	if pos.r < b.NRows {
		up = b.VertiWalls.Get(pos)
	}

	var right bool
	if pos.c < b.NCols {
		right = b.HorizWalls.Get(pos)
	}

	var down bool
	if pos.r > 0 {
		down = b.VertiWalls.Get(pos.D())
	}

	var left bool
	if pos.c > 0 {
		left = b.HorizWalls.Get(pos.L())
	}

	charIdx := 0
	if up {
		charIdx += 1
	}
	if right {
		charIdx += 2
	}
	if down {
		charIdx += 4
	}
	if left {
		charIdx += 8
	}
	return boxChars[charIdx]
}

func (av *AsciiVisualizer) horizChar(b *board.Board, pos *board.Pos) string {
	if b.HorizWalls.Get(pos) {
		return "─"
	} else {
		return " "
	}
}

func (av *AsciiVisualizer) vertiChar(b *board.Board, pos *board.Pos) string {
	if b.VertiWalls.Get(pos) {
		return "│"
	} else {
		return " "
	}
}

func (av *AsciiVisualizer) midChar(b *board.Board, pos *board.Pos) string {
	if b.Pos0.Equal(pos) {
		return "0"
	} else if b.Pos1.Equal(pos) {
		return "1"
	} else {
		return " "
	}
}
