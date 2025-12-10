package main

import (
	"fmt"
	"strconv"
	"strings"
)

type productRange struct {
	start int
	end   int
}

// type productRanges struct {
// 	productRanges []productRange
// 	seenFreshIds  map[int]struct{}
// }

func NewProductRange(rng string) productRange {
	s := strings.Split(rng, "-")
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])

	return productRange{start: start, end: end}
}

func convertIds(rawIds []string) []int {
	ids := []int{}
	for _, v := range rawIds {
		x, _ := strconv.Atoi(v)
		ids = append(ids, x)
	}
	return ids
}

func isFresh(pr productRange, id int) bool {
	if id >= pr.start && id <= pr.end {
		return true
	}
	return false
}

// func isFresh2(pr productRange,) []int {
// 	l := []int{}
// 	for i := pr.start; i <= pr.end; i++ {
// 		l = append(l, i)
// 	}
// 	return l
// }

func isFresh2(pr productRange, seenIds *map[int]struct{}) {
	drf := *seenIds
	for i := pr.start; i <= pr.end; i++ {
		drf[i] = struct{}{}
	}
}

func parseInput(input string) ([]string, []string) {
	splitString := strings.Split(input, "\n\n")
	ranges, ids := strings.Split(splitString[0], "\n"), strings.Split(splitString[1], "\n")

	return ranges, ids
}

func mergeRanges(prs []productRange) []productRange {
	// Given a list of a product ranges, that have a start and an end.
	// We want to remove any product ranges that are completely encapsulated by another
	// with a.start = 1, a.end = 15, b.start=5, b.end = 20
	// if a.start is less than b.start but a.end is less than b.end and greater than b.start, we have an overlap
	originalRanges := prs
	smallestStart := originalRanges[0].start
	biggestEnd := originalRanges[0].end
	mergedRanges := []productRange{}
	wasContig := false

	for currentPrIdx, currentPr := range originalRanges {
		if currentPrIdx == len(originalRanges)-1 {
			// Check if the last item increases the bounds
			if smallestStart >= currentPr.start {
				smallestStart = currentPr.start
			}

			if biggestEnd <= currentPr.end {
				biggestEnd = currentPr.end
			}

			if wasContig {
				mergedRanges = append(mergedRanges, productRange{start: smallestStart, end: biggestEnd})
				wasContig = false
			} else {
				mergedRanges = append(mergedRanges, currentPr)
			}

			continue
		}

		nextPr := originalRanges[currentPrIdx+1]
		if hasOverlap(currentPr, nextPr) {
			wasContig = true

			overlapSmallest, overlapBiggest := mergeRange(currentPr, nextPr)

			if smallestStart >= overlapSmallest {
				smallestStart = overlapSmallest
			}

			if biggestEnd <= overlapBiggest {
				biggestEnd = overlapBiggest
			}

		} else {
			if wasContig {
				mergedRanges = append(mergedRanges, productRange{start: smallestStart, end: biggestEnd})
				wasContig = false
			} else {
				mergedRanges = append(mergedRanges, currentPr)
			}

			smallestStart = nextPr.start
			biggestEnd = nextPr.end

		}
	}

	fmt.Println(smallestStart, biggestEnd)

	return mergedRanges
}

func mergeRange(a, b productRange) (int, int) {
	return a.start, b.end
}

func hasOverlap(a, b productRange) bool {
	return a.end >= b.start || a.end == b.end
}

func main() {

}
