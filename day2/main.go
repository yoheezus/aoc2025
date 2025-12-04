package main

import (
	"aoc2025/helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type productRange struct {
	start int
	end   int
}

func parseInput(raw_input string) []productRange {
	productRanges := []productRange{}
	// Split by ,
	for _, rng := range strings.Split(raw_input, ",") {
		start, _ := strconv.Atoi(strings.Split(rng, "-")[0])
		end, _ := strconv.Atoi(strings.Split(rng, "-")[1])

		productRanges = append(productRanges, productRange{start: start, end: end})
	}
	return productRanges

}

func getInvalidIds(start int, end int) []int {
	invalidIds := []int{}
	for i := start; i <= end; i++ {
		stringified := strconv.Itoa(i)
		// If the number of digits is odd, skip
		if len(stringified)%2 != 0 {
			continue
		}
		mid := len(stringified) / 2
		left, right := stringified[mid:], stringified[:mid]

		if left == right {
			invalidIds = append(invalidIds, i)
		}
	}
	if len(invalidIds) == 0 {
		invalidIds = []int{0}
	}
	return invalidIds
}

func main() {
	raw_input := helpers.ReadInput("../inputs/day2.txt")[0]
	parsedInput := parseInput(raw_input)

	unflattenedResults := [][]int{}
	for _, prdRange := range parsedInput {
		r := getInvalidIds(prdRange.start, prdRange.end)
		unflattenedResults = append(unflattenedResults, r)
	}

	results := slices.Concat(unflattenedResults...)
	total := 0
	for _, v := range results {
		total += v
	}

	fmt.Printf("Part 1 password: %d\n", total)
}
