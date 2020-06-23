package tictactoe

import (
	"fmt"
	"log"
)

type Game struct {
	p1            Player // player 1
	p2            Player // player 2
	currentPlayer Player // current player
	board         *Board // game board
	round         int    // current round
}

func NewGame(p1, p2 Player) *Game {
	return &Game{
		p1:            p1,
		p2:            p2,
		currentPlayer: p1,
		board:         newBoard(),
		round:         1,
	}
}

type IGame interface {
	Start()
}

func (g *Game) Start() {
	fmt.Println("TicTacToe, Start!")
	fmt.Println("=================")

	g.board.Draw()

	for !g.isOver() {
		// Print game info
		g.info()

		// Player's turn, wait for input
		i, j, err := g.currentPlayer.GetMove()
		if err != nil {
			fmt.Println("Invalid input format, please try again")
			log.Printf("Error message: %v\n", err)
			continue
		}

		if i < 0 || j < 0 || i > 2 || j > 2 {
			fmt.Println("Input can be number 0-2 only, please try again")
			continue
		}

		if g.board[i][j] != "_" {
			fmt.Println("Board cell already used, please try again")
			continue
		}

		// Assign player's mark to board[i][j]
		g.board[i][j] = g.currentPlayer.Mark()

		// Switch player
		g.switchPlayer()
		g.round++

		// Draw current board state
		g.board.Draw()
	}

	fmt.Println("End Game")
}

func (g *Game) switchPlayer() {
	if g.currentPlayer == g.p1 {
		g.currentPlayer = g.p2
	} else {
		g.currentPlayer = g.p1
	}
}

func (g *Game) isOver() bool {
	return g.board.HasWinner(g.p1.Mark(), g.p2.Mark()) || !g.board.HasEmptyCell()
}

func (g *Game) info() {
	fmt.Printf("Turn #%v\n", g.round)
	fmt.Printf("Player: %v\n", g.currentPlayer.Name())
}
