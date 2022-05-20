package tictactoe

import (
	"log"
	"sort"
)

type ValuedPosition struct {
	Position
	value uint8
}

type NextPositionSet struct {
	area              Area
	oponnentPositions PositionSet
	aiPositions       PositionSet
	freePositions     []*ValuedPosition
}

func (n *NextPositionSet) Positions() []Position {
	pos := make([]Position, len(n.freePositions))

	for _, val := range n.freePositions {
		pos = append(pos, Position{val.Row, val.Column})
	}

	return pos
}

func (n *NextPositionSet) Remove(pos Position) {
	var index int = -1

	for i, p := range n.freePositions {
		if p.Row == pos.Row && p.Column == pos.Column {
			index = i
			break
		}
	}

	if index == -1 {
		log.Printf("%v not found in free positions %v", pos, n.freePositions)
		return
	}

	n.freePositions = append(n.freePositions[:index], n.freePositions[:index+1]...)
}

func (n *NextPositionSet) AddNeighboursOf(p Position) {
	for _, neighbour := range p.surroundings() {
		if !neighbour.isWithinBounds(n.area) {
			continue
		}

		if n.oponnentPositions.Has(p) {
			continue
		}

		if n.aiPositions.Has(p) {
			continue
		}

		n.freePositions = append(n.freePositions, &ValuedPosition{neighbour, 0})
	}
}

func (n *NextPositionSet) Evaluate() {
	//for each free position, count number of neighbouring non-free positions
	for _, pos := range n.freePositions {
		for _, neighbour := range pos.surroundings() {
			var value uint8 = 0
			if n.aiPositions.Has(neighbour) {
				value++
			}
			if n.oponnentPositions.Has(neighbour) {
				value++
			}
			pos.value = value
		}
	}

	var dataToSort byValueSort = byValueSort(n.freePositions)
	sort.Sort(dataToSort)
}

func (n *NextPositionSet) First() Position {
	return n.freePositions[0].Position
}

//------------------------------------------------------------

type byValueSort []*ValuedPosition

func (b byValueSort) Len() int {
	return len(b)
}

func (b byValueSort) Less(i, j int) bool {
	return b[i].value < b[j].value
}

func (b byValueSort) Swap(i, j int) {
	var temp *ValuedPosition = b[i]
	b[i] = b[j]
	b[j] = temp
}
