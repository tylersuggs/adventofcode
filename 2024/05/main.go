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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse rules (X|Y means X must come before Y)
	rules := make(map[[2]int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		rules[[2]int{a, b}] = true
	}

	// Parse updates
	var updates [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		var pages []int
		for _, p := range parts {
			n, _ := strconv.Atoi(p)
			pages = append(pages, n)
		}
		updates = append(updates, pages)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(rules, updates))
	fmt.Println("Part 2:", part2(rules, updates))
}

func isOrdered(rules map[[2]int]bool, pages []int) bool {
	for i := 0; i < len(pages); i++ {
		for j := i + 1; j < len(pages); j++ {
			if rules[[2]int{pages[j], pages[i]}] {
				return false
			}
		}
	}
	return true
}

// Part 1: Sum middle pages of correctly-ordered updates.
func part1(rules map[[2]int]bool, updates [][]int) int {
	total := 0
	for _, pages := range updates {
		if isOrdered(rules, pages) {
			total += pages[len(pages)/2]
		}
	}
	return total
}

// Part 2: Sort incorrectly-ordered updates by rules, sum their middle pages.
func part2(rules map[[2]int]bool, updates [][]int) int {
	total := 0
	for _, pages := range updates {
		if isOrdered(rules, pages) {
			continue
		}
		sorted := make([]int, len(pages))
		copy(sorted, pages)
		sort.Slice(sorted, func(i, j int) bool {
			return rules[[2]int{sorted[i], sorted[j]}]
		})
		total += sorted[len(sorted)/2]
	}
	return total
}
