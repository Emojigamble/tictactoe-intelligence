package game

import "testing"

func TestGameInit(t *testing.T) {
	game := TicTacToeGame{}
	game.Init()

	for _, f := range game.Board {
		if f != Empty {
			t.Log("Board not empty:", game.Board)
			t.FailNow()
		}
	}

	if game.ActivePlayer != One && game.ActivePlayer != Two {
		t.Log("No player is active")
		t.Fail()
	}
}

func TestLegalMoves(t *testing.T) {
	game := TicTacToeGame{}
	game.Init()

	moves := game.LegalMoves()
	if len(moves) < 9 {
		t.Log("Not all legal moves were detected")
		t.FailNow()
	}
}

func TestBoardHash(t *testing.T) {
	game := TicTacToeGame{}
	game.Init()

	if game.Hash() != "43bd6f03270a43a5488e069a0539c37a30ab402b" {
		t.Log("Hash for empty board did not match up")
		t.FailNow()
	}
}
