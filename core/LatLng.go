package core

import (
	"github.com/bencalnan/geo-tools/utilities"
)

//LatLng - Reprsents WGS84 Coorindates (GeoCentric)
type LatLng struct {
	Lat float64
	Lng float64
	Alt int //Optional Altitude
}

// ConvertToRadian - Convert degrees to radians
func (l *LatLng) ConvertToRadian() LatLng {
	n := LatLng{
		Lat: utilities.DegToRad(l.Lat),
		Lng: utilities.DegToRad(l.Lng),
	}
	return n
}

// ConvertToDegrees - Convert Radians to degrees
func (l *LatLng) ConvertToDegrees() LatLng {
	n := LatLng{
		Lat: utilities.RadToDegree(l.Lat),
		Lng: utilities.RadToDegree(l.Lng),
	}
	return n
}
