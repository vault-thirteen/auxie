package mcci

const (
	Err_AreaIsMalformed = "area is malformed"
)

// Area is a rectangle area of pixels represented by its diagonal corners.
// First pixel in the corner has coordinates (0;0).
// Each pixel has width and height equal to 1.
// Example: area from point (0;0) to point (1;1) has width and height = 2.
type Area[T ICoordinate] struct {
	BottomLeft Point[T]
	TopRight   Point[T]
}

func NewAreaFromPoints[T ICoordinate](p1, p2 Point[T]) (a *Area[T]) {
	a = &Area[T]{
		BottomLeft: NewPoint[T](min(p1.X, p2.X), min(p1.Y, p2.Y)),
		TopRight:   NewPoint[T](max(p1.X, p2.X), max(p1.Y, p2.Y)),
	}
	return a
}
func NewAreaFromPointAndVector[T ICoordinate](p Point[T], v Vector[T]) (r *Area[T]) {
	var p2 = p
	p2.Shift(v)
	return NewAreaFromPoints(p, p2)
}
func NewAreaFromDimensions[T ICoordinate](width, height T) (a *Area[T]) {
	if (width <= 0) || (height <= 0) {
		return nil
	}

	var p1 = NewPoint[T](0, 0)
	var p2 = p1
	var delta = NewVector[T](width-1, height-1)
	p2.Shift(delta)

	return NewAreaFromPoints(p1, p2)
}

func (a *Area[T]) IsMalformed() bool {
	return (a.Width() <= 0) || (a.Height() <= 0)
}
func (a *Area[T]) IsEqualTo(that *Area[T]) bool {
	return a.BottomLeft.IsEqualTo(that.BottomLeft) && a.TopRight.IsEqualTo(that.TopRight)
}
func (a *Area[T]) HasSameDimension(that *Area[T]) bool {
	return (a.Width() == that.Width()) && (a.Height() == that.Height())
}
func (a *Area[T]) Width() (width T) {
	return a.TopRight.X - a.BottomLeft.X + 1 // Each pixel has width = 1.
}
func (a *Area[T]) Height() (height T) {
	return a.TopRight.Y - a.BottomLeft.Y + 1 // Each pixel has height = 1.
}
func (a *Area[T]) Shift(v Vector[T]) {
	a.BottomLeft.Shift(v)
	a.TopRight.Shift(v)
}
func (a *Area[T]) UnShift(v Vector[T]) {
	a.BottomLeft.UnShift(v)
	a.TopRight.UnShift(v)
}
func (a *Area[T]) IsVertical() bool {
	return a.Height() > a.Width()
}
func (a *Area[T]) IsHorizontal() bool {
	return a.Width() > a.Height()
}
func (a *Area[T]) IsSquare() bool {
	return a.Width() == a.Height()
}
func (a *Area[T]) ContainsPoint(p Point[T]) bool {
	return (a.BottomLeft.X <= p.X) && (p.X <= a.TopRight.X) &&
		(a.BottomLeft.Y <= p.Y) && (p.Y <= a.TopRight.Y)
}
func (a *Area[T]) ContainsArea(that *Area[T]) bool {
	return a.ContainsPoint(that.BottomLeft) && a.ContainsPoint(that.TopRight)
}
