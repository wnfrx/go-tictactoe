package tictactoe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	MARK1 = "X"
	MARK2 = "O"
)

func TestEmptyBoard(t *testing.T) {
	var board Board

	/**
	Board State:
		_ _ _
		_ _ _
		_ _ _
	*/
	board[0][0] = "_"
	board[0][1] = "_"
	board[0][2] = "_"
	board[1][0] = "_"
	board[1][1] = "_"
	board[1][2] = "_"
	board[2][0] = "_"
	board[2][1] = "_"
	board[2][2] = "X"

	hasEmptyCell := board.HasEmptyCell()

	assert.Equal(t, true, hasEmptyCell)
}

func TestFullBoard(t *testing.T) {
	var board Board

	/**
	Board State:
		X O X
		O O X
		X X O
	*/
	board[0][0] = MARK1
	board[0][1] = MARK2
	board[0][2] = MARK1
	board[1][0] = MARK2
	board[1][1] = MARK2
	board[1][2] = MARK1
	board[2][0] = MARK1
	board[2][1] = MARK1
	board[2][2] = MARK2

	hasEmptyCell := board.HasEmptyCell()

	assert.Equal(t, false, hasEmptyCell)
}

func TestHasAWinner(t *testing.T) {
	var board Board

	/**
	Board State:
		X O X
		_ O X
		_ O _
	*/
	board[0][0] = MARK1
	board[0][1] = MARK2
	board[0][2] = MARK1
	board[1][0] = "_"
	board[1][1] = MARK2
	board[1][2] = MARK1
	board[2][0] = "_"
	board[2][1] = MARK2
	board[2][2] = "_"

	winner := board.GetWinner(MARK1, MARK2)

	assert.Equal(t, MARK2, winner)
}

func TestHasNoWinner(t *testing.T) {
	var board Board

	/**
	Board State:
		O O X
		X X O
		O X X
	*/
	board[0][0] = MARK2
	board[0][1] = MARK2
	board[0][2] = MARK1
	board[1][0] = MARK1
	board[1][1] = MARK1
	board[1][2] = MARK2
	board[2][0] = MARK2
	board[2][1] = MARK1
	board[2][2] = MARK1

	winner := board.GetWinner(MARK1, MARK2)

	assert.Equal(t, "", winner)
}
