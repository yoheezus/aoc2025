package main

import (
	"aoc2025/helpers"
	"fmt"
	"slices"
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

	// fmt.Printf("bank: %s, bank ints: %v, trimmed slice:%v\n", pb.bank, bankInts, trimmedSlice)

	joltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", pb.max, trimmedSlice[0]))
	return joltage
}

func (pb PowerBank) maxJoltagePart2() int {
	// Find the highest digit in the bank. Check if there are at least 12 digits after this number
	// If there is no valid digit, check for the next smallest digit and repeat
	// Use that number as the first digit. Chop slice to a subslice with atleast 12 digits
	// Continue looking for the biggest number, check if that number is 12 - 1 away
	currentMax := 9
	maxFound := false
	bankInts := []int{}
	desiredLength := 12

	// Convert the string to a slice of Ints
	for _, v := range pb.bank {
		x, _ := strconv.Atoi(string(v))
		bankInts = append(bankInts, x)
	}

	// Find the highest number that has atleast 12 items to the right of it
	for !maxFound {
		for i, x := range bankInts {

			if x == currentMax {
				if i+desiredLength > len(bankInts) {
					break
				}
				pb.max, pb.maxIdx = x, i
				maxFound = true
				desiredLength -= 1
				break
			}
		}
		currentMax -= 1
	}

	fmt.Printf("max: %d, maxIdx: %d, desiredLength: %d\n", pb.max, pb.maxIdx, desiredLength)
	trimmedSlice := bankInts[pb.maxIdx+1:]
	result := []int{pb.max}
	// Reset current max back to 9:
	currentMax = slices.Max(trimmedSlice)

	for len(result) < 12 {
		for i, x := range trimmedSlice {
			if x == currentMax && i <= len(trimmedSlice)-desiredLength {
				result = append(result, x)
				desiredLength -= 1
				if len(trimmedSlice) > 1 {
					trimmedSlice = trimmedSlice[i+1:]
				}
				// Avoids cases of the trimmed slice being 0
				if len(trimmedSlice) != 0 {
					currentMax = slices.Max(trimmedSlice)
				}
				break
			} else if i > len(trimmedSlice)-desiredLength {
				currentMax -= 1
				break
			}

		}

	}

	var s string

	for _, v := range result {
		s += strconv.Itoa(v)
	}

	joltage, _ := strconv.Atoi(s)
	fmt.Println(joltage)

	return joltage
}

func main() {
	raw_input := helpers.ReadInput("../inputs/day3.txt")
	// raw_input = []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	pbs := []PowerBank{}
	for _, v := range raw_input {
		pb := PowerBank{
			bank: v,
			max:  9,
		}
		pbs = append(pbs, pb)
	}

	numbers := []int{}
	part2 := []int{}
	for _, v := range pbs {
		numbers = append(numbers, v.maxJoltage())
		part2 = append(part2, v.maxJoltagePart2())
	}

	fmt.Println("Part one:", helpers.Sum(numbers))
	fmt.Println("Part two:", helpers.Sum((part2)))

}
