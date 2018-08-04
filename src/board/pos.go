package board

import "fmt"

type Pos struct {
	Row int
	Col int
}

func (p *Pos) U() *Pos {
	return &Pos{
		Row: p.Row + 1,
		Col: p.Col,
	}
}

func (p *Pos) D() *Pos {
	return &Pos{
		Row: p.Row - 1,
		Col: p.Col,
	}
}

func (p *Pos) R() *Pos {
	return &Pos{
		Row: p.Row,
		Col: p.Col + 1,
	}
}

func (p *Pos) L() *Pos {
	return &Pos{
		Row: p.Row,
		Col: p.Col - 1,
	}
}

func (p *Pos) Copy() *Pos {
	return &Pos{
		Row: p.Row,
		Col: p.Col,
	}
}

func (p *Pos) Equal(pos2 *Pos) bool {
	return p.Row == pos2.Row && p.Col == pos2.Col
}

func (p *Pos) Show() {
	fmt.Println(p.Row, p.Col)
}

// So far unused
// func (p *Pos) Validate(NRows int, NCols int) bool {
// 	return p.Row > 0 && p.Col > 0 && p.Row < NRows && p.Col < NCols
// }
