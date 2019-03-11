package main

//PolyLine - Aka Polygonal chain, linestring,
type PolyLine []Line

// Creates a Polyline from a slice of Points
func createPolylineFromPoints(points []Point) PolyLine {
	var p PolyLine
	for i, pt := range points {
		if i > 0 {
			line := createLine(points[i-1], pt)
			p = append(p, line)
		}
	}
	return p
}

// Creates a Polyline from a slice of lines
func createPolyLineFromLines(lines []Line) PolyLine {
	var p PolyLine
	for _, l := range lines {
		p = append(p, l)
	}
	return p
}

//Returns length of a polyline.
func (p *PolyLine) getPolyLineLength() float64 {
	var d float64
	for _, l := range *p {
		d = d + l.getLineLength()
	}
	return d
}

// Returns all vertices in Polyline.
func (p *PolyLine) getVertices() []Point {
	var v []Point
	for i, l := range *p {
		if i == 0 {
			v = append(v, Point{X: l[0].X, Y: l[0].Y}) // Add first point as well on first line.
		}
		v = append(v, Point{X: l[1].X, Y: l[1].Y}) // Only add 2nd Point normally so don't duplicate.
	}
	return v
}

// NumEdges returns the number of edges in this shape. // Copied from S2
func (p *PolyLine) getNumEdges() int {
	if len(*p) == 0 {
		return 0
	}
	return len(*p) - 1
}

//ClosedChain - Check if is a closed chain of lines (i.e. it is a Polygon)
func (p PolyLine) ClosedChain() bool {
	start := p[0][0]
	end := p[len(p)-1][1]
	x, y := false, false
	if start.X == end.X {
		x = true
	}
	if start.Y == end.Y {
		y = true
	}
	if x == true && y == true {
		return true
	}
	return false

}

// - Todo

// - Is valid ? i.e. contains 1 less edge than number of points.
// - Is it a simple polygonal chain.
// - Self intersecting?
