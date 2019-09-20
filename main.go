package main

import (
	"fmt"

	c "github.com/bencalnan/geo-tools/core"
)

func main() {
	point := c.Point{123, 2323}
	fmt.Println(point.GeomType())
}
