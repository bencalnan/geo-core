package core

import (
	"math"
)

// MidPointGeographicCoordinates - Get the mid point of two points (geographic)
// Input and Output in Radians.
func MidPointGeographicCoordinates(a LatLng, b LatLng) LatLng {

	Bx := math.Cos(b.Lat) * math.Cos(b.Lng-a.Lng)
	By := math.Cos(b.Lat) * math.Sin(b.Lng-a.Lng)

	lngMid := a.Lng + math.Atan2(By, math.Cos(a.Lat)+Bx)
	latMid := math.Atan2(math.Sin(a.Lat)+math.Sin(b.Lat), math.Sqrt((math.Cos(a.Lat)+Bx)*(math.Cos(a.Lat)+Bx)+By*By))
	return LatLng{Lng: lngMid, Lat: latMid}
}
