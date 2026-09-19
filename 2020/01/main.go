package main

import "fmt"

const target = 2020

func main() {

	p1result := twoSum(input, target)

	p1one := input[p1result[0]]
	p1two := input[p1result[1]]

	fmt.Println("Part 1:")
	fmt.Printf("%v + %v = 2020\n", p1one, p1two)
	fmt.Printf("%v * %v = %v\n\n", p1one, p1two, p1one*p1two)

	p2result := threeSum(input, target)

	p2one := input[p2result[0]]
	p2two := input[p2result[1]]
	p2three := input[p2result[2]]

	fmt.Println("Part 2:")
	fmt.Printf("%v + %v + %v = 2020\n", p2one, p2two, p2three)
	fmt.Printf("%v * %v * %v = %v\n", p2one, p2two, p2three, p2one*p2two*p2three)

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

func threeSum(nums []int, target int) []int {
	for k, v := range nums {
		t := target - v
		ts := twoSum(input, t)

		if ts != nil {
			return []int{ts[0], ts[1], k}
		}
	}
	return nil
}
