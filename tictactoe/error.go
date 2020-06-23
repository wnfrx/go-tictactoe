package tictactoe

import "errors"

var (
	ErrInvalidInputFormat = errors.New("Invalid input format")
	ErrInvalidInputCount  = errors.New("Input must be row & column")
)
