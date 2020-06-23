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
	mark1Streak := [3]string{mark1, mark1, mark1}
	mark2Streak := [3]string{mark2, mark2, mark2}

	// Check rows
	rows := b.rows()
	for _, row := range rows {
		if row == mark1Streak || row == mark2Streak {
			return true
		}
	}

	// Check columns
	columns := b.columns()
	for _, col := range columns {
		if col == mark1Streak || col == mark2Streak {
			return true
		}
	}

	// Check diagonals
	diagonals := b.diagonals()
	for _, diagonal := range diagonals {
		if diagonal == mark1Streak || diagonal == mark2Streak {
			return true
		}
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

func (b *Board) rows() [3][3]string {
	return *b
}

func (b *Board) columns() [3][3]string {
	var columns [3][3]string

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			columns[j][i] = b[i][j]
		}
	}

	return columns
}

func (b *Board) diagonals() [2][3]string {
	var diagonals [2][3]string

	diagonals[0] = [3]string{b[0][0], b[1][1], b[2][2]}
	diagonals[1] = [3]string{b[2][0], b[1][1], b[0][2]}

	return diagonals
}
