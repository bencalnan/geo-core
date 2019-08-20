package main

import "math"

//Radius of earth in meters
const earthRadius = 6371000

// DegToRad - Convert degrees to radians
func DegToRad(angle float64) float64 {
	return angle * math.Pi / 180
}

// RadToDegree - Convert radians to degrees
func RadToDegree(angle float64) float64 {
	return angle * 180 / math.Pi
}

// FindBearing - Returns bearing between a start and end point
// Input and Output in Radians.
func FindBearing(startPoint, endPoint Point) float64 {
	var y = math.Sin(endPoint.X-startPoint.X) * math.Cos(endPoint.Y)
	var x = math.Cos(startPoint.Y)*math.Sin(endPoint.Y) - math.Sin(startPoint.Y)*math.Cos(endPoint.Y)*math.Cos(endPoint.X-startPoint.X)
	return math.Atan2(y, x)
}

//Normalise bearing
// var normalisedACBearing = 360 - math.Mod((radToDegree(bearingAC)+360), 360)
//findCrossTrackDistance - Cross-track distance: Returns min distance between a line and another point (Lat longs).
// This is the distance between a point and great arc, not a segment
//Input and Output in Radians
func findCrossTrackDistance(distanceAC, bearingAC, bearingAB float64) float64 {
	var ctd = math.Asin(math.Sin(distanceAC/6371000)*math.Sin(bearingAC-bearingAB)) * 6371000
	return ctd
}

//pointToPointDistanceCosine Get distance (in meters) bewteen two Lat/Longs using cosine formula. Has some benefits over using Haversine formula.
//Input and Output in Radians
func pointToPointDistanceCosine(startPoint, endPoint Point) float64 {
	var d = math.Acos(math.Sin(startPoint.Y)*math.Sin(endPoint.Y)+math.Cos(startPoint.Y)*math.Cos(endPoint.Y)*math.Cos(startPoint.X-endPoint.X)) * earthRadius
	return d
}

func pointToPointHaversine(startLat, startLon, endLat, endLon float64) float64 {

	var startLatr = DegToRad(startLat)
	var endLatR = DegToRad(endLat)
	var difLat = DegToRad(endLat - startLat)
	var difLon = DegToRad(endLon - startLon)

	var a = math.Sin(difLat/2)*math.Sin(difLat/2) + math.Cos(startLatr)*math.Cos(endLatR)*math.Sin(difLon/2)*math.Sin(difLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = earthRadius * c
	return d
}
