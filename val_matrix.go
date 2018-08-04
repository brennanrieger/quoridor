package main

type ValMatrix struct {
	n_rows int
	n_cols int
	grid   []interface{}
}

func (vm *ValMatrix) Init(n_rows int, n_cols int) {
	vm.n_rows = n_rows
	vm.n_cols = n_cols

	var gridSize = n_rows * n_cols
	vm.grid = make([]interface{}, gridSize)
}

func (vm *ValMatrix) Get(pos *Pos) interface{} {
	var idx = pos.r*vm.n_cols + pos.c
	return vm.grid[idx]
}

func (vm *ValMatrix) Set(pos *Pos, val interface{}) {
	var idx = pos.r*vm.n_cols + pos.c
	vm.grid[idx] = val
}

func (vm *ValMatrix) Copy() *ValMatrix {
	grid := make([]interface{}, vm.n_rows*vm.n_cols)
	copy(grid, vm.grid)
	return &ValMatrix{
		n_rows: vm.n_rows,
		n_cols: vm.n_cols,
		grid:   grid,
	}
}

func (vm *ValMatrix) Flip() *ValMatrix {
	grid := make([]interface{}, vm.n_rows*vm.n_cols)
	for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
		grid[i], grid[j] = grid[j], grid[i]
	}
	return &ValMatrix{
		n_rows: vm.n_rows,
		n_cols: vm.n_cols,
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
