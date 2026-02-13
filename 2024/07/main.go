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
	var equations []equation
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		fields := strings.Fields(parts[1])
		var nums []int
		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			nums = append(nums, n)
		}
		equations = append(equations, equation{target, nums})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(equations))
	fmt.Println("Part 2:", part2(equations))
}

type equation struct {
	target int
	nums   []int
}

func canSolve(target int, nums []int, idx int, current int, useConcat bool) bool {
	if current > target {
		return false
	}
	if idx == len(nums) {
		return current == target
	}
	if canSolve(target, nums, idx+1, current+nums[idx], useConcat) {
		return true
	}
	if canSolve(target, nums, idx+1, current*nums[idx], useConcat) {
		return true
	}
	if useConcat {
		concat := concat(current, nums[idx])
		if canSolve(target, nums, idx+1, concat, useConcat) {
			return true
		}
	}
	return false
}

func concat(a, b int) int {
	mul := 1
	for mul <= b {
		mul *= 10
	}
	return a*mul + b
}

// Part 1: Sum test values of equations solvable with + and *.
func part1(equations []equation) int {
	total := 0
	for _, eq := range equations {
		if canSolve(eq.target, eq.nums, 1, eq.nums[0], false) {
			total += eq.target
		}
	}
	return total
}

// Part 2: Same but also allow concatenation operator.
func part2(equations []equation) int {
	total := 0
	for _, eq := range equations {
		if canSolve(eq.target, eq.nums, 1, eq.nums[0], true) {
			total += eq.target
		}
	}
	return total
}
