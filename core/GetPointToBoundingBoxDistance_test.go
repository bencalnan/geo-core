package core

import (
	"fmt"
	"testing"
)

func TestPointToBoundingBoxGeographic(t *testing.T) {

	testBBox := BoundingBox{
		Min: Point{X: 1.070678, Y: 51.279336},
		Max: Point{X: 1.071720, Y: 51.277559},
	}

	// //Point which should return a distnace to a point on the line.
	startPoint := LatLng{Lat: 51.27936, Lng: 1.07263}
	distance := geoBoxDist(startPoint, testBBox)
	fmt.Printf("%v meters \n", distance)

	// if int64(distance) != 77 {
	// 	t.Errorf("Distance between two points is wrong, got: %d meters, want: %d meters.", int64(distance), 77)
	// }

}
