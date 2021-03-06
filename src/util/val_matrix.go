package util

import (
	"board"
	"fmt"
	"reflect"
)

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
	for i := 0; i < len(grid); i++ {
		grid[i] = vm.grid[len(grid)-1-i]
	}
	return &ValMatrix{
		NRows: vm.NRows,
		NCols: vm.NCols,
		grid:  grid,
	}
}

func (vm *ValMatrix) Equal(vm2 *ValMatrix) bool {
	return vm.NRows == vm2.NRows && vm.NCols == vm2.NCols && reflect.DeepEqual(vm.grid, vm2.grid)
}

func (vm *ValMatrix) Show() {
	var disp string
	for r := vm.NRows - 1; r >= 0; r-- {
		for c := 0; c < vm.NCols; c++ {
			if vm.grid[r*vm.NCols+c] != nil {
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
