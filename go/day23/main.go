package main

import (
	"fmt"
	"maps"
	"slices"
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
	connections := util.GetLinesTransformed[[2]string](filename, func(s string) ([2]string, error) {
		split := strings.Split(s, "-")
		return [2]string{split[0], split[1]}, nil
	})
	adjMap := make(map[string][]string, 0)
	for _, conn := range connections {
		l1 := util.GetOrDefault(adjMap, conn[0], []string{})
		l1 = append(l1, conn[1])
		adjMap[conn[0]] = l1

		l2 := util.GetOrDefault(adjMap, conn[1], []string{})
		l2 = append(l2, conn[0])
		adjMap[conn[1]] = l2
	}

	computerSets := make(map[string][]string)
	for comp1, conns := range adjMap {
		for _, comp2 := range conns {
			if comp2 == comp1 {
				continue
			}

			union := sliceUnion(conns, adjMap[comp2])
			for _, common := range union {
				computers := []string{comp1, comp2, common}
				slices.Sort(computers)
				key := fmt.Sprintf("%s%s%s", computers[0], computers[1], computers[2])
				if _, exist := computerSets[key]; !exist {
					computerSets[key] = computers
				}
			}
		}
	}

	withTOnly := make([][]string, 0)
	for comps := range maps.Values(computerSets) {
		for _, comp := range comps {
			if strings.HasPrefix(comp, "t") {
				withTOnly = append(withTOnly, comps)
				break
			}
		}
	}

	fmt.Println(len(withTOnly))
}

func part2(filename string) {
	connections := util.GetLinesTransformed[[2]string](filename, func(s string) ([2]string, error) {
		split := strings.Split(s, "-")
		return [2]string{split[0], split[1]}, nil
	})
	adjMap := make(map[string][]string, 0)
	for _, conn := range connections {
		l1 := util.GetOrDefault(adjMap, conn[0], []string{})
		l1 = append(l1, conn[1])
		adjMap[conn[0]] = l1

		l2 := util.GetOrDefault(adjMap, conn[1], []string{})
		l2 = append(l2, conn[0])
		adjMap[conn[1]] = l2
	}

	var maxLAN []string

	for comp1, conns := range adjMap {
		for _, comp2 := range conns {
			network := make([]string, 0)
			union := sliceUnion(conns, adjMap[comp2])
			for _, common := range union {
				neighbors := adjMap[common]
				hasAll := true
				for _, comp := range union {
					if common == comp {
						continue
					}
					if !slices.Contains(neighbors, comp) {
						hasAll = false
						break
					}
				}
				if hasAll {
					network = append(network, common)
				}
			}

			if len(network)+2 > len(maxLAN) { // +2 for the computers we unioned at first
				maxLAN = append(union, comp1, comp2)
			}
		}
	}

	slices.Sort(maxLAN)
	fmt.Println(strings.Join(maxLAN, ","))
}

func sliceUnion[T comparable](s1 []T, s2 []T) []T {
	result := make([]T, 0)
	for _, v1 := range s1 {
		if slices.Contains(s2, v1) {
			result = append(result, v1)
		}
	}

	return result
}
