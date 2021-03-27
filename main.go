package main

import (
	"fmt"
	"github.com/Emojigamble/tictactoe-intelligence/ai"
	"github.com/Emojigamble/tictactoe-intelligence/game"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {

	agent := ai.Agent{}
	if _, err := os.Stat("qtable.gob"); err == nil {
		agent.LoadQTable()
	}

	var wg sync.WaitGroup
	iterations := 500000
	goroutines := 1

	for goroutines > 0 {
		wg.Add(1)

		reportProgress := false
		if goroutines == 1 {
			reportProgress = true
		}

		go func() {
			train(iterations, &agent, reportProgress)
			wg.Done()
		}()

		goroutines--
	}

	wg.Wait()

	fmt.Println()
	agent.SaveQTable()
	fmt.Printf("Saved QTable with %d entries", len(agent.QTable))
}

func train(iterations int, agent *ai.Agent, reportProgress bool) {
	i := 1
	step := iterations/300

	g := game.TicTacToeGame{}
	g.Init()

	var history []ai.Move

	for i < iterations {

		if w, tie := g.HasWinner(); w != game.Empty || tie != false {
			if w == game.One {
				agent.GiveReward(history)
			}
			history = []ai.Move{}
			g.Init()
			continue
		}

		if g.ActivePlayer == game.One {
			optimalMove := agent.OptimalMove(g, true, float32(i)/float32(iterations))
			history = append(history, ai.Move{Board: g.Board, Move: optimalMove})
			_ = g.Input(optimalMove, g.ActivePlayer)
		} else {
			_ = g.Input(g.LegalMoves()[rand.Intn(len(g.LegalMoves()))], g.ActivePlayer)
		}

		i++

		if i%step == 0 && reportProgress {
			fmt.Println(fmt.Sprintf("%.2f", (float32(i)/float32(iterations))*100), "%")
			rand.Seed(time.Now().Unix())
		}

	}
}
