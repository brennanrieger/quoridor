package player

import (
	"board"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HumanPlayer struct{}

func (hp *HumanPlayer) Move(b *board.Board) *board.Move {
	var move *board.Move
	validMove := false
	boardCopy := b.Copy()

	for !validMove {
		move = hp.move()
		if err := boardCopy.MakeMove(b.CurPlayer, move); err == nil {
			validMove = true
		} else {
			hp.helpText()
		}
	}

	return move
}

func (hp *HumanPlayer) move() *board.Move {
	rawText := hp.promptUser()
	inputs := strings.Split(rawText, " ")

	switch inputs[0] {
	case "a":
		return board.StepMove(board.Left)
	case "w":
		return board.StepMove(board.Up)
	case "s":
		return board.StepMove(board.Down)
	case "d":
		return board.StepMove(board.Right)
	case "h":
		pos, err := hp.parsePos(inputs)
		if err != nil {
			hp.helpText()
			return hp.move()
		} else {
			return &board.Move{
				Mt:  board.HorizWall,
				Pos: pos,
			}
		}
	case "v":
		pos, err := hp.parsePos(inputs)
		if err != nil {
			hp.helpText()
			return hp.move()
		} else {
			return &board.Move{
				Mt:  board.VertiWall,
				Pos: pos,
			}
		}
	case "j":
		pos, err := hp.parsePos(inputs)
		if err != nil {
			hp.helpText()
			return hp.move()
		} else {
			return &board.Move{
				Mt:  board.Jump,
				Pos: pos,
			}
		}
	default:
		fmt.Println(inputs[0])
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
		return strings.TrimSuffix(text, "\n")
	}
}

func (hp *HumanPlayer) parsePos(inputs []string) (*board.Pos, error) {
	var pos *board.Pos
	if len(inputs) != 3 {
		return pos, fmt.Errorf("3 arguments required")
	}

	r, rErr := strconv.ParseInt(inputs[1], 10, 16)
	c, cErr := strconv.ParseInt(inputs[2], 10, 16)
	if rErr != nil || cErr != nil {
		return pos, fmt.Errorf("%v %v", rErr, cErr)
	} else {
		pos = &board.Pos{
			Row: int(r),
			Col: int(c),
		}
		return pos, nil
	}
}

func (hp *HumanPlayer) helpText() {
	fmt.Println(`

==========================================================================================
=
=   USAGE:
=
=   There are 7 valid move types: left, up, down, right, horizontal wall, vertical wall, and jump
=   To move your piece, press the a w s d keys to move it left, up, down, and right, respectively
=   To build a horizontal wall, press h followed by the row and column of the left corner wall
=   To build a vertical wall, press v followed by the row and column of the bottom corner wall
=   To jump over your opponent, press j followed by the row and column of the square to jump to
=
=   Ex: h 3 5
=   Ex: a
=   Ex: v 2 3
=   Ex: s
=   Ex: j 4 4
=
=========================================================================================
	`)
}
