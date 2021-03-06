package core

//Point - Ordered Pair
type Point struct {
	X float64
	Y float64
}

// GeomType - Describes geometry type
func (Point) GeomType() string {
	return "point"
}

func (p *Point) createPointFromCoords(x float64, y float64) {
	p.X = x
	p.Y = y
}

func (p *Point) createPointFromLatLng(latLng LatLng) {
	p.X = latLng.Lng
	p.Y = latLng.Lat
}

//Add Two points together
func (p Point) add(n Point) Point {
	return Point{p.X + n.X, p.Y + n.Y}
}

//Subtract one point from another.
func (p Point) subtract(n Point) Point {
	return Point{p.X - n.X, p.Y - n.Y}
}
