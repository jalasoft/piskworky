package tictactoe

import (
	"fmt"
	"log"
)

type GameStatus uint8

const (
	TIE GameStatus = iota
	AI_WON
	OPONNENT_WON
	PLAYING
)

func NewGame(rows uint16, columns uint16) *Game {
	oponentPositions := newPositionSet()
	aiPositions := newPositionSet()
	area := Area{rows, columns}

	nextPositions := &NextPositionSet{
		area:              area,
		oponnentPositions: oponentPositions,
		aiPositions:       aiPositions,
		freePositions:     make([]*ValuedPosition, 0),
	}

	return &Game{
		area:             area,
		oponentPositions: oponentPositions,
		aiPositions:      aiPositions,
		nextPositions:    nextPositions,
		status:           PLAYING,
	}
}

type Area struct {
	Rows    uint16
	Columns uint16
}

func (a Area) String() string {
	return fmt.Sprintf("Area[rows=%d,columns=%d]", a.Rows, a.Columns)
}

type Game struct {
	area             Area
	oponentPositions PositionSet
	aiPositions      PositionSet
	nextPositions    *NextPositionSet
	status           GameStatus
}

func (g *Game) NextPositions() []Position {
	return g.nextPositions.Positions()
}

func (g *Game) OponentMove(pos Position) GameStatus {
	log.Printf("Oponent move %s", pos)

	if g.status != PLAYING {
		return g.status
	}

	if !pos.isWithinBounds(g.area) {
		panic(fmt.Sprintf("Move %s is not within bounds %v.", pos, g.area))
	}

	if g.oponentPositions.Has(pos) || g.aiPositions.Has(pos) {
		panic(fmt.Sprintf("Move %s is already used.", pos))
	}

	g.nextPositions.Remove(pos)
	g.nextPositions.AddNeighboursOf(pos)
	g.oponentPositions.Add(pos)

	if g.oponentPositions.HasTerminalVector() {
		return OPONNENT_WON
	}

	return PLAYING
}

func (g *Game) AiMove() (GameStatus, Position) {

	var pos Position

	if g.oponentPositions.IsEmpty() && g.aiPositions.IsEmpty() {
		pos = Position{g.area.Rows / 2, g.area.Columns / 2} //first move
	} else {
		pos = g.computeAiMove()
	}

	if g.aiPositions.HasTerminalVector() {
		return AI_WON, Position{}
	}

	g.nextPositions.Remove(pos)
	g.aiPositions.Add(pos)
	g.nextPositions.AddNeighboursOf(pos)

	return PLAYING, pos
}

func (g *Game) computeAiMove() Position {
	g.nextPositions.Evaluate()
	return g.nextPositions.First()
}
