package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type idRange struct {
	start, end int64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ranges []idRange
	var ids []int64
	parsingRanges := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			start, _ := strconv.ParseInt(parts[0], 10, 64)
			end, _ := strconv.ParseInt(parts[1], 10, 64)
			ranges = append(ranges, idRange{start, end})
		} else {
			id, _ := strconv.ParseInt(line, 10, 64)
			ids = append(ids, id)
		}
	}

	fmt.Println("Part 1:")
	freshCount := countFreshIngredients(ranges, ids)
	fmt.Printf("Fresh ingredients: %d\n", freshCount)

	fmt.Println("\n\nPart 2:")
	totalFresh := countTotalFreshIDs(ranges)
	fmt.Printf("Total fresh IDs: %d\n", totalFresh)
}

func countTotalFreshIDs(ranges []idRange) int64 {
	// Sort ranges by start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	// Merge overlapping/adjacent ranges and count total IDs
	var total int64
	curStart := ranges[0].start
	curEnd := ranges[0].end

	for _, r := range ranges[1:] {
		if r.start <= curEnd+1 {
			// Overlapping or adjacent - extend current range
			if r.end > curEnd {
				curEnd = r.end
			}
		} else {
			// Gap - finalize current range
			total += curEnd - curStart + 1
			curStart = r.start
			curEnd = r.end
		}
	}
	// Don't forget the last range
	total += curEnd - curStart + 1

	return total
}

func countFreshIngredients(ranges []idRange, ids []int64) int {
	count := 0

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				count++
				break
			}
		}
	}

	return count
}
