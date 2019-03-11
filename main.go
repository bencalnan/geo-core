package main

import "fmt"

func main() {
	point := Point{123, 2323}
	fmt.Println(point.geomType())
}

// func TestPoint(t *testing.T) {
// 	point := Point{123, 2323}
// 	geomType := point.geomType()
// 	fmt.Println(geomType)
// 	if geomType != "point" {
// 		t.Errorf("Geometry Type = %d; want point", geomType)
// 	}
// }
