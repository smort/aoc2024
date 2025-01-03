package main

import (
	"fmt"
	"math"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example", 2000)
	part1("input", 2000)
	part2("example3", 2000)
	part2("input", 2000)
}

func part1(filename string, iterations int) {
	nums := readInput(filename)
	result := 0
	for _, n := range nums {
		for range iterations {
			n = calcNextCode(n)
		}

		result += n
	}

	fmt.Println(result)
}

type bananaBid struct {
	price int
	diff  int
}

func part2(filename string, iterations int) {
	nums := readInput(filename)

	buyers := make([][]bananaBid, 0)
	for _, sc := range nums {
		prevPrice := sc % 10
		bids := make([]bananaBid, iterations)
		for idx := range iterations {
			sc = calcNextCode(sc)
			price := sc % 10
			diff := price - prevPrice
			bids[idx] = bananaBid{price: price, diff: diff}
			prevPrice = price
		}

		buyers = append(buyers, bids)
	}

	// store the sequence as well as the price for each buyer that had that sequence
	seqs := make(map[[4]int][]int)
	for buyerNum, buyerBids := range buyers {
		seq := [4]int{buyerBids[0].diff, buyerBids[1].diff, buyerBids[2].diff}
		for i := 3; i < len(buyerBids); i++ {
			curr := buyerBids[i]
			seq[3] = curr.diff

			if _, exists := seqs[seq]; !exists {
				seqs[seq] = make([]int, len(buyers))
			}

			// only want to store it the first time we see it. hopefully checking 0 is ok...
			if seqs[seq][buyerNum] == 0 {
				seqs[seq][buyerNum] = curr.price
			}

			// shift the sequence
			seq[0], seq[1], seq[2], seq[3] = seq[1], seq[2], seq[3], 0
		}
	}

	maxBananas := math.MinInt
	var maxSeq [4]int
	for seq, prices := range seqs {
		sum := 0
		for _, price := range prices {
			sum += price
		}

		if sum > maxBananas {
			maxBananas = sum
			maxSeq = seq
		}
	}

	fmt.Println(maxSeq, maxBananas)
}

func readInput(filename string) []int {
	return util.GetLinesTransformed(filename, func(s string) (int, error) {
		return util.MustConvAtoi(s), nil
	})
}

func calcNextCode(num int) int {
	num = ((64 * num) ^ num)
	num = num % 16777216
	num = int(num/32) ^ num
	num = num % 16777216
	num = (num * 2048) ^ num
	num = num % 16777216

	return num
}
