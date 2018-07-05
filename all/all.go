package main

import (
	"fmt"
	"math/rand"
)

func main() {
	bob := "yo"
	bob += "go"
	fmt.Println(bob + "sj")

	av := AsciiVisualizer{}
	game := &Game{}
	p0 := RandomPlayer{}
	p1 := RandomPlayer{}
	game.Init(5, 5, p0, p1, av)

}

type AsciiVisualizer struct{}

var (
	boxChars = [16]string{"·", "╵", "╶", "└", "╷", "│", "┌", "├", "╴", "┘", "─", "┴", "┐", "┤", "┬", "┼"}
)

func (av *AsciiVisualizer) Display(b *Board) {
	var disp string
	for r := 0; r < b.n_rows+1; r++ {
		disp += av.lineRow(b, r) + "\n"

		// no gapRow following last lineRow
		if r != b.n_rows {
			disp += av.gapRow(b, r) + "\n"
		}
	}
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
	if b.pos0 == pos {
		return "0"
	} else if b.pos1 == pos {
		return "1"
	} else {
		return " "
	}
}

// A visualizer is capable of displaying the board
type Visualizer interface {
	// display a visualization of the board
	Display(b *Board)
}

type NilVisualizer struct{}

func (nv *NilVisualizer) Display(b *Board) {
	return
}

// A player is capable of playing quoridor by specifying a move given the board setup
type Player interface {
	// return a move given the board
	Move(b *Board) (MoveType, *Pos)
}

type RandomPlayer struct{}

func (rp *RandomPlayer) Move(b *Board) (MoveType, *Pos) {
	var boardCopy = b.Copy()
	var option = rand.Intn(3)
	if option == 1 {
		return HorizWall, rp.makeWall(boardCopy, true)
	} else if option == 2 {
		return VertiWall, rp.makeWall(boardCopy, false)
	} else {
		return rp.movePiece(b), nil
	}
}

func (rp *RandomPlayer) makeWall(b *Board, horizontal bool) *Pos {
	var availablePositions []*Pos

	var moveType MoveType
	if horizontal {
		moveType = HorizWall
	} else {
		moveType = VertiWall
	}

	for r := 0; r < b.n_rows-1; r++ {
		for c := 0; c < b.n_cols-1; c++ {
			var boardCopy = b.Copy()
			pos := &Pos{
				r: r,
				c: c,
			}
			if err := boardCopy.Move(moveType, pos); err == nil && boardCopy.Validate() {
				availablePositions = append(availablePositions, pos)
			}
		}
	}

	return availablePositions[rand.Intn(len(availablePositions))]
}

func (rp *RandomPlayer) movePiece(b *Board) MoveType {
	var availableMoves []MoveType

	var curPos *Pos
	if b.curPlayer {
		curPos = b.pos1
	} else {
		curPos = b.pos0
	}

	if !b.horizWalls.Get(curPos) && curPos.r != 0 {
		availableMoves = append(availableMoves, Down)
	}
	if !b.horizWalls.Get(curPos.U()) && curPos.r != b.n_rows-1 {
		availableMoves = append(availableMoves, Up)
	}
	if !b.horizWalls.Get(curPos) && curPos.c != 0 {
		availableMoves = append(availableMoves, Left)
	}
	if !b.horizWalls.Get(curPos.R()) && curPos.c != b.n_cols-1 {
		availableMoves = append(availableMoves, Right)
	}
	return availableMoves[rand.Intn(len(availableMoves))]
}

type MoveType int

const (
	HorizWall      = 1
	VertiWall      = 2
	Up             = 3
	Down           = 4
	Right          = 5
	Left           = 6
	JumpUpUp       = 7
	JumpUpLeft     = 8
	JumpUpRight    = 9
	JumpDownDown   = 10
	JumpDownLeft   = 11
	JumpDownRight  = 12
	JumpRightRight = 13
	JumpRightUp    = 14
	JumpRightDown  = 15
	JumpLeftLeft   = 16
	JumpLeftUp     = 17
	JumpLeftDown   = 18
)

type Board struct {
	n_cols int
	n_rows int

	pos0 *Pos
	pos1 *Pos

	curPlayer bool

	vertiWalls *Matrix
	horizWalls *Matrix

	win chan bool
}

func (b *Board) Init(n_rows int, n_cols int, win chan bool) {
	b.n_rows = n_rows
	b.n_cols = n_cols

	b.pos0 = &Pos{
		r: 0,
		c: n_cols / 2,
	}

	b.pos1 = &Pos{
		r: n_rows - 1,
		c: n_cols / 2,
	}

	b.vertiWalls.Init(n_rows, n_cols+1)
	b.horizWalls.Init(n_rows+1, n_cols)

	b.win = win
}

func (b *Board) Move(moveType MoveType, wallPos *Pos) error {
	var boardCopy = b.Copy()
	if err := boardCopy.move(moveType, wallPos); err != nil {
		return err
	} else if !boardCopy.Validate() {
		return fmt.Errorf("New board is not valid")
	} else {
		b.move(moveType, wallPos)
		b.curPlayer = !b.curPlayer
	}
	return nil
}

func (b *Board) move(moveType MoveType, wallPos *Pos) error {

	var curPos *Pos
	if b.curPlayer {
		curPos = b.pos1
	} else {
		curPos = b.pos0
	}

	switch moveType {
	case HorizWall:
		if wallPos.r < 0 || wallPos.c < 0 || wallPos.r > b.n_rows || wallPos.c > b.n_cols-2 {
			return fmt.Errorf("wall out of bounds")
		} else if b.horizWalls.Get(wallPos) || b.horizWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall already exists")
		} else if wallPos.r != 0 && wallPos.r != b.n_rows && b.vertiWalls.Get(wallPos.D().R()) && b.vertiWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall intersects")
		} else {
			b.horizWalls.Set(wallPos)
			b.horizWalls.Set(wallPos.R())
		}
	case VertiWall:
		if wallPos.r < 0 || wallPos.c < 1 || wallPos.r > b.n_rows-2 || wallPos.c > b.n_cols-1 { // do not allow columns on far edges
			return fmt.Errorf("wall out of bounds")
		} else if b.vertiWalls.Get(wallPos) || b.vertiWalls.Get(wallPos.U()) {
			return fmt.Errorf("wall already exists")
		} else if b.vertiWalls.Get(wallPos.D().R()) && b.vertiWalls.Get(wallPos.R()) {
			return fmt.Errorf("wall intersects")
		} else {
			b.vertiWalls.Set(wallPos)
			b.vertiWalls.Set(wallPos.U())
		}
	case Down:
		if b.horizWalls.Get(curPos) {
			return fmt.Errorf("hit wall")
		} else if curPos.r == 0 && b.curPlayer {
			b.win <- true
		} else if curPos.r == 0 && !b.curPlayer {
			return fmt.Errorf("hit floor")
		} else {
			if b.curPlayer {
				b.pos1 = curPos.D()
			} else {
				b.pos0 = curPos.D()
			}
		}
	case Up:
		if b.horizWalls.Get(curPos.U()) {
			return fmt.Errorf("hit wall")
		} else if curPos.r == b.n_rows-1 && b.curPlayer {
			return fmt.Errorf("hit ceiling")
		} else if curPos.r == b.n_rows-1 && !b.curPlayer {
			b.win <- false
		} else {
			if b.curPlayer {
				b.pos1 = curPos.U()
			} else {
				b.pos0 = curPos.U()
			}
		}
	case Left:
		if b.vertiWalls.Get(curPos) {
			return fmt.Errorf("hit wall")
		} else if curPos.r == 0 {
			return fmt.Errorf("hit left border")
		} else {
			if b.curPlayer {
				b.pos1 = curPos.L()
			} else {
				b.pos0 = curPos.L()
			}
		}
	case Right:
		if b.vertiWalls.Get(curPos.R()) {
			return fmt.Errorf("hit wall")
		} else if curPos.r == b.n_cols-1 {
			return fmt.Errorf("hit right border")
		} else {
			if b.curPlayer {
				b.pos1 = curPos.R()
			} else {
				b.pos0 = curPos.R()
			}
		}
	default:
		return fmt.Errorf("Not a valid move type")
	}

	// only switch player if move was successful
	b.curPlayer = !b.curPlayer
	return nil
}

func (b *Board) Copy() *Board {
	var newBoard *Board
	newBoard.Init(b.n_rows, b.n_cols, b.win)
	newBoard.curPlayer = b.curPlayer
	newBoard.pos1 = b.pos1.Copy()
	newBoard.pos0 = b.pos0.Copy()
	newBoard.vertiWalls = b.vertiWalls.Copy()
	newBoard.horizWalls = b.horizWalls.Copy()
	return newBoard
}

func (b *Board) Validate() bool {
	var visited0, visited1 *Matrix
	visited0.Init(b.n_rows, b.n_cols)
	visited1.Init(b.n_rows, b.n_cols)
	return b.walk(b.pos0, visited0, false) && b.walk(b.pos1, visited1, true)
}

func (b *Board) walk(pos *Pos, visited *Matrix, curWalker bool) bool {
	visited.Set(pos)

	var neighbors []*Pos
	if !b.vertiWalls.Get(pos) && pos.c != 0 {
		neighbors = append(neighbors, pos.L())
	}
	if !b.vertiWalls.Get(pos.R()) && pos.c != b.n_cols-1 {
		neighbors = append(neighbors, pos.R())
	}
	if !b.horizWalls.Get(pos) && pos.r == 0 && curWalker {
		return true
	} else if b.horizWalls.Get(pos) && pos.r != 0 {
		neighbors = append(neighbors, pos.D())
	}
	if !b.horizWalls.Get(pos.U()) && pos.r == b.n_rows-1 && !curWalker {
		return true
	} else if b.horizWalls.Get(pos) && pos.r != b.n_rows-1 {
		neighbors = append(neighbors, pos.U())
	}

	for _, neighborPos := range neighbors {
		if !visited.Get(neighborPos) && b.walk(neighborPos, visited, curWalker) {
			return true
		}
	}
	return false
}

type Game struct {
	board *Board

	p0         *Player
	p1         *Player
	visualizer *Visualizer

	win chan bool
}

func (g *Game) Init(n_rows int, n_cols int, p0 Player, p1 Player, v Visualizer) {
	g.p0 = &p0
	g.p1 = &p1

	win := make(chan bool)
	g.board.Init(n_rows, n_cols, win)
	g.visualizer = &v
}

func (g *Game) Play() bool {
	for {
		select {
		case winner := <-g.win:
			return winner
		default:
			g.Display()
			var moveType MoveType
			var wallPos *Pos
			if g.board.curPlayer {
				moveType, wallPos = (*g.p1).Move(g.board)
			} else {
				moveType, wallPos = (*g.p0).Move(g.board)
			}
			if err := g.board.Move(moveType, wallPos); err != nil {
				// if player makes invalid move other player wins
				return !g.board.curPlayer
			}
		}
	}
}

func (g *Game) Display() {
	v := *g.visualizer
	v.Display(g.board)
}

type Matrix struct {
	n_rows int
	n_cols int
	grid   []bool
}

func (m *Matrix) Init(n_rows int, n_cols int) {
	m.n_rows = n_rows
	m.n_cols = n_cols

	var gridSize = n_rows * n_cols
	m.grid = make([]bool, gridSize)
}

func (m *Matrix) Get(pos *Pos) bool {
	var idx = pos.r*m.n_cols + pos.c
	return m.grid[idx]
}

func (m *Matrix) Set(pos *Pos) {
	var idx = pos.r*m.n_cols + pos.c
	m.grid[idx] = true
}

func (m *Matrix) Copy() *Matrix {
	grid := make([]bool, m.n_rows*m.n_cols)
	copy(grid, m.grid)
	return &Matrix{
		n_rows: m.n_rows,
		n_cols: m.n_cols,
		grid:   grid,
	}
}

type Pos struct {
	r int
	c int
}

func (p *Pos) U() *Pos {
	return &Pos{
		r: p.r + 1,
		c: p.c,
	}
}

func (p *Pos) D() *Pos {
	return &Pos{
		r: p.r - 1,
		c: p.c,
	}
}

func (p *Pos) R() *Pos {
	return &Pos{
		r: p.r,
		c: p.c + 1,
	}
}

func (p *Pos) L() *Pos {
	return &Pos{
		r: p.r,
		c: p.c - 1,
	}
}

func (p *Pos) Copy() *Pos {
	return &Pos{
		r: p.r,
		c: p.c,
	}
}

// So far unused
// func (p *Pos) Validate(n_rows int, n_cols int) bool {
//  return p.r > 0 && p.c > 0 && p.r < n_rows && p.c < n_cols
// }
