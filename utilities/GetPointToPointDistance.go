package utilities

import "math"

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
