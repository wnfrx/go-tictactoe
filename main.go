package main

import (
	"log"
	"os"

	"github.com/wnfrx/go-tictactoe/tictactoe"
)

func main() {
	// initialize logger
	l := log.New(os.Stdout, "", log.LstdFlags)

	// initialize player and game
	p1 := tictactoe.NewPlayer("Player 1", "X")
	p2 := tictactoe.NewPlayer("Player 2", "O")
	g := tictactoe.NewGame(*p1, *p2, l)

	// game start
	g.Start()
}
