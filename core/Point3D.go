package core

import (
	"math"

	"github.com/bencalnan/geo-tools/utilities"
)

//Point3D - 3 Dimensional Point
type Point3D struct {
	X float64
	Y float64
	Z float64
}

// GeomType - Describes geometry type
func (Point) GeomType() string {
	return "point3D"
}

//ConvertToLatLng - Takes a 3D Cartesian Point (ECEF) and returns Lat/Lngs in Radians. (Geocentric -> Geodetic)
func (p *Point3D) ConvertToLatLng() LatLng {
	l := LatLng{
		Lat: math.Asin(p.Z / utilities.EarthRadius),
		Lng: math.Atan2(p.Y, p.X),
	}

	return l
}
