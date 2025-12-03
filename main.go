package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(path string) []string {
	var content []string
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}

type Wheel struct {
	length int
	idx    int
	wheel  []int
	clicks int
}

func (w *Wheel) NewWheel() {
	w.wheel = make([]int, w.length)
	for i := range w.length {
		w.wheel[i] = i
	}
}

func parseInstructions(instructions []string) []int {
	parsedInstructions := make([]int, len(instructions))
	for i, v := range instructions {
		dir := string(v[0])
		n, _ := strconv.Atoi(v[1:])
		if strings.HasPrefix(dir, "L") {
			parsedInstructions[i] = int(math.Copysign(float64(n), -1))
		}
		if strings.HasPrefix(dir, "R") {
			parsedInstructions[i] = int(math.Copysign(float64(n), 1))
		}

	}
	return parsedInstructions

}

func (w *Wheel) Reset() {
	w.idx = 50
}

func (w *Wheel) Spin(instruction int) int {
	started_zero := false
	if w.idx == 0 {
		started_zero = true
	}

	new_pos := w.idx
	new_pos += instruction

	// While the new position is below 0 or above 100
	// Keep subtracting the length of the list until it's a valid index
	for new_pos >= w.length || new_pos < 0 {
		// if new_pos == 0 {
		// 	started_zero = true
		// }

		// There is currently an error where it is counting passing 0 if it stats on 0

		if new_pos == 100 {
			new_pos = new_pos - w.length
			if !started_zero {
				w.clicks += 1
				started_zero = true
			}

		} else if new_pos < 0 {
			new_pos = w.length + new_pos // Adding as it's a negative number

		} else if new_pos >= w.length {
			new_pos = new_pos - w.length

		}
		if started_zero {
			started_zero = false
		} else {
			w.clicks += 1
		}

	}

	w.idx = new_pos
	return w.idx
}

func DoSpins(instructions []int, w *Wheel) []int {
	results := make([]int, len(instructions))
	for i, v := range instructions {
		results[i] = w.Spin(v)
	}
	return results
}

func main() {
	w := Wheel{
		length: 100,
		idx:    0,
	}
	w.NewWheel()

	var raw_instructions []string
	for _, v := range readInput("inputs/day1-1.txt") {
		raw_instructions = append(raw_instructions, string(v))
	}
	// instructions := parseInstructions(raw_instructions)
	// instructions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	instructions := []int{200}
	unfiltered := DoSpins(instructions, &w)
	var zero_count int
	for i := range unfiltered {
		if unfiltered[i] == 0 {
			zero_count += 1
		}
	}

	fmt.Printf("Part 1s: %d, Part 2s: %d, answer: %d idx: %d\n", zero_count, w.clicks, zero_count+w.clicks, w.idx)

}
