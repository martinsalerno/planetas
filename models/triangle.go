package models

import (
	"math"
	"planets/constants"
)

type Triangle struct {
	Vertex1, Vertex2, Vertex3 Point
}

func (triangle Triangle) Stats() (float64, bool, bool) {
	area := triangle.area()
	pointsAligned := triangle.pointsAreAligned(area)
	zeroPointIncluded := triangle.includesPoint(Point{X: 0.0, Y: 0.0})

	return area, pointsAligned, zeroPointIncluded
}

func (triangle Triangle) area() float64 {
	area1 := triangle.Vertex1.X * (triangle.Vertex2.Y - triangle.Vertex3.Y)
	area2 := triangle.Vertex2.X * (triangle.Vertex3.Y - triangle.Vertex1.Y)
	area3 := triangle.Vertex3.X * (triangle.Vertex1.Y - triangle.Vertex2.Y)

	return math.Abs((area1 + area2 + area3) * 0.5)
}

func (triangle Triangle) includesPoint(point Point) bool {
	s1 := triangle.Vertex3.Y - triangle.Vertex1.Y
	s2 := triangle.Vertex3.X - triangle.Vertex1.X
	s3 := triangle.Vertex2.Y - triangle.Vertex1.Y
	s4 := point.Y - triangle.Vertex1.Y

	w0 := (s3*s2 - (triangle.Vertex2.X-triangle.Vertex1.X)*s1)

	if isZero(w0) {
		return true
	}

	w1 := (triangle.Vertex1.X*s1 + s4*s2 - point.X*s1) / w0
	w2 := (s4 - w1*s3) / s1

	return w1 >= 0 && w2 >= 0 && (w1+w2) <= 1
}

func (triangle Triangle) pointsAreAligned(area float64) bool {
	return isZero(area)
}

func isZero(num float64) bool {
	return math.Abs(num) <= constants.FloatPrecision
}
