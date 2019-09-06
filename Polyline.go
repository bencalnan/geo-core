package main

//PolyLine - Aka Polygonal chain, linestring,
type PolyLine []Line

func (PolyLine) geomType() string {
	return "polyline"
}

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

//GetLength - Returns length of a polyline.
func (p *PolyLine) length() float64 {
	var d float64
	for _, l := range *p {
		d = d + l.length()
	}
	return d
}

// GetVertices - returns all vertices in Polyline.
func (p *PolyLine) vertices() []Point {
	var v []Point
	for i, l := range *p {
		if i == 0 {
			v = append(v, Point{X: l[0].X, Y: l[0].Y}) // Add first point as well on first line.
		}
		v = append(v, Point{X: l[1].X, Y: l[1].Y}) // Only add 2nd Point normally so don't duplicate.
	}
	return v
}

func (p *PolyLine) getNumVertices() int {
	return len(*p) + 1
}

func (p *PolyLine) bbox() BoundingBox {
	points := p.vertices()
	var minX float64
	var minY float64
	var maxX float64
	var maxY float64

	for _, pt := range points {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}

	}
	return BoundingBox{Point{X: minX, Y: minY}, Point{X: maxX, Y: maxY}}
}

// NumEdges returns the number of edges in this shape. // Copied from S2 //Move out, and create interface.
func (p *PolyLine) getNumEdges() int {
	if len(*p) == 0 {
		return 0
	}
	return len(*p) - 1
}

//ClosedChain - Check if is a closed chain of lines (i.e. it is a Polygon)
func (p PolyLine) checkClosedChain() bool {
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

// func (p *PolyLine) centroid() Point {
// 	//http://www.ae.msstate.edu/vlsm/shape/centroid_of_a_line/straight.htm

// }

// - Todo

// - Is valid ? i.e. contains 1 less edge than number of points.
// - Is it a simple polygonal chain.
// - Self intersecting?
// ///Centroid
