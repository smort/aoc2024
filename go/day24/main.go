package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example")
	part1("example2")
	part1("input")
	part2("input")
}

func part1(filename string) {
	result := int64(0)
	registers, wires := readInput(filename)

	numZ := 0
	for _, w := range wires {
		if strings.HasPrefix(w.output, "z") {
			numZ++
		}
	}

	z := make([]string, numZ)
	for numZ > 0 {
		for _, curr := range wires {
			if _, exists := registers[curr.output]; exists {
				continue
			}
			allIn := true
			for _, w := range curr.wires {
				if _, exist := registers[w]; !exist {
					allIn = false
					break
				}
			}

			if !allIn {
				continue
			}

			val1 := registers[curr.wires[0]]
			val2 := registers[curr.wires[1]]
			op := ops[curr.op]

			out := op(val1, val2)
			registers[curr.output] = out
			if strings.HasPrefix(curr.output, "z") {
				numZ--
				z[len(z)-1-util.MustConvAtoi(curr.output[1:len(curr.output)])] = strconv.Itoa(int(out))
			}
		}
	}

	result, _ = strconv.ParseInt(strings.Join(z, ""), 2, 64)
	fmt.Println(result)
}

func part2(filename string) {
	result := 0

	registers, _ := readInput(filename)
	builderX := strings.Builder{}
	builderY := strings.Builder{}
	for _, k := range slices.Backward(slices.Sorted(maps.Keys(registers))) {
		if strings.HasPrefix(k, "y") {
			builderY.WriteString(strconv.Itoa(int(registers[k])))
		}

		if strings.HasPrefix(k, "x") {
			fmt.Println(k)
			builderX.WriteString(strconv.Itoa(int(registers[k])))
		}
	}

	x, _ := strconv.ParseInt(builderX.String(), 2, 64)
	y, _ := strconv.ParseInt(builderY.String(), 2, 64)

	fmt.Println(x, y, x+y)
	fmt.Println(result)
}

var ops = map[string]func(uint8, uint8) uint8{
	"AND": AND,
	"OR":  OR,
	"XOR": XOR,
}

type wire struct {
	output string
	op     string
	wires  []string
}

func readInput(filename string) (map[string]uint8, map[string]wire) {
	registers := make(map[string]uint8)
	wires := make(map[string]wire, 0)
	lines := util.GetLines(filename)

	isRegister := true
	for _, line := range lines {
		if line == "" {
			isRegister = false
			continue
		}

		if isRegister {
			parts := strings.Split(line, ":")
			registers[parts[0]] = uint8(util.MustConvAtoi(strings.TrimSpace(parts[1])))
			continue
		}

		parts := strings.Split(line, " ")
		wire1 := parts[0]
		op := parts[1]
		wire2 := parts[2]
		output := parts[4]
		wires[output] = wire{
			output: output,
			op:     op,
			wires:  []string{wire1, wire2},
		}
	}

	return registers, wires
}

func XOR(n, n1 uint8) uint8 {
	return n ^ n1
}

func AND(n, n1 uint8) uint8 {
	return n & n1
}

func OR(n, n1 uint8) uint8 {
	return n | n1
}
