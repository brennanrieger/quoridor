package main

import "fmt"

type AsciiVisualizer struct{}

var (
	boxChars = [16]string{"·", "╵", "╶", "└", "╷", "│", "┌", "├", "╴", "┘", "─", "┴", "┐", "┤", "┬", "┼"}
)

func (av *AsciiVisualizer) Display(b *Board) {
	var disp string
	for r := b.n_rows; r >= 0; r-- {
		// no gapRow before first lineRow
		if r != b.n_rows {
			disp += av.gapRow(b, r) + "\n"
		}
		disp += av.lineRow(b, r) + "\n"
	}
	disp += "\n"
	fmt.Println(disp)
}

func (av *AsciiVisualizer) lineRow(b *Board, r int) string {
	var lineRow string
	for c := 0; c < b.n_cols+1; c++ {
		pos := &Pos{
			r: r,
			c: c,
		}
		lineRow += av.intersectionChar(b, pos)

		// no horizChar following last intersectionChar
		if c != b.n_cols {
			lineRow += av.horizChar(b, pos)
		}
	}
	return lineRow
}

func (av *AsciiVisualizer) gapRow(b *Board, r int) string {
	var gapRow string
	for c := 0; c < b.n_cols+1; c++ {
		pos := &Pos{
			r: r,
			c: c,
		}
		gapRow += av.vertiChar(b, pos)

		// no midChar following last vertiChar
		if c != b.n_cols {
			gapRow += av.midChar(b, pos)
		}
	}
	return gapRow
}

func (av *AsciiVisualizer) intersectionChar(b *Board, pos *Pos) string {
	var up bool
	if pos.r < b.n_rows {
		up = b.vertiWalls.Get(pos)
	}

	var right bool
	if pos.c < b.n_cols {
		right = b.horizWalls.Get(pos)
	}

	var down bool
	if pos.r > 0 {
		down = b.vertiWalls.Get(pos.D())
	}

	var left bool
	if pos.c > 0 {
		left = b.horizWalls.Get(pos.L())
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

func (av *AsciiVisualizer) horizChar(b *Board, pos *Pos) string {
	if b.horizWalls.Get(pos) {
		return "─"
	} else {
		return " "
	}
}

func (av *AsciiVisualizer) vertiChar(b *Board, pos *Pos) string {
	if b.vertiWalls.Get(pos) {
		return "│"
	} else {
		return " "
	}
}

func (av *AsciiVisualizer) midChar(b *Board, pos *Pos) string {
	if b.pos0.Equal(pos) {
		return "0"
	} else if b.pos1.Equal(pos) {
		return "1"
	} else {
		return " "
	}
}
