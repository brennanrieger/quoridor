package main

import "board"

type ValMatrix struct {
	NRows int
	NCols int
	grid  []interface{}
}

func (vm *ValMatrix) Init(NRows int, NCols int) {
	vm.NRows = NRows
	vm.NCols = NCols

	var gridSize = NRows * NCols
	vm.grid = make([]interface{}, gridSize)
}

func (vm *ValMatrix) Get(pos *board.Pos) interface{} {
	var idx = pos.r*vm.NCols + pos.c
	return vm.grid[idx]
}

func (vm *ValMatrix) Set(pos *board.Pos, val interface{}) {
	var idx = pos.r*vm.NCols + pos.c
	vm.grid[idx] = val
}

func (vm *ValMatrix) Copy() *ValMatrix {
	grid := make([]interface{}, vm.NRows*vm.NCols)
	copy(grid, vm.grid)
	return &ValMatrix{
		NRows: vm.NRows,
		NCols: vm.NCols,
		grid:  grid,
	}
}

func (vm *ValMatrix) Flip() *ValMatrix {
	grid := make([]interface{}, vm.NRows*vm.NCols)
	for i, j := 0, len(grid)-1; i < j; i, j = i+1, j-1 {
		grid[i], grid[j] = grid[j], grid[i]
	}
	return &ValMatrix{
		NRows: vm.NRows,
		NCols: vm.NCols,
		grid:  grid,
	}
}

// func (m *Matrix) Show() {
// 	var disp string
// 	for r := m.NRows - 1; r >= 0; r-- {
// 		for c := 0; c < m.NCols; c++ {
// 			if m.grid[r*m.NCols+c] {
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
