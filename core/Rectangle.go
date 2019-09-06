package core

//Rectangle - Defined by min and max points.
type Rectangle struct {
	Min Point
	Max Point
}

// GeomType - Describes geometry type
func (Rectangle) GeomType() string {
	return "rectangle"
}
