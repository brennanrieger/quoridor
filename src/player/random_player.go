package player

import (
	"board"
	"math/rand"
	"util"
)

type RandomPlayer struct {
	playerNum bool
}

func (rp *RandomPlayer) Init(playerNum bool) {
	rp.playerNum = playerNum
}

func (rp *RandomPlayer) Move(b *board.Board) *board.Move {
	if rp.playerNum != b.CurPlayer {
		panic("it's not my turn")
	}

	var availableMoves = util.AvailableMoves(b)
	return availableMoves[rand.Intn(len(availableMoves))]
}
