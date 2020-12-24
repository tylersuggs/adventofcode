package main

import "fmt"

const target = 2020

func main() {

	result := twoSum(input, target)

	one := input[result[0]]
	two := input[result[1]]

	fmt.Printf("%v + %v = 2020\n", one, two)
	fmt.Printf("%v * %v = %v", one, two, one*two)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
