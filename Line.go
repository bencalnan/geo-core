package main

import "math"

//Line - A edge created from 2 Points.
type Line [2]Point

func (Line) geomType() string {
	return "line"
}

//GetBBox - Return Bounding box of Line
func (l *Line) GetBBox() BoundingBox {
	maxX := math.Max(l[0].X, l[1].X)
	maxY := math.Max(l[0].Y, l[1].Y)
	minX := math.Min(l[0].X, l[1].X)
	minY := math.Min(l[0].Y, l[1].Y)
	return BoundingBox{Min: Point{X: minX, Y: minY}, Max: Point{X: maxX, Y: maxY}}
}

func createLine(a Point, b Point) Line {
	return Line{a, b}
}

//The Euclidean distance between two points of the plane with Cartesian coordinates
//Calcuate Hypoteneuse between two Points.
func (l *Line) getLineLength() float64 {
	return math.Sqrt(math.Pow((l[1].X-l[0].X), 2) + math.Pow((l[1].Y-l[0].Y), 2))
}

func (l *Line) getCentroid() Point {
	x := math.Sqrt(math.Pow((l[1].X-l[0].X), 2)) / 2
	y := math.Sqrt(math.Pow((l[1].Y-l[0].Y), 2)) / 2
	return Point{X: x, Y: y}
}

// //Centroid
//Distance to centorid
//Lowest Range to line.
