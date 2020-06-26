package tictactoe

import (
	"fmt"
	"os"
)

type Player struct {
	name string
	mark string
}

func NewPlayer(n string, m string) *Player {
	return &Player{
		name: n,
		mark: m,
	}
}

type IPlayer interface {
	GetMove() (int, int, error)
	Name() string
	Mark() string
}

func (p *Player) GetMove(in *os.File) (int, int, error) {
	if in == nil {
		in = os.Stdin
	}

	var i, j int

	fmt.Print("Enter row and column number [0-2]: ")

	_, err := fmt.Fscanf(in, "%d %d\n", &i, &j)
	if err != nil {
		return 0, 0, ErrInvalidInputFormat
	}

	return i, j, nil
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Mark() string {
	return p.mark
}
