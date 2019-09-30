package core

import (
	"fmt"
	"testing"
)

func TestPointToPointDistanceCosine(t *testing.T) {

	startLat := 51.27936
	startLng := 1.07263
	endLat := 51.27741
	endLng := 1.07281

	startPoint := Point{X: DegToRad(startLng), Y: DegToRad(startLat)}
	endPoint := Point{X: DegToRad(endLng), Y: DegToRad(endLat)}

	distance := PointToPointDistanceCosine(startPoint, endPoint)

	fmt.Printf("%v meters \n", distance)

	if int64(distance) != 217 {
		t.Errorf("Distace between two points is wrong, got: %d meters, want: %d meters.", int64(distance), 217)
	}

}

func TestPointToPointDistanceHaversine(t *testing.T) {

	startLat := 51.27936
	startLng := 1.07263
	endLat := 51.27741
	endLng := 1.07281

	startLatLng := LatLng{Lat: startLat, Lng: startLng}
	endLatLng := LatLng{Lat: endLat, Lng: endLng}

	distance := PointToPointHaversine(startLatLng, endLatLng)

	fmt.Printf("%v meters \n", distance)

	if int64(distance) != 217 {
		t.Errorf("Distace between two points is wrong, got: %d meters, want: %d meters.", int64(distance), 217)
	}

}
