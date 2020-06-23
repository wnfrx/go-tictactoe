package tictactoe

import "fmt"

type Board [3][3]string

func NewBoard() *Board {
	var b Board

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = "_"
		}
	}

	return &b
}

type IBoard interface {
	Draw()
	HasWinner() bool
	HasEmptyCell() bool
}

func (b *Board) Draw() {
	var str string

	str += "-------------\n"
	str += "    Board\n"
	str += "-------------\n"

	for _, row := range b {
		str += "   "
		for _, v := range row {
			str += v + "  "
		}
		str += "\n"
	}

	fmt.Println(str)
}

func (b *Board) HasWinner(mark1, mark2 string) bool {
	// Check rows
	if (b[0][0] == mark1 && b[0][1] == mark1 && b[0][2] == mark1) ||
		(b[1][0] == mark1 && b[1][1] == mark1 && b[1][2] == mark1) ||
		(b[2][0] == mark1 && b[2][1] == mark1 && b[2][2] == mark1) ||
		(b[0][0] == mark2 && b[0][1] == mark2 && b[0][2] == mark2) ||
		(b[1][0] == mark2 && b[1][1] == mark2 && b[1][2] == mark2) ||
		(b[2][0] == mark2 && b[2][1] == mark2 && b[2][2] == mark2) {
		return true
	}

	// Check columns
	if (b[0][0] == mark1 && b[1][0] == mark1 && b[2][0] == mark1) ||
		(b[0][1] == mark1 && b[1][1] == mark1 && b[2][1] == mark1) ||
		(b[0][2] == mark1 && b[1][2] == mark1 && b[2][2] == mark1) ||
		(b[0][0] == mark2 && b[1][0] == mark2 && b[2][0] == mark2) ||
		(b[0][1] == mark2 && b[1][1] == mark2 && b[2][1] == mark2) ||
		(b[0][2] == mark2 && b[1][2] == mark2 && b[2][2] == mark2) {
		return true
	}

	// Check diagonals
	if (b[0][0] == mark1 && b[1][1] == mark1 && b[2][2] == mark1) ||
		(b[0][2] == mark1 && b[1][1] == mark1 && b[2][0] == mark1) ||
		(b[0][0] == mark2 && b[1][1] == mark2 && b[2][2] == mark2) ||
		(b[0][2] == mark2 && b[1][1] == mark2 && b[2][0] == mark2) {
		return true
	}

	return false
}

func (b *Board) HasEmptyCell() bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == "_" {
				return true
			}
		}
	}

	return false
}
