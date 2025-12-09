package main

import (
	"aoc2025/helpers"
	"testing"
)

func TestInput(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day4-t.txt")

	g := NewGrid(raw_input)

	want := 13
	result := g.findRolls()

	if want != result {
		t.Errorf("Got: %d, wanted: %d", result, want)
	}
}

func TestUpdatingBoard(t *testing.T) {
	raw_input := helpers.ReadInputAll("../inputs/day4-t.txt")

	g := NewGrid(raw_input)
	_ = g.findRolls()
	g.UpdateGrid()

	want := 12
	result := g.findRolls()

	if want != result {
		t.Errorf("Got: %d, wanted: %d", result, want)
	}
}
