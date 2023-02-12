package rng

import (
	"errors"
	"math"
)

const (
	ErrStartIsReversed = "start is reversed"
)

type Range struct {
	start  float64
	end    float64
	middle float64
	radius float64
}

func New(start float64, end float64) (r *Range, err error) {
	if end < start {
		return nil, errors.New(ErrStartIsReversed)
	}

	middle := (start + end) / 2

	return &Range{
		start:  start,
		end:    end,
		middle: middle,
		radius: middle - start,
	}, nil
}

func (r *Range) Contains(value float64) (contains bool) {
	return (r.start <= value) && (value <= r.end)
}

func (r *Range) GetStart() (mid float64) {
	return r.start
}

func (r *Range) GetEnd() (mid float64) {
	return r.end
}

func (r *Range) GetMiddle() (mid float64) {
	return r.middle
}

func (r *Range) GetRadius() (rad float64) {
	return r.radius
}

func (r *Range) HasIntersectionWith(that *Range) (intersects bool) {
	return math.Abs(r.middle-that.middle) <= (r.radius + that.radius)
}

func IsSequence(r1 *Range, r2 *Range) (isSequence bool) {
	return r1.end == r2.start
}
