package tictactoe

import (
	"fmt"
	"log"
)

type gameState struct {
	area              Area
	opponentPositions PositionSet
	aiPositions       PositionSet
}

type game struct {
	gameState
	nextPositions *NextPositionSet
	status        GameStatus
}

func (g *game) NextPositions() []Position {
	return g.nextPositions.Positions()
}

func (g *game) Area() Area {
	return g.area
}

func (g *game) OpponentMove(pos Position) GameStatus {
	log.Printf("Opponent move %s", pos)

	if g.status != PLAYING {
		return g.status
	}

	if !pos.isWithinBounds(g.area) {
		panic(fmt.Sprintf("Move %s is not within bounds %v.", pos, g.area))
	}

	if g.opponentPositions.Has(pos) || g.aiPositions.Has(pos) {
		panic(fmt.Sprintf("Move %s is already used.", pos))
	}

	g.nextPositions.Remove(pos)
	g.nextPositions.AddNeighboursOf(pos)
	g.opponentPositions.Add(pos)

	if g.opponentPositions.HasTerminalVector() {
		return OPONNENT_WON
	}

	return PLAYING
}

func (g *game) AiMove() (GameStatus, Position) {

	var pos Position

	if g.opponentPositions.IsEmpty() && g.aiPositions.IsEmpty() {
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

func (g *game) computeAiMove() Position {
	g.nextPositions.Evaluate()
	return g.nextPositions.First()
}

type PositionDirection uint

const (
	Top PositionDirection = iota
	TopRight
	Right
	BottomRight
	Bottom
	BottomLeft
	Left
	TopLeft
)

func (p PositionDirection) opposite() PositionDirection {
	switch p {
	case Top:
		return Bottom
	case TopRight:
		return BottomLeft
	case Right:
		return Left
	case BottomRight:
		return TopLeft
	case Bottom:
		return Top
	case BottomLeft:
		return TopRight
	case Left:
		return Right
	case TopLeft:
		return BottomRight
	default:
		panic("Unknown direction")
	}
}
