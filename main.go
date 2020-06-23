package main

import (
	"github.com/wnfrx/go-tictactoe/tictactoe"
)

func main() {
	// initialize player and game
	p1 := tictactoe.NewPlayer("Player 1", "X")
	p2 := tictactoe.NewPlayer("Player 2", "O")
	g := tictactoe.NewGame(*p1, *p2)

	// game start
	g.Start()
}
