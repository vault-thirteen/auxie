package mcci

type Vector[T ICoordinate] struct {
	X T
	Y T
}

func NewVector[T ICoordinate](x T, y T) Vector[T] {
	return Vector[T]{X: x, Y: y}
}

func (v *Vector[T]) IsEqualTo(that Vector[T]) bool {
	return (v.X == that.X) && (v.Y == that.Y)
}
func (v *Vector[T]) IsHorizontal() bool {
	return v.Y == 0
}
func (v *Vector[T]) IsVertical() bool {
	return v.X == 0
}
func (v *Vector[T]) Shift(delta Vector[T]) {
	v.X = v.X + delta.X
	v.Y = v.Y + delta.Y
}
func (v *Vector[T]) UnShift(delta Vector[T]) {
	v.X = v.X - delta.X
	v.Y = v.Y - delta.Y
}
func (v *Vector[T]) IsZero() bool {
	return (v.X == 0) && (v.Y == 0)
}
