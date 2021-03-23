package main

import (
	"fmt"
	"github.com/Emojigamble/tictactoe-intelligence/ai"
	"github.com/Emojigamble/tictactoe-intelligence/game"
	"math/rand"
)

func main() {
	g := game.TicTacToeGame{}
	g.Init()

	agent := ai.Agent{}

	i := 0
	for i < 1000000 {

		if w, tie := g.HasWinner(); w != game.Empty || tie != false {
			g.Init()
			if tie {
				agent.ResetHistory()
				continue
			}
			if w == game.One {
				agent.GiveReward()
				agent.ResetHistory()
			}
			agent.ResetHistory()
			continue
		}

		if g.ActivePlayer == game.One {
			_ = g.Input(agent.OptimalMove(g, true), g.ActivePlayer)
		} else {
			_ = g.Input(g.LegalMoves()[rand.Intn(len(g.LegalMoves()))], g.ActivePlayer)
		}

		i++
		fmt.Println(i)

	}

	fmt.Println(len(agent.QTable))
}
