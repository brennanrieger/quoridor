// WIP

package main

// type ManhattanDistance struct {
// 	board *Board
// }

// func (md *ManhattanDistance) Val(b *Board) {
// 	visited := [b.n_rows][b.n_cols]int{}
// }

// func (md *ManhattanDistance) walk(pos *Pos, visited *Matrix) bool {
// 	visited.Set(pos)

// 	var neighbors []*Pos
// 	if !b.vertiWalls.Get(pos) && pos.c != 0 {
// 		neighbors = append(neighbors, pos.L())
// 	}
// 	if !b.vertiWalls.Get(pos.R()) && pos.c != b.n_cols-1 {
// 		neighbors = append(neighbors, pos.R())
// 	}
// 	if !b.horizWalls.Get(pos) && pos.r == 0 && curWalker {
// 		return true
// 	} else if !b.horizWalls.Get(pos) && pos.r != 0 {
// 		neighbors = append(neighbors, pos.D())
// 	}
// 	if !b.horizWalls.Get(pos.U()) && pos.r == b.n_rows-1 && !curWalker {
// 		return true
// 	} else if !b.horizWalls.Get(pos.U()) && pos.r != b.n_rows-1 {
// 		neighbors = append(neighbors, pos.U())
// 	}

// 	for _, neighborPos := range neighbors {
// 		if !visited.Get(neighborPos) && b.walk(neighborPos, visited, curWalker) {
// 			return true
// 		}
// 	}
// 	return false
// }
