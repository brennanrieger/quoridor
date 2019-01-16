package player

import (
	"board"
	"feature"
	"math"
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

	bestVal := math.Inf(-1)
	bestI := 0
	win := make(chan bool, 2)
	for i, move := range availableMoves {
		bNew := b.Copy()
		bNew.MakeMove(move, fp.playerNum, win)
		select {
		case <-win:
			return move
		default:
			move.Show()
			var md = &feature.ManhattanDistance{}
			val0, val1 := md.Val(bNew)
			if fp.playerNum == 0 {
				if val1-val0 > bestVal {
					bestVal = val1 - val0
					bestI = i
				}
			} else {
				if val0-val1 > bestVal {
					bestVal = val1 - val0
					bestI = i
				}
			}
		}
	}
	return availableMoves[bestI]
}
