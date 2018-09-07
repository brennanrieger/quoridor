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
	var availableMoves = util.AvailableMoves(b, rp.playerNum)
	return availableMoves[rand.Intn(len(availableMoves))]
}
