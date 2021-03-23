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

func (g *TicTacToeGame) Hash() string {
	delimiter := ","
	boardString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(g.Board)), delimiter), "[]")

	h := sha1.New()
	h.Write([]byte(boardString))
	return hex.EncodeToString(h.Sum(nil))
}
