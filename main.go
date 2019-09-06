package main

import (
	"fmt"

	"github.com/bencalnan/geo-tools/core"
)

func main() {
	point := core.Point{123, 2323}
	fmt.Println(point.GeomType())
}
