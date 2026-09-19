package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reports [][]int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		var levels []int
		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			levels = append(levels, n)
		}
		reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(reports))
	fmt.Println("Part 2:", part2(reports))
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}
		if increasing && diff < 0 {
			return false
		}
		if !increasing && diff > 0 {
			return false
		}
	}
	return true
}

// Part 1: Count reports that are safe as-is.
func part1(reports [][]int) int {
	count := 0
	for _, r := range reports {
		if isSafe(r) {
			count++
		}
	}
	return count
}

// Part 2: Count reports that are safe, or become safe by removing one level.
func part2(reports [][]int) int {
	count := 0
	for _, r := range reports {
		if isSafe(r) {
			count++
			continue
		}
		for i := range r {
			dampened := make([]int, 0, len(r)-1)
			dampened = append(dampened, r[:i]...)
			dampened = append(dampened, r[i+1:]...)
			if isSafe(dampened) {
				count++
				break
			}
		}
	}
	return count
}
