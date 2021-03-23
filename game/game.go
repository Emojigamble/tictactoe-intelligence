package game

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
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

func (g *TicTacToeGame) Input(field int, player Player) error {
	if player != g.ActivePlayer {
		return errors.New(fmt.Sprint(player, "is currently not active"))
	}

	if len(g.LegalMoves()) == 0 {
		return errors.New("no more moves available")
	}

	if !g.IsLegalMove(field) {
		return errors.New("move not legal")
	}

	g.Board[field] = player
	if g.ActivePlayer == One {
		g.ActivePlayer = Two
	} else {
		g.ActivePlayer = One
	}

	return nil
}

func (g *TicTacToeGame) IsLegalMove(field int) bool {
	moves := g.LegalMoves()

	legal := false
	for _, m := range moves {
		if m == field {
			legal = true
		}
	}

	return legal
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
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
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

func Hash(board [9]Player) string {
	delimiter := ","
	boardString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(board)), delimiter), "[]")

	h := sha1.New()
	h.Write([]byte(boardString))
	return hex.EncodeToString(h.Sum(nil))
}

func (g *TicTacToeGame) PrintBoard() {
	for i, f := range g.Board {
		switch f {
			case One:
				print("X")
			case Two:
				print("O")
			default:
				print("-")
		}
		if (i+1)%3 == 0 {
			print("\n")
		} else {
			print(" | ")
		}
	}
}
