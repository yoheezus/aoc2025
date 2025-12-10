package main

import (
	"aoc2025/helpers"
	"fmt"
	"slices"
	"testing"
)

func TestInput(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day5-t.txt")
	ranges, rawIds := parseInput(raw_input)
	productRanges := []productRange{}
	ids := convertIds(rawIds)
	want := 3
	var result int

	for _, v := range ranges {
		productRanges = append(productRanges, NewProductRange(string(v)))
	}

	seenFreshIds := []int{}

	for _, id := range ids {
		for _, pr := range productRanges {
			if isFresh(pr, id) && !slices.Contains(seenFreshIds, id) {
				result += 1
				seenFreshIds = append(seenFreshIds, id)
			}
		}
	}

	if want != result {
		t.Errorf("Wanted: %d, Got: %d", want, result)
	}

}

func TestInputReal(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day5.txt")
	ranges, rawIds := parseInput(raw_input)
	productRanges := []productRange{}
	ids := convertIds(rawIds)
	want := 607
	var result int

	for _, v := range ranges {
		productRanges = append(productRanges, NewProductRange(string(v)))
	}

	seenFreshIds := []int{}

	for _, id := range ids {
		for _, pr := range productRanges {
			if isFresh(pr, id) && !slices.Contains(seenFreshIds, id) {
				result += 1
				seenFreshIds = append(seenFreshIds, id)
			}
		}
	}

	if want != result {
		t.Errorf("Wanted: %d, Got: %d", want, result)
	} else {
		fmt.Println(result)
	}

}

func TestInputPart2(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day5-t.txt")
	ranges, _ := parseInput(raw_input)
	productRanges := []productRange{}
	want := 14
	var result int

	for _, v := range ranges {
		productRanges = append(productRanges, NewProductRange(string(v)))
	}

	seenFreshIdsMap := make(map[int]struct{})

	for _, pr := range productRanges {
		isFresh2(pr, &seenFreshIdsMap)
	}

	result = len(seenFreshIdsMap)

	if want != result {
		t.Errorf("Wanted: %d, Got: %d", want, result)
	} else {
		fmt.Println("Result:", result)
	}

}

func TestInputRealPart2(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day5.txt")
	ranges, _ := parseInput(raw_input)
	productRanges := []productRange{}
	want := 14
	var result int

	for _, v := range ranges {
		productRanges = append(productRanges, NewProductRange(string(v)))
	}

	// seenFreshIds := []int{}
	seenFreshIdsMap := make(map[int]struct{})

	for _, pr := range productRanges {
		// seenFreshIds = append(seenFreshIds, isFresh2(pr)...)
		isFresh2(pr, &seenFreshIdsMap)
	}

	for range seenFreshIdsMap {
		result += 1
	}

	if want != result {
		t.Errorf("Wanted: %d, Got: %d", want, result)
	} else {
		fmt.Println("Result:", result)
	}

}

func TestMergeRangesDummy(t *testing.T) {
	productRanges := []productRange{productRange{start: 4, end: 15}, productRange{start: 5, end: 20}, productRange{start: 1, end: 3}}
	want := 20

	slices.SortFunc(productRanges, func(i, j productRange) int {
		return i.start - j.start
	})

	r := mergeRanges(productRanges)
	result := 0
	for _, v := range r {
		result += (v.end - v.start + 1)
	}

	if want != result {
		t.Errorf("Wanted: %v, Got: %v", want, result)
	}
}

func TestMergeRangesTest(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day5-t.txt")
	ranges, _ := parseInput(raw_input)
	productRanges := []productRange{}
	want := 14
	var result int

	for _, v := range ranges {
		productRanges = append(productRanges, NewProductRange(string(v)))
	}

	slices.SortFunc(productRanges, func(i, j productRange) int {
		return i.start - j.start
	})

	r := mergeRanges(productRanges)

	for _, v := range r {
		result += (v.end - v.start + 1)
	}

	if want != result {
		t.Errorf("Wanted: %v, Got: %v", want, result)
	} else {
		fmt.Println(result)
	}
}

func TestMergeRangesReal(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day5.txt")
	ranges, _ := parseInput(raw_input)
	productRanges := []productRange{}
	want := 14
	var result int

	for _, v := range ranges {
		productRanges = append(productRanges, NewProductRange(string(v)))
	}

	// Sort product ranges with lowest start first
	slices.SortFunc(productRanges, func(i, j productRange) int {
		return i.start - j.start
	})

	r := mergeRanges(productRanges)
	startingLen := len(r)
	for {
		// Re-sort to re-run and merge any more that needed merging
		slices.SortFunc(productRanges, func(i, j productRange) int {
			return i.start - j.start
		})
		r = mergeRanges(r)
		if startingLen == len(r) {
			break
		} else {
			startingLen = len(r)
		}

	}

	for _, v := range r {
		result += (v.end - v.start + 1)
	}

	if want != result {
		t.Errorf("Wanted: %v, Got: %v", want, result)
	}
}
