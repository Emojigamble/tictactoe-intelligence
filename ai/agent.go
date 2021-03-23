package ai

import (
	"github.com/Emojigamble/tictactoe-intelligence/game"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Agent struct {
	QTable  []StateSet
	History []Move
}

type Move struct {
	Board [9]game.Player
	Move  int
}

type StateSet struct {
	Hash   string
	Fields []Field
}

type Field struct {
	Index int
	Value float64
}

func ModifiedSigmoid(x float64) float64 {
	return 1/(1+math.Exp(-x+4))
}

func (a *Agent) GiveReward() {
	for i, m := range a.History {
		qEntry, qIndex := a.GetQEntry(game.Hash(m.Board))

		reward := 1
		rewardFactor := (i+1)/len(a.History)

		field := GetField(m.Move, qEntry)
		if field == nil {
			field = &Field{
				Index: m.Move,
			}
			qEntry.Fields = append(qEntry.Fields, *field)
		}

		field.Value = ModifiedSigmoid(field.Value + (1-field.Value) * float64(rewardFactor*reward))
		if qIndex == -1 {
			a.QTable = append(a.QTable, *qEntry)
		}
	}
}

// ẞ != ß   þ þ √Þ≈ç√∫
func (a *Agent) ResetHistory() {
	a.History = []Move{}
}

func (a *Agent) OptimalMove(g game.TicTacToeGame, train bool) int {
	h := game.Hash(g.Board)
	stateSet, _ := a.GetQEntry(h)

	sort.Slice(stateSet.Fields[:], func(i, j int) bool {
		return stateSet.Fields[i].Value < stateSet.Fields[j].Value
	})

	rand.Seed(time.Now().Unix())

	legalMoves := g.LegalMoves()
	move := legalMoves[rand.Intn(len(legalMoves))]

	if ((train == true && rand.Intn(2) == 1) || train == false) && len(stateSet.Fields) > 0 {
		move = stateSet.Fields[0].Index
	}

	a.History = append(a.History, Move{g.Board, move})
	return move
}

func GetField(index int, set *StateSet) *Field {
	for i, _ := range set.Fields {
		if set.Fields[i].Index == index {
			return &set.Fields[i]
		}
	}
	return nil
}

func (a *Agent) GetQEntry(hash string) (*StateSet, int) {
	for i, q := range a.QTable {
		if q.Hash == hash {
			return &q, i
		}
	}
	return &StateSet{
		Hash: hash,
	}, -1
}
