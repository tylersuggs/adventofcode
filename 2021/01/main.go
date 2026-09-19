package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var depths []int

	for scanner.Scan() {
		ln, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		depths = append(depths, ln)
	}

	fmt.Println("Part 1:")
	findIncreases(depths)

	fmt.Println("\n\nPart 2:")
	findIncreasesWithinWindow(depths)
}

func findIncreases(depths []int) {
	var increases int
	for index, depth := range depths {
		var direction string

		if index == 0 {
			direction = "none - first measurement"
			continue
		}

		if depths[index] > depths[index-1] {
			direction = "increase"
			increases++
		} else if depths[index] < depths[index-1] {
			direction = "decrease"
		} else {
			direction = "no change"
		}

		fmt.Printf("%d (%s)\n", depth, direction)
	}
	fmt.Printf("Number of Increases: %d", increases)
}

func findIncreasesWithinWindow(depths []int) {
	numDepths := len(depths)
	numWindows := numDepths / 3
	remainingDepths := numDepths % 3

	fmt.Printf("Depth Count: %d\nWindow Count: %d\nRemainder: %d", numDepths, numWindows, remainingDepths)
}
