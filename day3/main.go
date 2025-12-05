package main

import (
	"aoc2025/helpers"
	"fmt"
	"sort"
	"strconv"
)

type PowerBank struct {
	bank   string
	max    int
	maxIdx int // Points to the position of the highest number
}

func (pb PowerBank) maxJoltage() int {
	// Find the highest number and it's index
	// If the possible max is not there, decrement possible max
	// If the highest number is at the end, decrement highest number
	currentMax := 9
	maxFound := false
	bankInts := []int{}

	// Convert the string to a slice of Ints
	for _, v := range pb.bank {
		x, _ := strconv.Atoi(string(v))
		bankInts = append(bankInts, x)
	}

	// Find the highest number that is not at the end of the list
	for !maxFound {
		for i, x := range bankInts {

			if x == currentMax {
				if i == len(bankInts)-1 {
					break
				}
				pb.max, pb.maxIdx = x, i
				maxFound = true
				break
			}
		}
		currentMax -= 1
	}

	// Remove everything including and before the highest number
	trimmedSlice := bankInts[pb.maxIdx+1:]

	// Sort descending
	sort.Slice(trimmedSlice, func(i, j int) bool {
		return trimmedSlice[i] > trimmedSlice[j]
	})

	fmt.Printf("bank: %s, bank ints: %v, trimmed slice:%v\n", pb.bank, bankInts, trimmedSlice)

	joltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", pb.max, trimmedSlice[0]))
	return joltage
}

func main() {
	raw_input := helpers.ReadInput("../inputs/day3.txt")
	pbs := []PowerBank{}
	for _, v := range raw_input {
		pb := PowerBank{
			bank: v,
			max:  9,
		}
		pbs = append(pbs, pb)
	}

	numbers := []int{}
	for _, v := range pbs {
		numbers = append(numbers, v.maxJoltage())
	}

	fmt.Println("Part one:", helpers.Sum(numbers))

}
