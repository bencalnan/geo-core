package main

// // //Geometry - All Types have
// type Geometry interface {
// 	area() float64
// 	perim() float64
// }

//PolygonalChain - Polylines and Polygons
type PolygonalChain interface {
	ClosedChain() bool
}

// func getCentroid(pc PolygonalChain) Point {
// 	p := Point{X: 0, Y: 0}

// }

//Signed Area = Area inside polygon.
