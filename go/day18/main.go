package main

import (
	"fmt"
	"strings"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example", 21, 7, 7)
	part1("input", 1024, 71, 71)
	part2("example", 12, 7, 7)
	part2("input", 1024, 71, 71)
}

func part1(filename string, numBytes int, lenX, lenY int) {
	bytes := readInput(filename)
	byteSet := make(map[util.Point]struct{}, numBytes)
	for idx := range numBytes {
		byteSet[bytes[idx]] = struct{}{}
	}
	adjMap := makeAdjMap(lenX, lenY, byteSet)
	steps := djikstra(adjMap, util.Point{X: 0, Y: 0}, util.Point{X: lenX - 1, Y: lenY - 1})

	fmt.Println(steps)
}

func part2(filename string, startByte, lenX, lenY int) {
	bytes := readInput(filename)
	byteSet := make(map[util.Point]struct{}, startByte)
	for idx := range startByte {
		byteSet[bytes[idx]] = struct{}{}
	}

	for numBytes := startByte; numBytes < len(bytes); numBytes++ {
		byteSet[bytes[numBytes-1]] = struct{}{}
		adjMap := makeAdjMap(lenX, lenY, byteSet)
		steps := djikstra(adjMap, util.Point{X: 0, Y: 0}, util.Point{X: lenX - 1, Y: lenY - 1})
		if steps == -1 {
			fmt.Printf("byte idx: %d - %#v\n", numBytes-1, bytes[numBytes-1])
			break
		}
	}
}

func readInput(filename string) []util.Point {
	return util.GetLinesTransformed[util.Point](filename, func(s string) (util.Point, error) {
		parts := strings.Split(s, ",")
		return util.Point{X: util.MustConvAtoi(parts[0]), Y: util.MustConvAtoi(parts[1])}, nil
	})
}

func makeAdjMap(boundX, boundY int, walls map[util.Point]struct{}) map[util.Point][]util.Point {
	g := util.Grid{
		LenX: boundX,
		LenY: boundY,
	}

	adjMap := make(map[util.Point][]util.Point, boundX*boundY)
	for y := range boundY {
		for x := range boundX {
			p := util.Point{X: x, Y: y}
			if _, exists := walls[p]; exists {
				continue
			}

			neighbors := make([]util.Point, 0)
			for _, dir := range util.Directions {
				maybeP := p.Add(dir)
				if _, exists := walls[maybeP]; exists || !maybeP.In(g) {
					continue
				}

				neighbors = append(neighbors, maybeP)
			}
			adjMap[p] = neighbors
		}
	}

	return adjMap
}

type step struct {
	util.Point
	score int
}

func djikstra(adjMap map[util.Point][]util.Point, start, end util.Point) int {
	pq := util.NewMinHeap[step]()
	pq.Push(step{start, 0}, 0)

	visited := make(map[util.Point]struct{})
	visited[start] = struct{}{}
	for pq.Len() > 0 {
		curr := pq.Pop()

		if curr.Point == end {
			return curr.score
		}

		newScore := curr.score + 1
		for _, neighbor := range adjMap[curr.Point] {
			if _, exists := visited[neighbor]; exists {
				continue
			}

			visited[neighbor] = struct{}{}
			pq.Push(step{neighbor, newScore}, newScore)
		}
	}

	return -1
}
