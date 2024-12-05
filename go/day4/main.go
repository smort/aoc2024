package main

import (
	"fmt"
	"strings"

	"github.com/smort/aoc2024/util"
)

var xmas = []string{"X", "M", "A", "S"}

type dirFunc func(grid, int, int) (int, int, bool)

var directions = []dirFunc{Above, Below, Left, Right, AboveLeft, AboveRight, BelowLeft, BelowRight}

func main() {
	part1("example")
	part1("input")
	part2("example")
	part2("input")
}

func part1(filename string) {
	lines := util.GetLinesTransformed[[]string](filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})

	result := 0
	g := grid{lines}
	var x, y int
	rowLen := len(g.Values[y])
	for y < len(g.Values) {
		if ValueAt(g, x, y) == xmas[0] {
			for _, dir := range directions {
				search(g, x, y, 1, dir, &result)
			}
		}

		x++
		if x >= rowLen {
			x = 0
			y++
		}
	}

	fmt.Println(result)
}

func part2(filename string) {
	lines := util.GetLinesTransformed[[]string](filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})

	result := 0
	g := grid{lines}
	var x, y int
	rowLen := len(g.Values[y])
	for y < len(g.Values) {
		if ValueAt(g, x, y) == "A" {
			if doesDiagonalMatch(g, AboveLeft, BelowRight, x, y) && doesDiagonalMatch(g, AboveRight, BelowLeft, x, y) {
				result++
			}
		}

		x++
		if x >= rowLen {
			x = 0
			y++
		}
	}
	fmt.Println(result)
}

func doesDiagonalMatch(g grid, diag1 dirFunc, diag2 dirFunc, x, y int) bool {
	x1, y1, exists1 := diag1(g, x, y)
	x2, y2, exists2 := diag2(g, x, y)

	if !exists1 || !exists2 {
		return false
	}

	val1 := ValueAt(g, x1, y1)
	val2 := ValueAt(g, x2, y2)
	word := val1 + "A" + val2

	return word == "MAS" || word == "SAM"
}

func search(g grid, x, y, currLetter int, f dirFunc, count *int) {
	newX, newY, exists := f(g, x, y)
	if !exists {
		return
	}

	if ValueAt(g, newX, newY) == xmas[currLetter] {
		nextLetter := currLetter + 1
		if nextLetter >= len(xmas) {
			*count++
			return
		}

		search(g, newX, newY, nextLetter, f, count)
	}
}

type grid struct {
	Values [][]string
}

func ValueAt(g grid, x, y int) string {
	return g.Values[y][x]
}

func Above(g grid, x, y int) (int, int, bool) {
	above := y - 1
	if above < 0 || above > len(g.Values)-1 {
		return 0, 0, false
	}
	return x, above, true
}

func Below(g grid, x, y int) (int, int, bool) {
	below := y + 1
	if below < 0 || below > len(g.Values)-1 {
		return 0, 0, false
	}

	return x, below, true
}

func Right(g grid, x, y int) (int, int, bool) {
	row := g.Values[y]
	right := x + 1
	if right < 0 || right > len(row)-1 {
		return 0, 0, false
	}

	return right, y, true
}

func Left(g grid, x, y int) (int, int, bool) {
	row := g.Values[y]
	left := x - 1
	if left < 0 || left > len(row)-1 {
		return 0, 0, false
	}

	return left, y, true
}

func AboveRight(g grid, x, y int) (int, int, bool) {
	x, y, _ = Above(g, x, y)
	return Right(g, x, y)
}

func AboveLeft(g grid, x, y int) (int, int, bool) {
	x, y, _ = Above(g, x, y)
	return Left(g, x, y)
}

func BelowRight(g grid, x, y int) (int, int, bool) {
	x, y, _ = Below(g, x, y)
	return Right(g, x, y)
}

func BelowLeft(g grid, x, y int) (int, int, bool) {
	x, y, _ = Below(g, x, y)
	return Left(g, x, y)
}
