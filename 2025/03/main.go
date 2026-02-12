package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	totalJoltage := calculateTotalJoltage(lines)
	fmt.Printf("Total active joltage: %d\n", totalJoltage)

	fmt.Println("\n\nPart 2:")
	totalJoltage12 := calculateTotalJoltage12(lines)
	fmt.Printf("Total active joltage (12 batteries): %d\n", totalJoltage12)
}

func calculateTotalJoltage(lines []string) int {
	total := 0

	for _, line := range lines {
		// Convert to digit array
		digits := make([]int, len(line))
		for i, ch := range line {
			digits[i] = int(ch - '0')
		}

		// Find maximum joltage from any pair (i, j) where i < j
		maxJoltage := 0
		for i := 0; i < len(digits); i++ {
			for j := i + 1; j < len(digits); j++ {
				joltage := digits[i]*10 + digits[j]
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}

		total += maxJoltage
	}

	return total
}

func calculateTotalJoltage12(lines []string) int64 {
	var total int64

	for _, line := range lines {
		maxJoltage := findMaxJoltage12Batteries(line)
		total += maxJoltage
	}

	return total
}

func findMaxJoltage12Batteries(line string) int64 {
	digits := make([]int, len(line))
	for i, ch := range line {
		digits[i] = int(ch - '0')
	}

	n := len(digits)
	numBatteries := 12

	// Greedy approach: at each step, pick the maximum digit
	// while ensuring we have enough remaining positions
	selectedDigits := make([]int, 0, numBatteries)
	lastPos := -1

	for i := 0; i < numBatteries; i++ {
		remaining := numBatteries - i
		maxDigit := -1
		maxPos := -1

		// We can look from lastPos+1 to n-remaining (inclusive)
		// to ensure we have enough positions left for remaining batteries
		for pos := lastPos + 1; pos <= n-remaining; pos++ {
			if digits[pos] > maxDigit {
				maxDigit = digits[pos]
				maxPos = pos
			}
		}

		selectedDigits = append(selectedDigits, maxDigit)
		lastPos = maxPos
	}

	// Convert selected digits to number
	var result int64
	for _, d := range selectedDigits {
		result = result*10 + int64(d)
	}

	return result
}
