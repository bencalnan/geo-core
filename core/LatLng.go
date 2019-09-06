package core

import (
	"math"

	"github.com/bencalnan/geo-tools/utilities"
)

//LatLng - Lat Long Coordinate
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

//ConvertToXYZ - Converts Lat Long (In Radians) into ECEF Cartesian coordinates (Earth-Centered, Earth-Fixed) (Geodetic -> Geocentric)
func (l *LatLng) ConvertToXYZ() Point3D {
	p := Point3D{
		X: utilities.EarthRadius * math.Cos(l.Lat) * math.Cos(l.Lng),
		Y: utilities.EarthRadius * math.Cos(l.Lat) * math.Sin(l.Lng),
		Z: utilities.EarthRadius * math.Sin(l.Lat),
	}
	return p
}
