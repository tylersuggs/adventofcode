package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// type Set struct {
// 	red, green, blue int
// }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:")
	fmt.Println(part1(lines))

	fmt.Println("\nPart 2:")
	fmt.Println(part2(lines))
}

func part1(input []string) int {
	total := 0

	var setLimits = map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, i := range input {

		// game: "Game #"
		// sets_string: "1 red, 2 blue, 3 green; ..."
		game, sets_string, _ := strings.Cut(i, ": ")

		// game: #
		game, _ = strings.CutPrefix(game, "Game ")
		gameId, _ := strconv.Atoi(game)

		possible := true

		// sets: ["1 red, 2 blue, 3 green", ...]
		sets := strings.Split(sets_string, "; ")

		for _, set := range sets {
			for _, dice := range strings.Split(set, ", ") {
				s := strings.Split(dice, " ")
				s0int, _ := strconv.Atoi(s[0])

				if s0int > setLimits[s[1]] {
					possible = false
				}
			}
		}
		if possible {
			total += gameId
		}
	}

	return total
}

func part2(input []string) int {
	total := 0
	// var parsedSets []Set

	for _, i := range input {
		max := map[string]int{}

		// game: "Game #"
		// sets_string: "1 red, 2 blue, 3 green; ..."
		_, sets_string, _ := strings.Cut(i, ": ")

		// possible := true

		// sets: ["1 red, 2 blue, 3 green", ...]
		sets := strings.Split(sets_string, "; ")

		for _, set := range sets {
			for _, dice := range strings.Split(set, ", ") {
				s := strings.Split(dice, " ")
				s0int, _ := strconv.Atoi(s[0])

				if s0int > max[s[1]] {
					max[s[1]] = s0int
				}
			}
		}
		total += max["red"] * max["green"] * max["blue"]
	}

	return total
}
