package game

import "testing"

func TestTicTacToeGame_Init(t *testing.T) {
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

func TestTicTacToeGame_LegalMoves(t *testing.T) {
	game := TicTacToeGame{}
	game.Init()

	moves := game.LegalMoves()
	if len(moves) < 9 {
		t.Log("Not all legal moves were detected")
		t.FailNow()
	}
}

func TestTicTacToeGame_Hash(t *testing.T) {
	game := TicTacToeGame{}
	game.Init()

	if Hash(game.Board) != "43bd6f03270a43a5488e069a0539c37a30ab402b" {
		t.Log("Hash for empty board did not match up")
		t.FailNow()
	}
}

func TestTicTacToeGame_HasWinner(t *testing.T) {
	game := TicTacToeGame{}
	game.Init()

	// Test winner
	game.Board[0] = One
	game.Board[4] = One
	game.Board[8] = One

	if w, _ := game.HasWinner(); w != One {
		t.Log("Did not detect winner properly")
		t.Fail()
	}

	// Test undecided board
	game.Board[0] = Empty
	if w, tie := game.HasWinner(); w != Empty || tie == true {
		t.Log("Did not detect a board without winner or tie properly")
		t.Fail()
	}

	// Test board with tie
	game.Board[0] = Two
	game.Board[1] = One
	game.Board[2] = One
	game.Board[3] = One
	game.Board[4] = Two
	game.Board[5] = Two
	game.Board[6] = One
	game.Board[7] = Two
	game.Board[8] = One
	if w, tie := game.HasWinner(); w != Empty || tie == false {
		t.Log("Did not detect a tie properly")
		t.Fail()
	}
}
