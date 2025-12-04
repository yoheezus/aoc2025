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

func remove[T comparable](l []T, item T) []T {
	out := make([]T, 0)
	for _, element := range l {
		if element != item {
			out = append(out, element)
		}
	}
	return out
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

func getInvalidIdsPart2(start int, end int) []int {
	invalidIds := []int{}
	invalidIdsSet := make(map[int]struct{})
	for i := start; i <= end; i++ {
		stringified := strconv.Itoa(i)

		// Determine the mid point, as that will result in the smallest possible
		mid := len(stringified) / 2
		chunks := []string{}
		for i := 0; i < mid; i++ {
			chunks = append(chunks, stringified[:i+1])
		}

		// fmt.Println(chunks)
		for _, chunk := range chunks {
			split := strings.Split(stringified, chunk)
			count := len(remove(split, ""))
			if count == 0 {
				invalidIdsSet[i] = struct{}{}
			}
		}

		// Scan accross the digits until the mid point for possible combinations
		// e.g. 12341234 = [1, 12, 123, 1234]
		// check how many times each of those appear, Repitions must be at least 2
	}
	for k, _ := range invalidIdsSet {
		invalidIds = append(invalidIds, k)
	}
	return invalidIds
}

func main() {
	raw_input := helpers.ReadInput("../inputs/day2.txt")[0]
	// raw_input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	parsedInput := parseInput(raw_input)

	unflattenedResults := [][]int{}
	for _, prdRange := range parsedInput {
		r := getInvalidIds(prdRange.start, prdRange.end)
		unflattenedResults = append(unflattenedResults, r)
	}

	part1 := slices.Concat(unflattenedResults...)
	part1Total := 0
	for _, v := range part1 {
		part1Total += v
	}

	unflattenedResultsPart2 := [][]int{}
	for _, prdRange := range parsedInput {
		r := getInvalidIdsPart2(prdRange.start, prdRange.end)
		unflattenedResultsPart2 = append(unflattenedResultsPart2, r)
	}

	part2 := slices.Concat(unflattenedResultsPart2...)
	part2Total := 0
	for _, v := range part2 {
		part2Total += v
	}

	fmt.Printf("Part 1 password: %d\n", part1Total)
	fmt.Printf("Part 2 password: %d\n", part2Total)
}
