package tictactoe

type PositionSet map[Position]bool

func newPositionSet() PositionSet {
	return PositionSet(make(map[Position]bool))
}

func (p *PositionSet) IsEmpty() bool {
	return len(*p) == 0
}

func (p PositionSet) Has(position Position) bool {
	_, has := p[position]
	return has
}

func (p PositionSet) Add(pos Position) {
	p[pos] = true
}

func (p PositionSet) ForEach(cb func(p Position)) {

}

func (p *PositionSet) HasTerminalVector() bool {
	for position, _ := range *p {
		if p.investigateTerminalVector(position, 4, Top) {
			return true
		}
		if p.investigateTerminalVector(position, 4, TopLeft) {
			return true
		}
		if p.investigateTerminalVector(position, 4, Left) {
			return true
		}
		if p.investigateTerminalVector(position, 4, BottomLeft) {
			return true
		}
		if p.investigateTerminalVector(position, 4, Bottom) {
			return true
		}
	}
	return false
}

func (p *PositionSet) investigateTerminalVector(pos Position, depth uint, direction PositionDirection) bool {
	if depth == 0 {
		return true
	}

	if !p.Has(pos) {
		return false
	}

	return p.investigateTerminalVector(pos.next(direction), depth-1, direction)
}
