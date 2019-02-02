package board

import "fmt"

type MoveType int

const (
	HorizWall = 1
	VertiWall = 2
	Up        = 3
	Down      = 4
	Right     = 5
	Left      = 6
	Jump      = 7
)

type Move struct {
	Mt  MoveType
	Pos *Pos
}

func (m *Move) Copy() *Move {
	return &Move{
		Mt:  m.Mt,
		Pos: m.Pos.Copy(),
	}
}

func (m *Move) Show() {
	fmt.Println("Move type: ", m.Mt, "\nPosition: ")
	m.Pos.Show()
}

func StepMove(moveType MoveType) *Move {
	if moveType < 3 || moveType > 6 {
		panic(fmt.Sprintf("moveType %d is not a step move", moveType))
	}
	return &Move{
		Mt:  moveType,
		Pos: &Pos{},
	}
}
