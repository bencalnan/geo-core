package utilities
import "math"

//Normalise bearing
// var normalisedACBearing = 360 - math.Mod((radToDegree(bearingAC)+360), 360)
//findCrossTrackDistance - Cross-track distance: Returns min distance between a line and another point (Lat longs).
// This is the distance between a point and great arc, not a segment
//Input and Output in Radians
func findCrossTrackDistance(distanceAC, bearingAC, bearingAB float64) float64 {
	var ctd = math.Asin(math.Sin(distanceAC/6371000)*math.Sin(bearingAC-bearingAB)) * 6371000
	return ctd
}