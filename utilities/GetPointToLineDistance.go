
package utilities

import (
	"math"

	"github.com/bencalnan/geo-tools/core"
)

// PointToLineCartesianDistance - Get the min distance from a point to a line (Cartesian Coordinates)
// https://en.m.wikipedia.org/wiki/Distance_from_a_point_to_a_line
func PointToLineCartesianDistance(p core.Point, l core.Line) float64 {
	var top = math.Abs(((l[1].Y - l[0].Y) * p.X) - ((l[1].X - l[0].X) * p.Y) + (l[1].X * l[0].Y) - (l[1].Y * l[0].X))
	var bottom = math.Sqrt(((l[1].Y - l[1].Y) * 2) + (l[1].X - l[0].X*2))
	var result = top / bottom
	return result
}


