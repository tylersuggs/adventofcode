package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`\d`)
	for _, i := range input {
		ints := re.FindAllString(i, -1)
		num, _ := strconv.Atoi(ints[0] + ints[len(ints)-1])
		total += num
	}

	return total
}

func part2(input []string) int {
	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	total := 0

	for _, line := range input {
		firstIndex := -1
		firstNumber := ""

		lastIndex := -1
		lastNumber := ""

		for numKey, numValue := range numberMap {
			findex := strings.Index(line, numKey)
			lindex := strings.LastIndex(line, numKey)

			if findex != -1 && (firstIndex == -1 || findex < firstIndex) {
				firstIndex = findex
				firstNumber = numValue
			}

			if lindex != -1 && (lastIndex == -1 || lindex > lastIndex) {
				lastIndex = lindex
				lastNumber = numValue
			}
		}

		num, _ := strconv.Atoi(firstNumber + lastNumber)
		total += num
	}

	return total
}
