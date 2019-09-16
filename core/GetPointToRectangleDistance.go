package core

import "math"



//PointToRectangleDistance - Get distance to closet point on a rectangle
//Cartesian coordinates.
func PointToRectangleDistance(p Point, r Rectangle) float64 {
	preMaxX := math.Max(r.Min.X-p.X, 0)
	dx := math.Max(preMaxX, p.X-r.Max.X)
	preMaxY := math.Max(r.Min.Y-p.Y, 0)
	dy := math.Max(preMaxY, p.Y-r.Max.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

//https://stackoverflow.com/questions/5254838/calculating-distance-between-a-point-and-a-rectangular-box-nearest-point
