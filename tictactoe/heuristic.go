package tictactoe

import (
	"fmt"
)

func EvaluateAIPositions(game gameState) {
	game.aiPositions.ForEach(func(p Position) {
		evaluatePosition(p, game.aiPositions, game.opponentPositions, game.area)
	})
}

func evaluatePosition(p Position, myPos PositionSet, opponentPos PositionSet, area Area) {
	discovery := discover(p, Left, p, Right, p, 0, myPos, opponentPos, area, make(map[segment]discovery))

	if discovery.status == 1 {
		fmt.Println("Nenasel jsem potencialni vyhru")
	} else {
		fmt.Printf("counts: %v\n", discovery.myPosCounts)
	}
}

type discovery struct {
	status      uint8
	myPosCounts []uint8
}

type segment struct {
	p1 Position
	p2 Position
}

func discover(p1 Position, d1 PositionDirection, p2 Position, d2 PositionDirection, newPos Position, distance uint8, myPos PositionSet, opponentPos PositionSet, area Area, memo map[segment]discovery) discovery {

	if !p1.isWithinBounds(area) || !p2.isWithinBounds(area) {
		return discovery{1, []uint8{}}
	}

	if opponentPos.Has(p1) || opponentPos.Has(p2) {
		return discovery{1, []uint8{}}
	}

	if distance == 4 {
		if myPos.Has(newPos) {
			return discovery{0, []uint8{1}}
		}
		return discovery{0, []uint8{0}}
	}

	newP1 := p1.next(d1)
	newSegment1 := segment{newP1, p2}

	discovery1, has1 := memo[newSegment1]

	if !has1 {
		memo[newSegment1] = discover(newP1, d1, p2, d2, newP1, distance+1, myPos, opponentPos, area, memo)
		discovery1 = memo[newSegment1]

		if discovery1.status == 0 && myPos.Has(newPos) {
			incrementDiscovery(&discovery1)
		}
	}

	newP2 := p2.next(d2)
	newSegment2 := segment{p1, newP2}

	discovery2, has2 := memo[newSegment2]

	if !has2 {
		memo[newSegment2] = discover(p1, d1, newP2, d2, newP2, distance+1, myPos, opponentPos, area, memo)
		discovery2 = memo[newSegment2]

		if discovery2.status == 0 && myPos.Has(newPos) {
			incrementDiscovery(&discovery2)
		}
	}

	if discovery1.status == 0 && discovery2.status == 0 {
		return *mergeDiscovery(&discovery1, &discovery2)
	} else if discovery1.status == 0 {
		return discovery1
	} else if discovery2.status == 0 {
		return discovery2
	} else {
		return discovery{1, []uint8{}}
	}
}

func incrementDiscovery(d *discovery) {
	newIncrements := make([]uint8, 0, len(d.myPosCounts))

	for _, c := range d.myPosCounts {
		newIncrements = append(newIncrements, uint8(c+1))
	}

	d.myPosCounts = newIncrements
}

func mergeDiscovery(d1 *discovery, d2 *discovery) *discovery {
	d1.myPosCounts = append(d1.myPosCounts, d2.myPosCounts...)
	return d1
}
