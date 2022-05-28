package tictactoe

import (
	"fmt"
)

type segment struct {
	p1 Position
	p2 Position
}

func (s *segment) String() string {
	return fmt.Sprintf("%v-%v", s.p1, s.p2)
}

type discovery struct {
	status   uint8
	segments map[segment]uint8
}

func (d *discovery) merge(another *discovery) discovery {
	newSegments := make(map[segment]uint8)

	for segment, count := range d.segments {
		newSegments[segment] = count
	}

	for segment, count := range another.segments {
		if _, has := newSegments[segment]; !has {
			newSegments[segment] = count
		}
	}
	return discovery{0, newSegments}
}

func (d *discovery) increment() discovery {
	newSegments := make(map[segment]uint8)

	for segment, count := range d.segments {
		newSegments[segment] = count + 1
	}
	return discovery{0, newSegments}
}

func (d discovery) String() string {
	str := "discovery{\n"
	for s, c := range d.segments {
		str = fmt.Sprintf("%s\t%v:%d\n", str, s, c)
	}
	str = str + "}\n"
	return str
}

//----------------------------------------------------------------

func EvaluateAIPositions(game gameState) discovery {
	all := discovery{0, make(map[segment]uint8)}
	memo := make(map[segment]discovery)

	game.aiPositions.ForEach(func(p Position) {
		discovery := evaluatePosition(p, game.aiPositions, game.opponentPositions, game.area, memo)
		all = all.merge(&discovery)
	})

	return all
}

func evaluatePosition(p Position, myPos PositionSet, opponentPos PositionSet, area Area, memo map[segment]discovery) discovery {
	d := discovery{}

	lr := discoverSegments(p, Left, Right, 0, myPos, opponentPos, area, memo)
	d = d.merge(&lr)

	tb := discoverSegments(p, Top, Bottom, 0, myPos, opponentPos, area, memo)
	d = d.merge(&tb)

	tlbr := discoverSegments(p, TopLeft, BottomRight, 0, myPos, opponentPos, area, memo)
	d = d.merge(&tlbr)

	trbl := discoverSegments(p, TopRight, BottomLeft, 0, myPos, opponentPos, area, memo)
	d = d.merge(&trbl)

	return d
}

func discoverSegments(p Position, d1 PositionDirection, d2 PositionDirection, distance uint8, myPos PositionSet, opponentPos PositionSet, area Area, memo map[segment]discovery) discovery {
	discovery := discover(segment{p, p}, d1, d2, distance, myPos, opponentPos, area, memo)

	if discovery.status == 0 && myPos.Has(p) {
		discovery = discovery.increment()
	}

	return discovery
}

func discover(s segment, d1 PositionDirection, d2 PositionDirection, distance uint8, myPos PositionSet, opponentPos PositionSet, area Area, memo map[segment]discovery) discovery {

	if !s.p1.isWithinBounds(area) || !s.p2.isWithinBounds(area) {
		return discovery{1, nil}
	}

	if opponentPos.Has(s.p1) || opponentPos.Has(s.p2) {
		return discovery{1, nil}
	}

	if distance == 4 {
		return discovery{0, map[segment]uint8{s: 0}}
	}

	newP1 := s.p1.next(d1)
	newSegment1 := segment{newP1, s.p2}

	discovery1, has1 := memo[newSegment1]

	if !has1 {
		memo[newSegment1] = discover(newSegment1, d1, d2, distance+1, myPos, opponentPos, area, memo)
		discovery1 = memo[newSegment1]
	}

	if discovery1.status == 0 && myPos.Has(newP1) {
		discovery1 = discovery1.increment()
	}

	newP2 := s.p2.next(d2)
	newSegment2 := segment{s.p1, newP2}

	discovery2, has2 := memo[newSegment2]

	if !has2 {
		memo[newSegment2] = discover(newSegment2, d1, d2, distance+1, myPos, opponentPos, area, memo)
		discovery2 = memo[newSegment2]
	}

	if discovery2.status == 0 && myPos.Has(newP2) {
		discovery2 = discovery2.increment()
	}

	if discovery1.status == 0 && discovery2.status == 0 {
		return discovery1.merge(&discovery2)
	} else if discovery1.status == 0 {
		return discovery1
	} else if discovery2.status == 0 {
		return discovery2
	} else {
		return discovery{1, nil}
	}
}
