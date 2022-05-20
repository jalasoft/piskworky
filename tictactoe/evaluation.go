package tictactoe

func evaluate(area Area, positions PositionSet, oponentPositions PositionSet) {
	positions.ForEach(func(p Position) {

	})
}

func investigatePosition(p Position, area Area, positions PositionSet, oponentPositions PositionSet) {

	topDepth, topCount := exploreDirection(p, Top, area, 5, positions, oponentPositions)
	bottomDepth, bottomCount := exploreDirection(p, Bottom, area, 5, positions, oponentPositions)

}

func exploreDirection(p Position, direction PositionDirection, area Area, maxDepth uint8, positions PositionSet, oponentPositions PositionSet) (depth uint8, posCount uint8) {
	if maxDepth == 0 {
		return 0, 0
	}

	if !p.isWithinBounds(area) {
		return 0, 0
	}

	if oponentPositions.Has(p) {
		return 0, 0
	}

	depth, posCount = exploreDirection(p.next(direction), direction, area, maxDepth-1, positions, oponentPositions)

	if positions.Has(p) {
		return depth + 1, posCount + 1
	}

	return depth + 1, posCount
}
