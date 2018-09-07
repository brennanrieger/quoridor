package player

import (
	"board"
	"feature"
	"util"
)

type FeaturePlayer struct {
	playerNum bool
}

func (fp *FeaturePlayer) Init(playerNum bool) {
	fp.playerNum = playerNum
}

func (fp *FeaturePlayer) Move(b *board.Board) *board.Move {
	var availableMoves = util.AvailableMoves(b, fp.playerNum)

	bestVal := 0.
	bestI := 0
	winCh := make(chan bool, 2)
	for i, move := range availableMoves {
		bNew := b.Copy()
		bNew.MakeMove(move, fp.playerNum, winCh)
		select {
		case <-winCh:
			return move
		default:
			move.Show()
			var md = &feature.ManhattanDistance{}
			val0, val1 := md.Val(bNew)
			if val1-val0 > bestVal { // TODO don't hardcode player 0
				bestVal = val1 - val0
				bestI = i
			}
		}
	}
	return availableMoves[bestI]
}
