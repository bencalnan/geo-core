package core

import (
	"fmt"
	"testing"
)

func TestPointToLineGeographic(t *testing.T) {
	testLine := Line{Point{X: 1.070678, Y: 51.279336}, Point{X: 1.071720, Y: 51.277559}}

	//Point which should return a distnace to a point on the line.
	startPoint := LatLng{Lat: 51.27936, Lng: 1.07263}
	distance := PointToLineGeographicDistance(startPoint, testLine)
	fmt.Printf("%v meters \n", distance)

	if int64(distance) != 128 {
		t.Errorf("Distance between two points is wrong, got: %d meters, want: %d meters.", int64(distance), 128)
	}

	//Point which should return a distance to the line's start point.
	startPoint2 := LatLng{Lat: 51.280345, Lng: 1.070411}
	distance = PointToLineGeographicDistance(startPoint2, testLine)

	fmt.Printf("%v meters \n", distance)

	if int64(distance) != 113 {
		t.Errorf("Distance between two points is wrong, got: %d meters, want: %d meters.", int64(distance), 113)
	}

	//Point which should return a distance to the line's end point.
	startPoint3 := LatLng{Lat: 51.277410, Lng: 1.072814}
	distance = PointToLineGeographicDistance(startPoint3, testLine)

	fmt.Printf("%v meters \n", distance)

	if int64(distance) != 77 {
		t.Errorf("Distance between two points is wrong, got: %d meters, want: %d meters.", int64(distance), 77)
	}

}
