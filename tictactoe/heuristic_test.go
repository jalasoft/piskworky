package tictactoe

import (
	"testing"
)

func assertThat(d discovery, t *testing.T) *DiscoveryAssertion {
	return &DiscoveryAssertion{
		d: d,
		t: t,
	}
}

type DiscoveryAssertion struct {
	d discovery
	t *testing.T
}

func (d *DiscoveryAssertion) hasSegment(expectedSegment segment, expectedCount uint8) *DiscoveryAssertion {

	actualCount, exists := d.d.segments[expectedSegment]

	if !exists {
		d.t.Errorf("%v was not discovered\n", expectedSegment)
		return d
	}

	if expectedCount != actualCount {
		d.t.Errorf("Counts of segment %v do NOT match. Expected: %d, actual: %d\n", expectedSegment, expectedCount, actualCount)
	}

	return d
}

func (d *DiscoveryAssertion) hasTotal(expected uint8) *DiscoveryAssertion {

	if len(d.d.segments) != int(expected) {
		d.t.Errorf("Discovery has unexpected number of segments. Expected: %d, actual: %d\n", expected, len(d.d.segments))
	}

	return d
}

//--------------------TESTS------------------------

//|ox.xx.|
func TestEvaluateAIPositions1(t *testing.T) {
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

	discovery := EvaluateAIPositions(state)

	assertThat(discovery, t).hasSegment(segment{Position{0, 1}, Position{0, 5}}, 3)
	assertThat(discovery, t).hasTotal(1)
}

//|...x.....|
func TestEvaluateAIPositions2(t *testing.T) {
	myPos := newPositionSet()
	myPos.Add(Position{0, 3})

	opponentPos := newPositionSet()
	area := Area{1, 9}

	state := gameState{
		area:              area,
		aiPositions:       myPos,
		opponentPositions: opponentPos,
	}

	discovery := EvaluateAIPositions(state)

	assertThat(discovery, t).hasSegment(segment{Position{0, 0}, Position{0, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 1}, Position{0, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 2}, Position{0, 6}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 3}, Position{0, 7}}, 1)
	assertThat(discovery, t).hasTotal(4)
}

//|o.x.x.x|
func TestEvaluateAIPositions3(t *testing.T) {
	myPos := newPositionSet()
	myPos.Add(Position{0, 2})
	myPos.Add(Position{0, 4})
	myPos.Add(Position{0, 6})

	opponentPos := newPositionSet()
	opponentPos.Add(Position{0, 0})
	area := Area{1, 7}

	state := gameState{
		area:              area,
		aiPositions:       myPos,
		opponentPositions: opponentPos,
	}

	discovery := EvaluateAIPositions(state)

	assertThat(discovery, t).hasSegment(segment{Position{0, 1}, Position{0, 5}}, 2)
	assertThat(discovery, t).hasSegment(segment{Position{0, 2}, Position{0, 6}}, 3)
	assertThat(discovery, t).hasTotal(2)
}

//|_____x_|
//|____x__|
//|_______|
//|__x____|
//|_______|
//|oo_____|
//|_______|
func TestEvaluateAIPositions4(t *testing.T) {
	myPos := newPositionSet()
	myPos.Add(Position{0, 5})
	myPos.Add(Position{1, 4})
	myPos.Add(Position{3, 2})

	opponentPos := newPositionSet()
	opponentPos.Add(Position{5, 0})
	opponentPos.Add(Position{5, 1})
	area := Area{7, 7}

	state := gameState{
		area:              area,
		aiPositions:       myPos,
		opponentPositions: opponentPos,
	}

	discovery := EvaluateAIPositions(state)

	t.Logf("%v\n", discovery)

	assertThat(discovery, t).hasSegment(segment{Position{0, 1}, Position{0, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 2}, Position{0, 6}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 5}, Position{4, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{1, 0}, Position{1, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{1, 1}, Position{1, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{1, 2}, Position{1, 6}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 4}, Position{4, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{1, 4}, Position{5, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{3, 0}, Position{3, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{3, 1}, Position{3, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{3, 2}, Position{3, 6}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 2}, Position{4, 2}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{1, 2}, Position{5, 2}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{2, 2}, Position{6, 2}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 5}, Position{4, 1}}, 3)
	assertThat(discovery, t).hasSegment(segment{Position{1, 0}, Position{5, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{2, 1}, Position{6, 5}}, 1)

	assertThat(discovery, t).hasTotal(17)
}

//|x____x|
//|ox___x|
//|___x__|
//|____x_|
//|____ox|
func TestEvaluateAIPositions5(t *testing.T) {
	myPos := newPositionSet()
	myPos.Add(Position{0, 0})
	myPos.Add(Position{0, 5})
	myPos.Add(Position{1, 1})
	myPos.Add(Position{1, 5})
	myPos.Add(Position{2, 3})
	myPos.Add(Position{3, 4})
	myPos.Add(Position{4, 5})

	opponentPos := newPositionSet()
	opponentPos.Add(Position{1, 0})
	opponentPos.Add(Position{4, 4})
	area := Area{5, 6}

	state := gameState{
		area:              area,
		aiPositions:       myPos,
		opponentPositions: opponentPos,
	}

	discovery := EvaluateAIPositions(state)

	t.Logf("%v\n", discovery)

	assertThat(discovery, t).hasSegment(segment{Position{0, 0}, Position{0, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 1}, Position{0, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{1, 1}, Position{1, 5}}, 2)
	assertThat(discovery, t).hasSegment(segment{Position{2, 0}, Position{2, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{2, 1}, Position{2, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{3, 0}, Position{3, 4}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{3, 1}, Position{3, 5}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 5}, Position{4, 5}}, 3)
	assertThat(discovery, t).hasSegment(segment{Position{0, 5}, Position{4, 1}}, 2)
	assertThat(discovery, t).hasSegment(segment{Position{0, 3}, Position{4, 3}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 1}, Position{4, 1}}, 1)
	assertThat(discovery, t).hasSegment(segment{Position{0, 1}, Position{4, 5}}, 3)
	assertThat(discovery, t).hasTotal(12)
}
