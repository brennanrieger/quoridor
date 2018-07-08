package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type HumanPlayer struct{}

func (hp *HumanPlayer) Move(b *Board) (MoveType, *Pos) {
	rawText := hp.promptUser
	inputs := Strings.split(rawText, ' ')

	switch inputs[0] {
	case "a":
		return Left, pos
	case "w":
		return Up, pos
	case "s":
		return Down, pos
	case "d":
		return Right, pos
	case "h":
		pos, err := parseWallPos(inputs)

	case "v":
		pos, err := parseWallPos(inputs)
	default:
		fmt.Printf("Invalid move. Please try again")
		hp.helpText()
		hp.promptUser()
	}
}

func (hp *HumanPlayer) move() MoveType, *Pos {

}

func (hp *HumanPlayer) promptUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter move: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input: ", err)
		return hp.promptUser()
	} else {
		return text
	}
}

func (hp *HumanPlayer) parseWallPos(inputs) *Pos, err {
	var pos *Pos
	if len(inputs) != 3 {
		return pos, fmt.Errorf("Building a wall requires 3 inputs")
	}

	r, rErr := strconv.ParseInt(inputs[1], 10, 32)
	c, cErr := strconv.ParseInt(inputs[2], 10, 32)
	if rErr != nil || cErr != nil {
		return pos, fmt.Errorf("%v %v", rErr, cErr)
	} else {
		return &Pos(r,c), nil
	}
}

func (hp *HumanPlayer) helpText() {
	fmt.Println("Usage:")
}

// func (rp *RandomPlayer) makeWall(b *Board, horizontal bool) (MoveType, *Pos) {
// 	var availablePositions []*Pos

// 	var moveType MoveType
// 	if horizontal {
// 		moveType = HorizWall
// 	} else {
// 		moveType = VertiWall
// 	}

// 	for r := 0; r < b.n_rows-1; r++ {
// 		for c := 0; c < b.n_cols-1; c++ {
// 			var boardCopy = b.Copy()
// 			pos := &Pos{
// 				r: r,
// 				c: c,
// 			}
// 			if err := boardCopy.Move(moveType, pos); err == nil && boardCopy.Validate() {
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

// func (rp *RandomPlayer) movePiece(b *Board) MoveType {
// 	var availableMoves []MoveType

// 	var curPos *Pos
// 	if b.curPlayer {
// 		curPos = b.pos1
// 	} else {
// 		curPos = b.pos0
// 	}

// 	if !b.horizWalls.Get(curPos) && (curPos.r != 0 || b.curPlayer == true) {
// 		availableMoves = append(availableMoves, Down)
// 	}
// 	if !b.horizWalls.Get(curPos.U()) && (curPos.r != b.n_rows-1 || b.curPlayer == false) {
// 		availableMoves = append(availableMoves, Up)
// 	}
// 	if !b.vertiWalls.Get(curPos) && curPos.c != 0 {
// 		availableMoves = append(availableMoves, Left)
// 	}
// 	if !b.vertiWalls.Get(curPos.R()) && curPos.c != b.n_cols-1 {
// 		availableMoves = append(availableMoves, Right)
// 	}

// 	return availableMoves[rand.Intn(len(availableMoves))]
// }
