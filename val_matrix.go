package main

type ValMatrix struct {
	n_rows int
	n_cols int
	grid   []interface{}
}

func (m *Matrix) Init(n_rows int, n_cols int) {
	m.n_rows = n_rows
	m.n_cols = n_cols

	var gridSize = n_rows * n_cols
	m.grid = make([]int, gridSize)
}

func (m *Matrix) Get(pos *Pos) interface{} {
	var idx = pos.r*m.n_cols + pos.c
	return m.grid[idx]
}

func (m *Matrix) Set(pos *Pos, val interface{}) {
	var idx = pos.r*m.n_cols + pos.c
	m.grid[idx] = val
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

func (m *Matrix) Flip() *Matrix {
	grid := make([]bool, m.n_rows*m.n_cols)
	for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
		grid[i], grid[j] = grid[j], grid[i]
	}
	return &Matrix{
		n_rows: m.n_rows,
		n_cols: m.n_cols,
		grid:   grid,
	}
}

// func (m *Matrix) Show() {
// 	var disp string
// 	for r := m.n_rows - 1; r >= 0; r-- {
// 		for c := 0; c < m.n_cols; c++ {
// 			if m.grid[r*m.n_cols+c] {
// 				disp += "1"
// 			} else {
// 				disp += "0"
// 			}
// 		}
// 		disp += "\n"
// 	}
// 	disp += "\n"
// 	fmt.Println(disp)
// }
