package tictactoe

import (
	"fmt"
)

type GameStatus uint8

const (
	TIE GameStatus = iota
	AI_WON
	OPONNENT_WON
	PLAYING
)

type Area struct {
	Rows    uint16
	Columns uint16
}

func (a Area) String() string {
	return fmt.Sprintf("Area[rows=%d,columns=%d]", a.Rows, a.Columns)
}

type Game interface {
	Area() Area
	OpponentMove(p Position) GameStatus
	AiMove() (GameStatus, Position)

	NextPositions() []Position
}

func NewGame(rows uint16, columns uint16) Game {
	oponentPositions := newPositionSet()
	aiPositions := newPositionSet()
	area := Area{rows, columns}

	nextPositions := &NextPositionSet{
		area:              area,
		oponnentPositions: oponentPositions,
		aiPositions:       aiPositions,
		freePositions:     make([]*ValuedPosition, 0),
	}

	return &game{
		gameState: gameState{
			area:              area,
			opponentPositions: oponentPositions,
			aiPositions:       aiPositions},
		nextPositions: nextPositions,
		status:        PLAYING,
	}
}
