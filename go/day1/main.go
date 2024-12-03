package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example")
	part1("input")
	part2("example")
	part2("input")
}

func part1(filename string) {
	lines := util.GetLines(filename)
	left := &util.IntHeap{}
	right := &util.IntHeap{}

	for _, line := range lines {
		vals := strings.Split(line, "   ")
		if len(vals) != 2 {
			panic(vals)
		}

		lNum, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		rNum, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		*left = append(*left, lNum)
		*right = append(*right, rNum)
	}

	heap.Init(left)
	heap.Init(right)

	diff := 0
	for left.Len() != 0 {
		l := heap.Pop(left).(int)
		r := heap.Pop(right).(int)
		diff += int(math.Abs(float64(l) - float64(r)))
	}

	fmt.Println(diff)
}

func part2(filename string) {
	lines := util.GetLines(filename)

	left := make([]int, 0, len(lines))
	right := make(map[int]int)
	for _, line := range lines {
		vals := strings.Split(line, "   ")
		l, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		if _, exits := right[r]; !exits {
			right[r] = 0
		}
		right[r] = right[r] + 1
	}

	similarity := 0
	for _, l := range left {
		if count, exists := right[l]; exists {
			similarity += count * l
		}
	}

	fmt.Println(similarity)
}
