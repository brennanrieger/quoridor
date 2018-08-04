// WIP

package main

// type ManhattanDistance struct {
// 	board *board.Board
// }

// func (md *ManhattanDistance) Val(b *board.Board) {
// 	visited := [b.NRows][b.NCols]int{}
// }

// func (md *ManhattanDistance) walk(pos *board.Pos, visited *Matrix) bool {
// 	visited.Set(pos)

// 	var neighbors []*board.Pos
// 	if !b.VertiWalls.Get(pos) && pos.c != 0 {
// 		neighbors = append(neighbors, pos.L())
// 	}
// 	if !b.VertiWalls.Get(pos.R()) && pos.c != b.NCols-1 {
// 		neighbors = append(neighbors, pos.R())
// 	}
// 	if !b.HorizWalls.Get(pos) && pos.r == 0 && curWalker {
// 		return true
// 	} else if !b.HorizWalls.Get(pos) && pos.r != 0 {
// 		neighbors = append(neighbors, pos.D())
// 	}
// 	if !b.HorizWalls.Get(pos.U()) && pos.r == b.NRows-1 && !curWalker {
// 		return true
// 	} else if !b.HorizWalls.Get(pos.U()) && pos.r != b.NRows-1 {
// 		neighbors = append(neighbors, pos.U())
// 	}

// 	for _, neighborPos := range neighbors {
// 		if !visited.Get(neighborPos) && b.walk(neighborPos, visited, curWalker) {
// 			return true
// 		}
// 	}
// 	return false
// }
