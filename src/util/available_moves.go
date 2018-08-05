package util

import (
	"board"
	"math/rand"
)

func AvailableMoves(b *Board, playerNum bool) []*board.Move {
  var horizWallMoves = availableWalls(b, true)
  var vertiWallMoves = availableWalls(b, false)
  var allWallMoves = append(horizWallMoves, vertiWallMoves...)
  return appent(allWallMoves, availableStepMoves(b, playerNum)...)
}

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

func availableWalls(b *board.Board, horizontal bool) ([]*board.Move) {
	var availableWalls []*board.Move

	var moveType board.MoveType
	if horizontal {
		moveType = board.HorizWall
	} else {
		moveType = board.VertiWall
	}

	for r := 0; r < b.NRows-1; r++ {
		for c := 0; c < b.NCols-1; c++ {
			var boardCopy = b.Copy()
			pos := &board.Pos{
				Row: r,
				Col: c,
			}
			dummyWinCh := make(chan bool, 2)
      var move = &board.Move{
        Mt: moveType,
        Pos: pos
      }
			if err := boardCopy.Move(move, rp.playerNum, dummyWinCh); err == nil && boardCopy.Validate() {
				availableWalls = append(availableWalls, move)
			}
		}
	}

  return availableWalls
}

func availableStepMoves(b *board.Board, playerNum bool) []*board.Move {
	var availableMoves []board.MoveType

	var curPos *board.Pos
	if rp.playerNum {
		curPos = b.Pos1
	} else {
		curPos = b.Pos0
	}

  var move = &board.Move{
    Mt: moveType,
    Pos: curPos
  }

	if !b.HorizWalls.Get(curPos) && (curPos.Row != 0 || rp.playerNum == true) {
    move.Mt = board.Down
    availableMoves = append(availableMoves, move.Copy())
	}
	if !b.HorizWalls.Get(curPos.U()) && (curPos.Row != b.NRows-1 || rp.playerNum == false) {
    move.Mt = board.Up
		availableMoves = append(availableMoves, move.Copy())
	}
	if !b.VertiWalls.Get(curPos) && curPos.Col != 0 {
	  move.Mt = board.Left
    availableMoves = append(availableMoves, move.Copy())
	}
	if !b.VertiWalls.Get(curPos.R()) && curPos.Col != b.NCols-1 {
	  move.Mt = board.Right
    availableMoves = append(availableMoves, move.Copy())
	}

	return availableMoves
}
