package core

//Polygon - Closed Chain Polyline
type Polygon []Line

// GeomType - Describes geometry type
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

//Vertices - Returns distinct vertices that make up the Polygon.
func (p Polygon) Vertices() []Point {
	var distinctPoints []Point
	// distinctPoints = append(distinctPoints, p[0][0])
	for _, l := range p {
		distinctPoints = append(distinctPoints, l[0])
	}
	return distinctPoints
}

// GetNumEdges - Returns NumEdges returns the number of edges in this shape.
// Copied from S2
func (p *Polygon) GetNumEdges() int {
	if len(*p) == 0 {
		return 0
	}
	return len(*p)
}

//Perimeter - Returns perimeter of polygon
func (p *Polygon) Perimeter() float64 {
	var d float64
	for _, l := range *p {
		d = d + l.length()
	}
	return d
}

// Area - Returns area of polygon
// https://www.mathopenref.com/coordpolygonarea.html
// Note does not work for self intersecting polygons. (need to add catch for this. )
func (p Polygon) Area() float64 {
	distinctPoints := p.Vertices()
	distinctPoints[len(distinctPoints)] = distinctPoints[0]
	var subTotal float64
	for i := 0; i < len(distinctPoints)-1; i++ {
		part := (distinctPoints[i].X * distinctPoints[i+1].Y) + (distinctPoints[i].Y * distinctPoints[i+1].X)
		subTotal = subTotal + part
	}
	return subTotal / 2
}

func (p *Polygon) bbox() BoundingBox {
	points := p.Vertices()
	points[len(points)] = points[0]

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


