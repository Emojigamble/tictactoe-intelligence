package main

import (
	"fmt"
	"github.com/Emojigamble/tictactoe-intelligence/ai"
	"github.com/Emojigamble/tictactoe-intelligence/game"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	agent := ai.Agent{}
	var wg sync.WaitGroup

	iterations := 500000
	goroutines := 50

	for goroutines > 0 {
		wg.Add(1)
		go func() {
			train(iterations, &agent)
			wg.Done()
		}()

		goroutines -= 1
	}

	wg.Wait()

	fmt.Println(len(agent.QTable))
}

func train(iterations int, agent *ai.Agent) {
	i := 1

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
			optimalMove := agent.OptimalMove(g, true, float32(i/iterations))
			history = append(history, ai.Move{Board: g.Board, Move: optimalMove})
			_ = g.Input(optimalMove, g.ActivePlayer)
		} else {
			_ = g.Input(g.LegalMoves()[rand.Intn(len(g.LegalMoves()))], g.ActivePlayer)
		}

		i++

		if i % 50000 == 0 {
			fmt.Println(i)
			rand.Seed(time.Now().Unix())
		}

	}
}