package core

import "math"

//FindCrossTrackDistance - Cross-track distance: Returns min distance between a great arc and another point (Lat longs).
//This is the distance between a point and great arc, not a segment
//Input and Output in Radians
func FindCrossTrackDistance(distanceAC, bearingAC, bearingAB float64) float64 {
	var ctd = math.Asin(math.Sin(distanceAC/EarthRadius)*math.Sin(bearingAC-bearingAB)) * EarthRadius
	return ctd
}
