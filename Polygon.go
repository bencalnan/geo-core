package main

//Create Geometry Interface

//Polygon - Closed Chain Polyline
type Polygon []Line

func (Polygon) geomType() string {
	return "polygon"
}

// CreatePolygonFromPoints - Creates a Polygon from a slice of Points
func CreatePolygonFromPoints(points []Point) Polygon {
	var p Polygon
	for i, pt := range points {
		if i > 0 {
			line := createLine(points[i-1], pt)
			p = append(p, line)
		}
	}

	//Final linestring to close end and first point.
	p = append(p, createLine(points[len(points)-1], points[0]))

	return p
}

//ExtractVertices - Returns distinct vertices that make up the Polygon.
func (p Polygon) getVertices() []Point {
	var distinctPoints []Point
	distinctPoints = append(distinctPoints, p[0][0])
	for _, l := range p {
		distinctPoints = append(distinctPoints, l[1])
	}
	return distinctPoints
}

// NumEdges returns the number of edges in this shape. // Copied from S2 //Move out, and create interface.
func (p *Polygon) getNumEdges() int {
	if len(*p) == 0 {
		return 0
	}
	return len(*p)
}

//GetPerimeter - Returns perimeter of polygon
func (p *Polygon) GetPerimeter() float64 {
	var d float64
	for _, l := range *p {
		d = d + l.getLineLength()
	}
	return d
}

// GetArea - Returns area of polygon
// https://www.mathopenref.com/coordpolygonarea.html
// Note does not work for self intersecting polygons. (need to add catch for this. )
func (p Polygon) GetArea() float64 {
	distinctPoints := p.getVertices()
	var subTotal float64
	for i := 0; i < len(distinctPoints)-1; i++ {
		part := (distinctPoints[i].X * distinctPoints[i+1].Y) + (distinctPoints[i].Y * distinctPoints[i+1].X)
		subTotal = subTotal + part
	}
	return subTotal / 2
}

//ClosedChain - Check if is a closed chain of lines (i.e. it is a Polygon)
func (p Polygon) ClosedChain() bool {
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

//TODO
//Self Intersecting
//Check if clockwise or anticlockwise
