package tictactoe

import (
	"fmt"
)

type Position struct {
	Row    uint16
	Column uint16
}

func (p Position) String() string {
	return fmt.Sprintf("Pos[row=%d,col=%d]", p.Row, p.Column)
}

func (p Position) isWithinBounds(area Area) bool {
	if p.Column < 0 || p.Column >= area.Columns {
		return false
	}

	if p.Row < 0 || p.Row >= area.Rows {
		return false
	}
	return true
}

func (p Position) topLeft() Position {
	return Position{p.Row - 1, p.Column - 1}
}

func (p Position) top() Position {
	return Position{p.Row - 1, p.Column}
}

func (p Position) topRight() Position {
	return Position{p.Row - 1, p.Column + 1}
}

func (p Position) right() Position {
	return Position{p.Row, p.Column + 1}
}

func (p Position) bottomRight() Position {
	return Position{p.Row + 1, p.Column + 1}
}

func (p Position) bottom() Position {
	return Position{p.Row + 1, p.Column}
}

func (p Position) bottomLeft() Position {
	return Position{p.Row + 1, p.Column - 1}
}

func (p Position) left() Position {
	return Position{p.Row, p.Column - 1}
}

func (p Position) next(d PositionDirection) Position {
	switch d {
	case Top:
		return p.top()
	case TopRight:
		return p.topRight()
	case Right:
		return p.right()
	case BottomRight:
		return p.bottomRight()
	case Bottom:
		return p.bottom()
	case BottomLeft:
		return p.bottomLeft()
	case Left:
		return p.left()
	case TopLeft:
		return p.topLeft()
	default:
		panic(fmt.Sprintf("Unknown direction %d", d))
	}
}

func (p Position) surroundings() []Position {
	return []Position{
		p.topLeft(),
		p.top(),
		p.topRight(),
		p.right(),
		p.bottomRight(),
		p.bottom(),
		p.bottomLeft(),
		p.left(),
	}
}
