package player

import (
	"board"
	"feature"
	"fmt"
	"math"
	"util"
)

type FeaturePlayer struct {
	playerNum   bool
	featureType feature.Feature
}

func (fp *FeaturePlayer) Init(playerNum bool, featureType feature.Feature) {
	fp.playerNum = playerNum
	fp.featureType = featureType
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
			val0, val1 := fp.featureType.Val(bNew)
			fmt.Println(val0)
			fmt.Println(val1)
			if fp.playerNum {
				// if player1, want player0's distance to be greater
				if val0-val1 > bestVal {
					bestVal = val0 - val1
					bestI = i
				}
			} else {
				if val1-val0 > bestVal {
					bestVal = val1 - val0
					bestI = i
				}
			}
		}
	}
	return availableMoves[bestI]
}
