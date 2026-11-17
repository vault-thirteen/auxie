package c

import (
	"golang.org/x/exp/constraints"
)

type Pos[T constraints.Integer | constraints.Float] struct {
	X T
	Y T
}

func (p *Pos[T]) Shift(dx, dy T) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func (p *Pos[T]) ShiftVec(d Vec[T]) {
	p.X = p.X + d.X
	p.Y = p.X + d.Y
}

func (p *Pos[T]) UnshiftVec(d Vec[T]) {
	p.X = p.X - d.X
	p.Y = p.Y - d.Y
}
