package board

import "fmt"

type Matrix struct {
	NRows int
	NCols int
	grid  []bool
}

func (m *Matrix) Init(NRows int, NCols int) {
	m.NRows = NRows
	m.NCols = NCols

	var gridSize = NRows * NCols
	m.grid = make([]bool, gridSize)
}

func (m *Matrix) Get(pos *Pos) bool {
	var idx = pos.r*m.NCols + pos.c
	return m.grid[idx]
}

func (m *Matrix) Set(pos *Pos) {
	var idx = pos.r*m.NCols + pos.c
	m.grid[idx] = true
}

func (m *Matrix) Copy() *Matrix {
	grid := make([]bool, m.NRows*m.NCols)
	copy(grid, m.grid)
	return &Matrix{
		NRows: m.NRows,
		NCols: m.NCols,
		grid:  grid,
	}
}

func (m *Matrix) Flip() *Matrix {
	grid := make([]bool, m.NRows*m.NCols)
	for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
		grid[i], grid[j] = grid[j], grid[i]
	}
	return &Matrix{
		NRows: m.NRows,
		NCols: m.NCols,
		grid:  grid,
	}
}

func (m *Matrix) Show() {
	var disp string
	for r := m.NRows - 1; r >= 0; r-- {
		for c := 0; c < m.NCols; c++ {
			if m.grid[r*m.NCols+c] {
				disp += "1"
			} else {
				disp += "0"
			}
		}
		disp += "\n"
	}
	disp += "\n"
	fmt.Println(disp)
}
