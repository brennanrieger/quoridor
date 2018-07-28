package main

import (
	"math/rand"
)

type RandomPlayer struct {
	playerNum bool
}

func (rp *RandomPlayer) Init(playerNum) {
	rp.playerNum = playerNum
}

func (rp *RandomPlayer) Move(b *Board) (MoveType, *Pos) {
	var boardCopy = b.Copy()
	var option = rand.Intn(3)
	var mt MoveType
	var po *Pos
	if option == 1 {
		mt, po = rp.makeWall(boardCopy, true)
	} else if option == 2 {
		mt, po = rp.makeWall(boardCopy, false)
	} else {
		mt, po = rp.movePiece(b), &Pos{}
	}
	return mt, po
}

func (rp *RandomPlayer) makeWall(b *Board, horizontal bool) (MoveType, *Pos) {
	var availablePositions []*Pos

	var moveType MoveType
	if horizontal {
		moveType = HorizWall
	} else {
		moveType = VertiWall
	}

	for r := 0; r < b.n_rows-1; r++ {
		for c := 0; c < b.n_cols-1; c++ {
			var boardCopy = b.Copy()
			pos := &Pos{
				r: r,
				c: c,
			}
			dummyWinCh := make(chan bool, 2)
			if err := boardCopy.Move(moveType, pos, dummyWinCh); err == nil && boardCopy.Validate() {
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

func (rp *RandomPlayer) movePiece(b *Board) MoveType {
	var availableMoves []MoveType

	var curPos *Pos
	if b.curPlayer {
		curPos = b.pos1
	} else {
		curPos = b.pos0
	}

	if !b.horizWalls.Get(curPos) && (curPos.r != 0 || b.curPlayer == true) {
		availableMoves = append(availableMoves, Down)
	}
	if !b.horizWalls.Get(curPos.U()) && (curPos.r != b.n_rows-1 || b.curPlayer == false) {
		availableMoves = append(availableMoves, Up)
	}
	if !b.vertiWalls.Get(curPos) && curPos.c != 0 {
		availableMoves = append(availableMoves, Left)
	}
	if !b.vertiWalls.Get(curPos.R()) && curPos.c != b.n_cols-1 {
		availableMoves = append(availableMoves, Right)
	}

	return availableMoves[rand.Intn(len(availableMoves))]
}
