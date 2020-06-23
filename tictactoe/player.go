package tictactoe

import (
	"fmt"
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

func (p *Player) GetMove() (int, int, error) {
	var i, j int

	fmt.Print("Enter row and column number [0-2]: ")

	input, err := fmt.Scanf("%d %d\n", &i, &j)
	if err != nil {
		return 0, 0, ErrInvalidInputFormat
	}

	if input != 2 {
		return 0, 0, ErrInvalidInputCount
	}

	return i, j, nil
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Mark() string {
	return p.mark
}
