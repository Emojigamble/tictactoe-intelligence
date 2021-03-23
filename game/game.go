package game

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type TicTacToeGame struct {
	Board        [9]Player
	ActivePlayer Player
}

func (g *TicTacToeGame) Init() {
	for i, _ := range g.Board {
		g.Board[i] = Empty
	}

	rand.Seed(time.Now().Unix())
	randInt := rand.Intn(2)
	g.ActivePlayer = Player(randInt + 1)
}

func (g *TicTacToeGame) LegalMoves() []int {
	var moves []int

	for i, _ := range g.Board {
		if g.Board[i] == Empty {
			moves = append(moves, i)
		}
	}

	return moves
}

func (g *TicTacToeGame) HasWinner() (winner Player, tie bool) {
	var winCombinations = [][]int{
		{0,1,2},
		{3,4,5},
		{6,7,8},
		{0,3,6},
		{1,4,7},
		{2,5,8},
		{0,4,8},
		{2,4,6},
	}

	for _, combination := range winCombinations {
		combinationType := Empty
		for _, entry := range combination {
			if g.Board[entry] == Empty {
				combinationType = Empty
				break
			} else if combinationType == Empty || combinationType == g.Board[entry] {
				combinationType = g.Board[entry]
			} else {
				combinationType = Empty
				break
			}
		}

		if combinationType != Empty {
			winner = combinationType
			break
		}
	}

	if winner != Empty {
		return winner, false
	} else if len(g.LegalMoves()) == 0 {
		return Empty, true
	} else {
		return Empty, false
	}
}

func (g *TicTacToeGame) Hash() string {
	delimiter := ","
	boardString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(g.Board)), delimiter), "[]")

	h := sha1.New()
	h.Write([]byte(boardString))
	return hex.EncodeToString(h.Sum(nil))
}
