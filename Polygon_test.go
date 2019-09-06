package main

import (
	"fmt"
	"testing"
)

func TestCreatePolygonFromPoints(t *testing.T) {
	poly := CreatePolygonFromPoints(
		[]Point{
			Point{0, 0},
			Point{3, 0},
			Point{3, 3},
			Point{0, 3},
		},
	)
	edges := poly.GetNumEdges()

	if edges != 4 {
		t.Errorf("Creating Polygon from points function was incorrect, got: %d edges, want: %d edges.", edges, 4)
	}

	closed := poly.ClosedChain()

	if closed != true {
		t.Error("Creating Polygon from points failed, polygon not closed")
	}
}

func TestGetVertices(t *testing.T) {
	poly := CreatePolygonFromPoints(
		[]Point{
			Point{0, 0},
			Point{3, 0},
			Point{3, 3},
			Point{0, 3},
		},
	)
	vertices := poly.GetVertices()
	fmt.Printf("%v\n", vertices)
	if len(vertices) != 4 {
		t.Errorf("Wrong Number of vertices, got %d, want %d", len(vertices), 4)
	}
}

func TestGetPerimeter(t *testing.T) {
	poly := CreatePolygonFromPoints(
		[]Point{
			Point{0, 0},
			Point{3, 0},
			Point{3, 3},
			Point{0, 3},
		},
	)
	total := poly.GetPerimeter()
	if total != 12 {
		t.Errorf("Perimeter calculation was incorrect, got: %d, want: %d.", int(total), 12)
	}
}

func TestGetArea(t *testing.T) {
	poly := CreatePolygonFromPoints(
		[]Point{
			Point{0, 0},
			Point{3, 0},
			Point{3, 3},
			Point{0, 3},
		},
	)
	total := poly.GetArea()
	if total != 9 {
		t.Errorf("Area calculation was incorrect, got: %d, want: %d.", int(total), 9)
	}
}

// func TestGetCentroid(t *testing.T) {
// 	poly := CreatePolygonFromPoints(
// 		[]Point{
// 			Point{0, 0},
// 			Point{4, 0},
// 			Point{4, 4},
// 			Point{0, 4},
// 		},
// 	)

// }
