package core

//PolyLine - Aka Polygonal chain, linestring,
type PolyLine []Line

// GeomType - Describes geometry type
func (PolyLine) GeomType() string {
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
func (p *PolyLine) checkClosedChain() bool {
	pV := *p
	start := pV[0][0]
	end := pV[len(pV)-1][1]
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

//centroid - Return centroid of a polyline
func (p *PolyLine) centroid() Point {

	var xTop = 0.0
	var yTop = 0.0
	var xBottom = 0.0
	var yBottom = 0.0
	for _, l := range *p {
		centroid := l.centroid()
		length := l.length()
		xTop = xTop + centroid.X*length
		yTop = yTop + centroid.Y*length
		xBottom = xBottom + length
		yBottom = yBottom + length
	}
	xCentroid := xTop / xBottom
	yCentroid := yTop / yBottom
	return Point{X: xCentroid, Y: yCentroid}
}
//https://www.geeksforgeeks.org/given-a-set-of-line-segments-find-if-any-two-segments-intersect/




///TRY THHIS ONE??? FOR TWO POLYLINES THOUGH
// https://stackoverflow.com/questions/27745972/test-if-polylines-intersects-using-python


// // https://stackoverflow.com/questions/4845569/c-sharp-polyline-is-self-crossing
//   func clockwise(p0, p1, p2 Point) int{
//         epsilon := 1e-13

//         dx1 := p1.X - p0.X
//         dy1 := p1.Y - p0.Y
//         dx2 := p2.X - p0.X
// 		dy2 := p2.Y - p0.Y
		
//         d := dx1 * dy2 - dy1 * dx2
//         if d > epsilon {
// 			return 1
// 		}
//         if d < -epsilon {
// 			return -1
// 		} 
//         if dx1*dx2 < -epsilon) || (dy1*dy2 < -epsilon {
// 			return -1
// 		} 
//         if dx1*dx1+dy1*dy1) < (dx2*dx2+dy2*dy2)+epsilon {
// 			return 1
// 		} 
//         return 0
// 	}
	
// 	func checkIntersection(p1, p2, q1, q2 Point) bool {
//         return (clockwise(p1,p2,q1) * clockwise(p1,p2,q2) <=0) &&
//             (clockwise(q1,q2,p1) * clockwise(q1,q2,p2) <=0);
// 	}
	
// 	func (*Polyline) IsSelfIntersecting() bool
//     {
//         if size <= 5 {
// 			return false
// 		}
			
			
//         first = body.Points.ElementAt(size - 1);
//         second = body.Points.ElementAt(size - 2);
//         for (int i = 0; i < size - 3; i++)
//         {
//             if (checkIntersection(first, second, body.Points.ElementAt(i),
//                 body.Points.ElementAt(i + 1)))
//             {
//                 return true;
//             }
//         }
//         return false;
//     }