package core

import (
	"math"
)

//PointToPointDistanceCosine Get distance (in meters) bewteen two Lat/Lons using cosine formula. Has some benefits over using Haversine formula.
//Input and Output in Radians
func PointToPointDistanceCosine(startPoint, endPoint Point) float64 {
	var d = math.Acos(math.Sin(startPoint.Y)*math.Sin(endPoint.Y)+math.Cos(startPoint.Y)*math.Cos(endPoint.Y)*math.Cos(startPoint.X-endPoint.X)) * EarthRadius
	return d
}

//PointToPointHaversine - Alternative method for getting distance (in meters) bewteen two Lat/Longs using cosine formula. Possibly slightly faster.
func PointToPointHaversine(start, end LatLng) float64 {

	var startRadian = start.ConvertToRadian()
	var endRadian = end.ConvertToRadian()
	var difLat = endRadian.Lat - startRadian.Lat
	var difLng = endRadian.Lng - startRadian.Lng

	var a = math.Sin(difLat/2)*math.Sin(difLat/2) + math.Cos(startRadian.Lat)*math.Cos(endRadian.Lat)*math.Sin(difLng/2)*math.Sin(difLng/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = EarthRadius * c
	return d
}

//PointToPointCartesianDistance - Get distance betweeen two cartesian points
func PointToPointCartesianDistance(source, target Point) float64 {
	distance := math.Sqrt(math.Pow((target.X-source.X), 2) + math.Pow((target.Y-source.Y), 2))
	return distance
}
