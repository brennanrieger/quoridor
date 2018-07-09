package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HumanPlayer struct{}

func (hp *HumanPlayer) Move(b *Board) (MoveType, *Pos) {
	return hp.move()
}

func (hp *HumanPlayer) move() (MoveType, *Pos) {
	rawText := hp.promptUser()
	inputs := strings.Split(rawText, " ")
	var pos *Pos

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
		pos, err := hp.parseWallPos(inputs)
		if err != nil {
			hp.helpText()
			return hp.move()
		} else {
			return HorizWall, pos
		}
	case "v":
		pos, err := hp.parseWallPos(inputs)
		if err != nil {
			hp.helpText()
			return hp.move()
		} else {
			return VertiWall, pos
		}
	default:
		fmt.Printf("Invalid move. Please try again")
		hp.helpText()
		return hp.move()
	}
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

func (hp *HumanPlayer) parseWallPos(inputs []string) (*Pos, error) {
	var pos *Pos
	if len(inputs) != 3 {
		return pos, fmt.Errorf("Building a wall requires 3 inputs")
	}

	r, rErr := strconv.ParseInt(inputs[1], 10, 16)
	c, cErr := strconv.ParseInt(inputs[2], 10, 16)
	if rErr != nil || cErr != nil {
		return pos, fmt.Errorf("%v %v", rErr, cErr)
	} else {
		pos = &Pos{
			r: int(r),
			c: int(c),
		}
		return pos, nil
	}
}

func (hp *HumanPlayer) helpText() {
	fmt.Println(`
Usage:
There are 6 valid move types: left, up, down, right, horizontal wall and vertical wall
To move your piece, press the a w s d keys to move it left, up, down, and right, respectively
To build a horizontal wall, press h followed by the row and column of the left corner wall
To build a vertical wall, press v followed by the row and column of the bottom corner wall
Ex: h 3 5
Ex: a
Ex: v 2 3
Ex: s
	`)
}
