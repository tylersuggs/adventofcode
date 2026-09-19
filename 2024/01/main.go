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
	var left, right []int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			continue
		}
		l, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[1])
		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(left, right))
	fmt.Println("Part 2:", part2(left, right))
}

// Part 1: Sort both lists, pair them up, sum the absolute differences.
func part1(left, right []int) int {
	l := make([]int, len(left))
	r := make([]int, len(right))
	copy(l, left)
	copy(r, right)
	sort.Ints(l)
	sort.Ints(r)

	total := 0
	for i := range l {
		d := l[i] - r[i]
		if d < 0 {
			d = -d
		}
		total += d
	}
	return total
}

// Part 2: Similarity score — each left number * its count in the right list.
func part2(left, right []int) int {
	counts := make(map[int]int)
	for _, v := range right {
		counts[v]++
	}

	total := 0
	for _, v := range left {
		total += v * counts[v]
	}
	return total
}
