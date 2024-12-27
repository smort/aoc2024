package util

import "math"

type Grid struct {
	LenX int
	LenY int
}

func ManhattanDistance(p0, p1 Point) int {
	return int(math.Abs(float64(p0.Y-p1.Y)) + math.Abs(float64(p0.X-p1.X)))
}
