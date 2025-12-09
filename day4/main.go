package main

import (
	"aoc2025/helpers"
	"fmt"
	"strings"
)

type Grid struct {
	grid          [][]string
	idx           int
	accessedRolls [][]int
}

func NewGrid(input string) Grid {
	grid := [][]string{}
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		grid = append(grid, []string{})
		for _, c := range line {
			grid[i] = append(grid[i], string(c))
		}

	}

	return Grid{grid: grid}
}

func (g *Grid) findRolls() int {
	// check i-1[j-1] i-1[j], i-1[j+1]
	//       i[j-1]   i[j],  i[j+1]
	//		 i+1[j-1] i+1[j] i+1[j+1]
	var rowDirections []int
	var columnDirections []int
	var finalCount int
	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j] != "@" {
				continue
			}
			count := 0
			if i == 0 {
				rowDirections = []int{0, 1}
			} else if i == len(g.grid)-1 {
				rowDirections = []int{-1, 0}
			} else {
				rowDirections = []int{-1, 0, 1}
			}

			if j == 0 {
				columnDirections = []int{0, 1}
			} else if j == len(g.grid[i])-1 {
				columnDirections = []int{-1, 0}
			} else {
				columnDirections = []int{-1, 0, 1}
			}

			for _, x := range rowDirections {
				for _, y := range columnDirections {
					// fmt.Print(g.grid[i+x][j+y])
					if x == 0 && y == 0 {
						continue
					}
					if g.grid[i+x][j+y] == "@" {
						count += 1
					}
					// fmt.Printf("position [%d][%d], read: %s\n", i+x, j+y, g.grid[i+x][j+y])
				}

			}
			if count < 4 {
				finalCount += 1
				g.accessedRolls = append(g.accessedRolls, []int{i, j})
			}

			count = 0
		}
	}

	return finalCount
}

func (g *Grid) UpdateGrid() {
	// Updated the board with the currently accessed rolls, set in the findRolls() func
	for _, v := range g.accessedRolls {
		g.grid[v[0]][v[1]] = "x"
	}
	// Reset the accessedRolls
	g.accessedRolls = [][]int{}
}

func main() {
	raw_input := helpers.ReadInputAll("../inputs/day4.txt")
	fmt.Println(raw_input)
	g := NewGrid(raw_input)
	total := g.findRolls()
	fmt.Println("Part 1:", total)

	// Do an initial run for Part 1
	// Check if the result of the last function run is not 0
	// Loop updateGrid and findRolls until the result is 0
	lastChanged := total
	for lastChanged != 0 {
		g.UpdateGrid()
		lastChanged = g.findRolls()
		total += lastChanged
	}
	fmt.Println("Part 2:", total)

}
