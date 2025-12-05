package main

import (
	"aoc2025/helpers"
	"reflect"
	"testing"
)

func TestInput(t *testing.T) {
	raw_input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	pbs := []PowerBank{}
	for _, v := range raw_input {
		pb := PowerBank{
			bank: v,
			max:  9,
		}
		pbs = append(pbs, pb)
	}

	want := []int{98, 89, 78, 92}
	result := []int{}
	for _, v := range pbs {
		result = append(result, v.maxJoltage())
	}

	if !reflect.DeepEqual(want, result) {
		t.Errorf("Got: %v, wanted: %v", result, want)
	}

}

func TestInputAnswer(t *testing.T) {
	raw_input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	pbs := []PowerBank{}
	for _, v := range raw_input {
		pb := PowerBank{
			bank: v,
			max:  9,
		}
		pbs = append(pbs, pb)
	}

	want := 357
	numbers := []int{}
	for _, v := range pbs {
		numbers = append(numbers, v.maxJoltage())
	}

	result := helpers.Sum(numbers)

	if result != want {
		t.Errorf("Got: %v, wanted: %v", result, want)
	}

}
