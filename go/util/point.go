package util

var Left = Point{X: -1, Y: 0}
var Up = Point{X: 0, Y: -1}
var Right = Point{X: 1, Y: 0}
var Down = Point{X: 0, Y: 1}

var Directions = []Point{Left, Up, Right, Down}

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - p.Y}
}

func (p Point) In(g Grid) bool {
	return p.X > -1 && p.Y > -1 && p.X < g.LenX && p.Y < g.LenY
}
