package ai

import (
	"encoding/gob"
	"fmt"
	"github.com/Emojigamble/tictactoe-intelligence/game"
	"math"
	"math/rand"
	"os"
	"sort"
)

type Agent struct {
	QTable  []StateSet
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

func (a *Agent) LoadQTable() {
	dataFile, err := os.Open("qtable.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&a.QTable)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Loaded QTable with %d entries\n", len(a.QTable))

	dataFile.Close()
}

func (a *Agent) SaveQTable() {
	dataFile, err := os.Create("qtable.gob")

	if err != nil {
		fmt.Println(err)
		return
	}

	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(a.QTable)

	dataFile.Close()
}

func (a *Agent) GiveReward(history []Move) {
	for i, m := range history {
		qEntry, qIndex := a.GetQEntry(game.Hash(m.Board))

		reward := 1
		rewardFactor := (i+1)/len(history)

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

func (a *Agent) OptimalMove(g game.TicTacToeGame, train bool, trainThreshold float32) int {
	h := game.Hash(g.Board)
	stateSet, _ := a.GetQEntry(h)

	sort.Slice(stateSet.Fields[:], func(i, j int) bool {
		return stateSet.Fields[i].Value < stateSet.Fields[j].Value
	})

	legalMoves := g.LegalMoves()
	move := legalMoves[rand.Intn(len(legalMoves))]

	if ((train == true && rand.Intn(100) > int(trainThreshold*50)) || train == false) && len(stateSet.Fields) > 0 {
		move = stateSet.Fields[0].Index
	}

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
