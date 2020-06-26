package tictactoe

import (
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	NAME = "PlayerName"
	MARK = "X"
	ROW  = 1
	COL  = 2
)

var player = NewPlayer(NAME, MARK)

func TestPlayerValidInput(t *testing.T) {
	temp, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}

	defer temp.Close()

	_, err = io.WriteString(temp, fmt.Sprintf("%d %d\n", ROW, COL))
	if err != nil {
		t.Fatal(err)
	}

	_, err = temp.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	i, j, err := player.GetMove(temp)

	assert.Equal(t, ROW, i)
	assert.Equal(t, COL, j)
	assert.Nil(t, err)
}

func TestPlayerInvalidInput(t *testing.T) {
	temp, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}

	defer temp.Close()

	_, err = io.WriteString(temp, fmt.Sprintf("%v %v\n", ROW, "abc"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = temp.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	i, j, err := player.GetMove(temp)

	assert.Equal(t, 0, i)
	assert.Equal(t, 0, j)
	assert.NotNil(t, err)
}

func TestPlayerName(t *testing.T) {
	assert.Equal(t, NAME, player.Name())
}

func TestPlayerMark(t *testing.T) {
	assert.Equal(t, MARK, player.Mark())
}
