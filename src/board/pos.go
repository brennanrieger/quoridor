package board

import "fmt"

type Pos struct {
	r int
	c int
}

func (p *Pos) U() *Pos {
	return &Pos{
		r: p.r + 1,
		c: p.c,
	}
}

func (p *Pos) D() *Pos {
	return &Pos{
		r: p.r - 1,
		c: p.c,
	}
}

func (p *Pos) R() *Pos {
	return &Pos{
		r: p.r,
		c: p.c + 1,
	}
}

func (p *Pos) L() *Pos {
	return &Pos{
		r: p.r,
		c: p.c - 1,
	}
}

func (p *Pos) Copy() *Pos {
	return &Pos{
		r: p.r,
		c: p.c,
	}
}

func (p *Pos) Equal(pos2 *Pos) bool {
	return p.r == pos2.r && p.c == pos2.c
}

func (p *Pos) Show() {
	fmt.Println(p.r, p.c)
}

// So far unused
// func (p *Pos) Validate(NRows int, NCols int) bool {
// 	return p.r > 0 && p.c > 0 && p.r < NRows && p.c < NCols
// }
