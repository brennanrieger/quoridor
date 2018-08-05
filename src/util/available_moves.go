package util

// import (
// 	"board"
// 	"math/rand"
// )

// func AvailableMoves(b *Board) []*board.Pos {
// 	return
// }

// func (rp *RandomPlayer) Move(b *board.Board) (board.MoveType, *board.Pos) {
// 	var boardCopy = b.Copy()
// 	var option = rand.Intn(3)
// 	var mt board.MoveType
// 	var po *board.Pos
// 	if option == 1 {
// 		mt, po = rp.makeWall(boardCopy, true)
// 	} else if option == 2 {
// 		mt, po = rp.makeWall(boardCopy, false)
// 	} else {
// 		mt, po = rp.movePiece(b), &board.Pos{}
// 	}
// 	return mt, po
// }

// func availableWalls(b *board.Board, horizontal bool) (board.MoveType, *board.Pos) {
// 	var availablePositions []*board.Pos

// 	var moveType board.MoveType
// 	if horizontal {
// 		moveType = board.HorizWall
// 	} else {
// 		moveType = board.VertiWall
// 	}

// 	for r := 0; r < b.NRows-1; r++ {
// 		for c := 0; c < b.NCols-1; c++ {
// 			var boardCopy = b.Copy()
// 			pos := &board.Pos{
// 				Row: r,
// 				Col: c,
// 			}
// 			dummyWinCh := make(chan bool, 2)
// 			if err := boardCopy.Move(moveType, pos, rp.playerNum, dummyWinCh); err == nil && boardCopy.Validate() {
// 				availablePositions = append(availablePositions, pos)
// 			}
// 		}
// 	}

// 	if len(availablePositions) > 0 {
// 		return moveType, availablePositions[rand.Intn(len(availablePositions))]
// 	} else {
// 		return rp.movePiece(b), nil
// 	}
// }

// func (rp *RandomPlayer) movePiece(b *board.Board) board.MoveType {
// 	var availableMoves []board.MoveType

// 	var curPos *board.Pos
// 	if rp.playerNum {
// 		curPos = b.Pos1
// 	} else {
// 		curPos = b.Pos0
// 	}

// 	if !b.HorizWalls.Get(curPos) && (curPos.Row != 0 || rp.playerNum == true) {
// 		availableMoves = append(availableMoves, board.Down)
// 	}
// 	if !b.HorizWalls.Get(curPos.U()) && (curPos.Row != b.NRows-1 || rp.playerNum == false) {
// 		availableMoves = append(availableMoves, board.Up)
// 	}
// 	if !b.VertiWalls.Get(curPos) && curPos.Col != 0 {
// 		availableMoves = append(availableMoves, board.Left)
// 	}
// 	if !b.VertiWalls.Get(curPos.R()) && curPos.Col != b.NCols-1 {
// 		availableMoves = append(availableMoves, board.Right)
// 	}

// 	return availableMoves[rand.Intn(len(availableMoves))]
// }
