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
	length  int
	idx     int
	clicks  int
	on_zero int
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

func (w *Wheel) clickAdd(x int) {
	w.clicks += x
	if w.clicks < 0 {
		w.clicks = 0
	}
}

func (w *Wheel) Spin(instruction int) int {
	var direction int
	start := w.idx
	negative := instruction < 0
	steps_left := int(math.Abs(float64(instruction)))

	if negative {
		direction = -1
	} else {
		direction = 1
	}

	for steps_left != 0 {
		w.idx += direction
		steps_left -= 1

		if w.idx == w.length {
			w.idx = 0
		} else if w.idx < 0 {
			w.idx = w.length - 1
		}
		if w.idx == 0 {
			w.clicks += 1
		}
	}
	fmt.Printf("Starting %d, instruction %d, ends at %d, clicks: %d on_zeroes:%d\n", start, instruction, w.idx, w.clicks, w.on_zero)

	if w.idx == 0 {
		w.clicks -= 1
		w.on_zero += 1
	}

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
		idx:    50,
	}

	var raw_instructions []string
	for _, v := range readInput("../inputs/day1-1.txt") {
		raw_instructions = append(raw_instructions, string(v))
	}
	instructions := parseInstructions(raw_instructions)
	// instructions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	// instructions := []int{-2}
	_ = DoSpins(instructions, &w)

	fmt.Printf("Lands: %d, Passes: %d, answer: %d idx: %d\n", w.on_zero, w.clicks, w.on_zero+w.clicks, w.idx)

}

//  Method that I could not get working but I don't know why
// func (w *Wheel) Spin(instruction int) int {
// 	start := w.idx
// 	// local_clicks := 0
// 	new_pos := start + instruction
// 	if new_pos == 100 {
// 		new_pos = 0
// 	}
// 	// started_zero := start == 0

// 	for new_pos > w.length-1 || new_pos < 0 { // While out of bounds
// 		w.clickAdd(1)

// 		if new_pos == 100 {
// 			new_pos = 0
// 		}

// 		// Over Bounds
// 		if new_pos > w.length-1 {
// 			new_pos -= w.length

// 		}
// 		// Under Bounds
// 		if new_pos < 0 {
// 			new_pos += w.length // Addition because number is negative

// 		}

// 		fmt.Printf("Starting %d, instruction %d, ends at %d, clicks: %d on_zeroes:%d\n", start, instruction, new_pos, w.clicks, w.on_zero)
// 	}

// 	// if started_zero {
// 	// 	w.clickAdd(-1)
// 	// }

// 	if new_pos == 0 {
// 		w.on_zero += 1
// 	}

// 	w.idx = new_pos
// 	// w.clicks += local_clicks
// 	return w.idx
// }
