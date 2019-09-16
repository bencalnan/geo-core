package core

import "math"

// DegToRad - Convert degrees to radians
func DegToRad(angle float64) float64 {
	return angle * math.Pi / 180
}

// RadToDegree - Convert radians to degrees
func RadToDegree(angle float64) float64 {
	return angle * 180 / math.Pi
}
