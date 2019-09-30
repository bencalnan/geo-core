package core

import (
	"fmt"
	"testing"
)

func TestPointToPointDistanceCosine(t *testing.T) {

	startPoint := Point{X: 1.07263, Y: 51.27936}
	endPoint := Point{X: 1.07281, Y: 51.27741}

	distance := PointToPointDistanceCosine(startPoint, endPoint)
	converted := RadToDeg(distance)
	fmt.Printf("%v\n", converted)

	if distance != 4 {
		t.Errorf("Distace between two points is wrong, got: %g edges, want: %g edges.", distance, 217.63)
	}

}
