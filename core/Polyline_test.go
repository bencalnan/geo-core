package core

import (
	"fmt"
	"testing"
)

///http://www.ae.msstate.edu/vlsm/shape/centroid_of_a_line/straight.htm
// www.ae.msstate.edu/vlsm/shape/centroid_of_a_line/example_1.htm
func TestPolylineCentroid(t *testing.T) {
	polyLine := createPolylineFromPoints(
		[]Point{
			Point{0, 0},
			Point{5, 0},
			Point{5, 4},
			Point{8, 4},
		},
	)
	centroid := polyLine.centroid()
	fmt.Printf("%+v", centroid)

}
