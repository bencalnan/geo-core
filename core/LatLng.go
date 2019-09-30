package core

import (
	"math"
)

//LatLng - Lat Long Coordinate
type LatLng struct {
	Lat float64
	Lng float64
	Alt int //Optional Altitude
}

// CreateLatLngFromPoint - Create Lat Long struct from Point
func (l *LatLng) CreateLatLngFromPoint(p Point) {
	l.Lat = p.Y
	l.Lng = p.X
}

// ConvertToRadian - Convert degrees to radians
func (l *LatLng) ConvertToRadian() LatLng {
	n := LatLng{
		Lat: DegToRad(l.Lat),
		Lng: DegToRad(l.Lng),
	}
	return n
}

// ConvertToPoint - Converts a lat long to a 2D point
func (l *LatLng) ConvertToPoint() Point {
	p := Point{
		X: l.Lng,
		Y: l.Lat,
	}
	return p
}

// ConvertToDegrees - Convert Radians to degrees
func (l *LatLng) ConvertToDegrees() LatLng {
	n := LatLng{
		Lat: RadToDeg(l.Lat),
		Lng: RadToDeg(l.Lng),
	}
	return n
}

//ConvertToXYZ - Converts Lat Long (In Radians) into ECEF Cartesian coordinates (Earth-Centered, Earth-Fixed) (Geodetic -> Geocentric)
func (l *LatLng) ConvertToXYZ() Point3D {
	p := Point3D{
		X: EarthRadius * math.Cos(l.Lat) * math.Cos(l.Lng),
		Y: EarthRadius * math.Cos(l.Lat) * math.Sin(l.Lng),
		Z: EarthRadius * math.Sin(l.Lat),
	}
	return p
}
