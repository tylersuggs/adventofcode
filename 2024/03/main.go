package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

// Part 1: Sum all valid mul(X,Y) instructions.
func part1(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	total := 0
	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		total += a * b
	}
	return total
}

// Part 2: Same, but do() enables and don't() disables mul instructions.
func part2(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	enabled := true
	total := 0
	for _, m := range matches {
		switch {
		case m[0] == "do()":
			enabled = true
		case m[0] == "don't()":
			enabled = false
		default:
			if enabled {
				a, _ := strconv.Atoi(m[1])
				b, _ := strconv.Atoi(m[2])
				total += a * b
			}
		}
	}
	return total
}
