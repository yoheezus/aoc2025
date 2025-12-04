package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestInput(t *testing.T) {
	raw_input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	parsedInput := parseInput(raw_input)
	unflattenedResults := [][]int{}
	for _, prdRange := range parsedInput {
		r := getInvalidIds(prdRange.start, prdRange.end)
		unflattenedResults = append(unflattenedResults, r)
	}

	results := slices.Concat(unflattenedResults...)
	expected := []int{11, 22, 99, 1010, 1188511885, 222222, 0, 446446, 38593859, 0, 0, 0}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("results: %v\nexpected:%v\n", results, expected)
	}
}

func TestSingleCaseTwoDigit(t *testing.T) {
	pair := productRange{
		start: 11,
		end:   22,
	}
	expected := []int{11, 22}
	result := getInvalidIds(pair.start, pair.end)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %d invalid product Ids, expected: %d ", len(result), len(expected))
	}
}

func TestSingleCaseSixDigit(t *testing.T) {
	pair := productRange{
		start: 446443,
		end:   446449,
	}
	expected := []int{446446}
	result := getInvalidIds(pair.start, pair.end)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %d invalid product Ids, expected: %d ", len(result), len(expected))
	}
}
