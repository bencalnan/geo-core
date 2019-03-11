package main

import "math"

//Line - A edge created from 2 Points.
type Line [2]Point

func createLine(a Point, b Point) Line {
	return Line{a, b}
}

//The Euclidean distance between two points of the plane with Cartesian coordinates
//Calcuate Hypoteneuse between two Points.
func (l *Line) getLineLength() float64 {
	return math.Sqrt(math.Pow((l[1].X-l[0].X), 2) + math.Pow((l[1].Y-l[0].Y), 2))
}
