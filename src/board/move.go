package board

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
