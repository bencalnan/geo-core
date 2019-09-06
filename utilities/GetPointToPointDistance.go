package utilities

import (
	"math"

	"github.com/bencalnan/geo-tools/core"
)

//PointToPointDistanceCosine Get distance (in meters) bewteen two Lat/Longs using cosine formula. Has some benefits over using Haversine formula.
//Input and Output in Radians
func PointToPointDistanceCosine(startPoint, endPoint core.Point) float64 {
	var d = math.Acos(math.Sin(startPoint.Y)*math.Sin(endPoint.Y)+math.Cos(startPoint.Y)*math.Cos(endPoint.Y)*math.Cos(startPoint.X-endPoint.X)) * EarthRadius
	return d
}

//PointToPointHaversine - Altnernative method for getting distance (in meters) bewteen two Lat/Longs using cosine formula.
func PointToPointHaversine(startLat, startLon, endLat, endLon float64) float64 {

	var startLatr = DegToRad(startLat)
	var endLatR = DegToRad(endLat)
	var difLat = DegToRad(endLat - startLat)
	var difLon = DegToRad(endLon - startLon)

	var a = math.Sin(difLat/2)*math.Sin(difLat/2) + math.Cos(startLatr)*math.Cos(endLatR)*math.Sin(difLon/2)*math.Sin(difLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = EarthRadius * c
	return d
}

//CartesianDistance - Get distance betweeen two cartesian points
func CartesianDistance(source, target core.Point) float64 {
	distance := math.Sqrt(math.Pow((target.X-source.X), 2) + math.Pow((target.Y-source.Y), 2))
	return distance
}
