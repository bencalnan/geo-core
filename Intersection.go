package main

//Intersection - For checking if two geometries intersect (cross one another)
type Intersection interface {
	Intersecting() bool
}

//two lines intersecting
//Two polygons intersecting
//Polygon and a polyline intersecting
