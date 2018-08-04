package main

import (
	"board"
	"math/rand"
)

type RandomPlayer struct {
	playerNum bool
}

func (rp *RandomPlayer) Init(playerNum bool) {
	rp.playerNum = playerNum
}

func (rp *RandomPlayer) Move(b *board.Board) (MoveType, *board.Pos) {
	var boardCopy = b.Copy()
	var option = rand.Intn(3)
	var mt MoveType
	var po *board.Pos
	if option == 1 {
		mt, po = rp.makeWall(boardCopy, true)
	} else if option == 2 {
		mt, po = rp.makeWall(boardCopy, false)
	} else {
		mt, po = rp.movePiece(b), &board.Pos{}
	}
	return mt, po
}

func (rp *RandomPlayer) makeWall(b *board.Board, horizontal bool) (MoveType, *board.Pos) {
	var availablePositions []*board.Pos

	var moveType MoveType
	if horizontal {
		moveType = HorizWall
	} else {
		moveType = VertiWall
	}

	for r := 0; r < b.NRows-1; r++ {
		for c := 0; c < b.NCols-1; c++ {
			var boardCopy = b.Copy()
			pos := &board.Pos{
				r: r,
				c: c,
			}
			dummyWinCh := make(chan bool, 2)
			if err := boardCopy.Move(moveType, pos, rp.playerNum, dummyWinCh); err == nil && boardCopy.Validate() {
				availablePositions = append(availablePositions, pos)
			}
		}
	}

	if len(availablePositions) > 0 {
		return moveType, availablePositions[rand.Intn(len(availablePositions))]
	} else {
		return rp.movePiece(b), nil
	}
}

func (rp *RandomPlayer) movePiece(b *board.Board) MoveType {
	var availableMoves []MoveType

	var curPos *board.Pos
	if rp.playerNum {
		curPos = b.Pos1
	} else {
		curPos = b.Pos0
	}

	if !b.HorizWalls.Get(curPos) && (curPos.r != 0 || rp.playerNum == true) {
		availableMoves = append(availableMoves, Down)
	}
	if !b.HorizWalls.Get(curPos.U()) && (curPos.r != b.NRows-1 || rp.playerNum == false) {
		availableMoves = append(availableMoves, Up)
	}
	if !b.VertiWalls.Get(curPos) && curPos.c != 0 {
		availableMoves = append(availableMoves, Left)
	}
	if !b.VertiWalls.Get(curPos.R()) && curPos.c != b.NCols-1 {
		availableMoves = append(availableMoves, Right)
	}

	return availableMoves[rand.Intn(len(availableMoves))]
}
