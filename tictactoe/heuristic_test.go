package tictactoe

import (
	"testing"
)

//|ox xx |
func TestEvaluateAIPositions(t *testing.T) {
	myPos := newPositionSet()
	myPos.Add(Position{0, 1})
	myPos.Add(Position{0, 3})
	myPos.Add(Position{0, 4})

	opponentPos := newPositionSet()
	opponentPos.Add(Position{0, 0})
	area := Area{1, 6}

	state := gameState{
		area:              area,
		aiPositions:       myPos,
		opponentPositions: opponentPos,
	}

	EvaluateAIPositions(state)
}
