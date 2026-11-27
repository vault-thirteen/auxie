package mcci

type Point[T ICoordinate] struct {
	X T
	Y T
}

func NewPoint[T ICoordinate](x T, y T) Point[T] {
	return Point[T]{X: x, Y: y}
}

func (p *Point[T]) IsEqualTo(that Point[T]) bool {
	return (p.X == that.X) && (p.Y == that.Y)
}
func (p *Point[T]) IsLefterThan(that Point[T]) bool {
	return p.X < that.X
}
func (p *Point[T]) IsRighterThan(that Point[T]) bool {
	return p.X > that.X
}
func (p *Point[T]) IsLowerThan(that Point[T]) bool {
	return p.Y < that.Y
}
func (p *Point[T]) IsUpperThan(that Point[T]) bool {
	return p.Y > that.Y
}
func (p *Point[T]) Shift(v Vector[T]) {
	p.X = p.X + v.X
	p.Y = p.Y + v.Y
}
func (p *Point[T]) UnShift(v Vector[T]) {
	p.X = p.X - v.X
	p.Y = p.Y - v.Y
}
func (p *Point[T]) IsZero() bool {
	return (p.X == 0) && (p.Y == 0)
}
