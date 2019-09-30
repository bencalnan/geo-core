package core

import (
	"math"
)

// PointToLineCartesianDistance - Get the min distance from a point to a line (Cartesian Coordinates)
// https://en.m.wikipedia.org/wiki/Distance_from_a_point_to_a_line
func PointToLineCartesianDistance(p Point, l Line) float64 {
	var top = math.Abs(((l[1].Y - l[0].Y) * p.X) - ((l[1].X - l[0].X) * p.Y) + (l[1].X * l[0].Y) - (l[1].Y * l[0].X))
	var bottom = math.Sqrt(((l[1].Y - l[1].Y) * 2) + (l[1].X - l[0].X*2))
	var result = top / bottom
	return result
}

// PointToLineGeographicDistance - Gets the min distnace from a point to a line (Geographic Cordinates)
func PointToLineGeographicDistance(p LatLng, l Line) float64 {
	var a = LatLng{
		Lat: DegToRad(l[0].Y),
		Lng: DegToRad(l[0].X),
	}
	var b = LatLng{
		Lat: DegToRad(l[1].Y),
		Lng: DegToRad(l[1].X),
	}
	var c = p.ConvertToRadian()

	var t = nearestPointGreatCircle2(a, b, c)
	aPoint := a.ConvertToPoint()
	bPoint := b.ConvertToPoint()
	cPoint := c.ConvertToPoint()
	tPoint := t.ConvertToPoint()

	//If closet point is on the line use that.
	if onSegment(aPoint, bPoint, tPoint) {
		return PointToPointDistanceCosine(tPoint, cPoint)
	}

	//Otherwise just use start or end point whichever is closer.
	var distanceAC = PointToPointDistanceCosine(aPoint, cPoint)
	var distanceBC = PointToPointDistanceCosine(bPoint, cPoint)

	if distanceAC < distanceBC {
		return distanceAC
	}
	return distanceBC

}

func nearestPointGreatCircle2(a, b, c LatLng) LatLng {
	var aCartesian = a.ConvertToXYZ()
	var bCartesian = b.ConvertToXYZ()
	var cCartesian = c.ConvertToXYZ()

	var G = vectorProduct(aCartesian, bCartesian)
	var F = vectorProduct(cCartesian, G)
	var t = vectorProduct(G, F)
	var norm = normalize(t)
	var multi = multiplyByScalar(norm, EarthRadius)
	var cart = multi.ConvertToLatLng()

	return cart
}

func vectorProduct(a, b Point3D) Point3D {
	var result = Point3D{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
	return result
}

func normalize(t Point3D) Point3D {
	var length = math.Sqrt((t.X * t.X) + (t.Y * t.Y) + (t.Z * t.Z))
	var result = Point3D{
		X: t.X / length,
		Y: t.Y / length,
		Z: t.Z / length,
	}
	return result
}

func multiplyByScalar(normalize Point3D, k float64) Point3D {
	var result = Point3D{
		X: normalize.X * k,
		Y: normalize.Y * k,
		Z: normalize.Z * k,
	}
	return result
}

//Needs Radians -- Checks if point is on the line by substracting distances to check if total lenght - two sub lenghts are near enough to be zero.
func onSegment(a, b, t Point) bool {
	var diff = math.Abs(PointToPointDistanceCosine(a, b) - PointToPointDistanceCosine(a, t) - PointToPointDistanceCosine(b, t))
	if diff < 0.1 {
		return true
	}
	return false
}
