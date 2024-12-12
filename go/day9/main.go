package main

import (
	"fmt"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example")
	part1("input")
	part2("example")
	part2("input")
}

type block struct {
	ID     int
	Length int
	IsFree bool
}

func part1(filename string) {
	result := 0
	totalLen := 0
	blocks := util.GetLinesTransformed[[]block](filename, func(s string) ([]block, error) {
		blocks := make([]block, 0, len(s))
		id := 0
		for idx, char := range s {
			num := util.MustConvAtoi(string(char))
			var b block
			if idx%2 != 0 {
				b = block{
					ID:     -1,
					Length: num,
					IsFree: true,
				}
			} else {
				b = block{
					ID:     id,
					Length: num,
				}
				id++
			}
			blocks = append(blocks, b)
			totalLen += b.Length
		}
		return blocks, nil
	})[0]

	i := 0
	fileSystem := make([]int, totalLen)
	for _, block := range blocks {
		for j := block.Length; j > 0; j-- {
			fileSystem[i] = block.ID
			i++
		}
	}

	i, j := 0, len(fileSystem)-1
	for i < j {
		if fileSystem[i] != -1 {
			i++
			continue
		}
		if fileSystem[j] == -1 {
			j--
			continue
		}

		fileSystem[i] = fileSystem[j]
		fileSystem[j] = -1
	}

	for idx, val := range fileSystem {
		if val == -1 {
			break
		}
		result += idx * val
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0
	totalLen := 0
	blocks := util.GetLinesTransformed[[]block](filename, func(s string) ([]block, error) {
		blocks := make([]block, 0, len(s))
		id := 0
		for idx, char := range s {
			num := util.MustConvAtoi(string(char))
			var b block
			if idx%2 != 0 {
				b = block{
					ID:     -1,
					Length: num,
					IsFree: true,
				}
			} else {
				b = block{
					ID:     id,
					Length: num,
				}
				id++
			}
			blocks = append(blocks, b)
			totalLen += b.Length
		}
		return blocks, nil
	})[0]

	i := 0
	fileSystem := make([]int, totalLen)
	for _, block := range blocks {
		for j := block.Length; j > 0; j-- {
			fileSystem[i] = block.ID
			i++
		}
	}

	for i = len(fileSystem) - 1; i > -1; i-- {
		if fileSystem[i] == -1 {
			continue
		}

		// find file block
		length := 0
		for j := i; j > -1; j-- {
			if fileSystem[j] != fileSystem[i] {
				break

			}

			length++
		}

		// find free block that matches size of file block
		freePos := -1
		for j := 0; j < i-length; j++ {
			if fileSystem[j] != -1 {
				continue
			}

			freeLength := 0
			for k := j; k < i-length+1; k++ {
				if fileSystem[k] != -1 {
					break
				}

				freeLength++

				// check if found an appropriately sized free block
				if freeLength >= length {
					freePos = j
					break
				}
			}

			if freePos != -1 {
				break
			}
		}

		// do the move if we found a free place
		if freePos != -1 {
			for j := 0; j < length; j++ {
				fileSystem[freePos+j] = fileSystem[i-j]
				fileSystem[i-j] = -1
			}
		}

		i = i - length + 1
	}

	// calculate score
	for idx, val := range fileSystem {
		if val == -1 {
			continue
		}
		result += idx * val
	}

	fmt.Println(result)
}
