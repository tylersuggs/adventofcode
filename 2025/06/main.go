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
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Lines 0-3 are number rows, line 4 is operations
	var rows [4][]int64
	for i := 0; i < 4; i++ {
		fields := strings.Fields(lines[i])
		for _, f := range fields {
			n, _ := strconv.ParseInt(f, 10, 64)
			rows[i] = append(rows[i], n)
		}
	}

	ops := strings.Fields(lines[4])

	fmt.Println("Part 1:")
	var total int64
	for col := 0; col < len(ops); col++ {
		var result int64
		if ops[col] == "+" {
			result = rows[0][col] + rows[1][col] + rows[2][col] + rows[3][col]
		} else {
			result = rows[0][col] * rows[1][col] * rows[2][col] * rows[3][col]
		}
		total += result
	}
	fmt.Printf("Sum of column results: %d\n", total)

	fmt.Println("\n\nPart 2:")
	total2 := solvePart2(lines)
	fmt.Printf("Sum of column results: %d\n", total2)
}

func solvePart2(lines []string) int64 {
	opLine := lines[4]
	dataLines := lines[:4]

	// Find max line length
	maxLen := len(opLine)
	for _, line := range dataLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Helper to get character at position, or space if out of bounds
	charAt := func(line string, pos int) byte {
		if pos < len(line) {
			return line[pos]
		}
		return ' '
	}

	// Find group boundaries from operation line
	type group struct {
		start int
		op    byte
	}
	var groups []group
	for i := 0; i < len(opLine); i++ {
		if opLine[i] == '*' || opLine[i] == '+' {
			groups = append(groups, group{i, opLine[i]})
		}
	}

	var totalSum int64

	for g := 0; g < len(groups); g++ {
		start := groups[g].start
		end := maxLen
		if g+1 < len(groups) {
			end = groups[g+1].start
		}

		// Read each character column in this group, form vertical numbers
		var numbers []int64
		for col := start; col < end; col++ {
			// Read digits top to bottom from data rows
			var digits []byte
			for row := 0; row < 4; row++ {
				ch := charAt(dataLines[row], col)
				if ch >= '0' && ch <= '9' {
					digits = append(digits, ch)
				}
			}
			if len(digits) > 0 {
				n, _ := strconv.ParseInt(string(digits), 10, 64)
				numbers = append(numbers, n)
			}
		}

		// Apply operation to all numbers in the group
		if len(numbers) > 0 {
			var result int64
			if groups[g].op == '+' {
				for _, n := range numbers {
					result += n
				}
			} else {
				result = 1
				for _, n := range numbers {
					result *= n
				}
			}
			totalSum += result
		}
	}

	return totalSum
}
