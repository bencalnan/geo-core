
package utilities

import "math"

import "../../main"

// FindBearing - Returns bearing between a start and end point
// Input and Output in Radians.
func FindBearing(startPoint, endPoint Point) float64 {
	var y = math.Sin(endPoint.X-startPoint.X) * math.Cos(endPoint.Y)
	var x = math.Cos(startPoint.Y)*math.Sin(endPoint.Y) - math.Sin(startPoint.Y)*math.Cos(endPoint.Y)*math.Cos(endPoint.X-startPoint.X)
	return math.Atan2(y, x)
}
