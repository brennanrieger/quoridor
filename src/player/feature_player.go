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
	var md = &feature.ManhattanDistance{}
	var availableMoves = util.AvailableMoves(b, fp.playerNum)

	bestVal := 0
	bestI := 0
	dummyWinCh := make(chan bool, 2)
	for i, move := range availableMoves {
		bNew := b.Copy()
		bNew.Move(move, fp.playerNum, dummyWinCh)
		val := md.Val(bNew)
		if val > bestVal {
			bestVal = val
			bestI = i
		}
	}
	return availableMoves[bestI]
}
