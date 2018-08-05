package util

import "board"

type ValMatrix struct {
	NRows int
	NCols int
	grid  []interface{}
}

func (vm *ValMatrix) Init(nRows int, nCols int) {
	vm.NRows = nRows
	vm.NCols = nCols

	var gridSize = nRows * nCols
	vm.grid = make([]interface{}, gridSize)
}

func (vm *ValMatrix) Get(pos *board.Pos) interface{} {
	var idx = pos.Row*vm.NCols + pos.Col
	return vm.grid[idx]
}

func (vm *ValMatrix) Set(pos *board.Pos, val interface{}) {
	var idx = pos.Row*vm.NCols + pos.Col
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
