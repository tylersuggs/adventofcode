# Day 05: Print Queue

## Problem

Validate and fix the ordering of pages in print updates based on ordering rules.

### Part 1
Given rules `X|Y` (page X must come before page Y), find updates already in correct order and sum their middle page numbers.

### Part 2
For incorrectly-ordered updates, sort them according to the rules and sum their middle page numbers.

## Solution

### Part 1
For each update, check all pairs to ensure no rule is violated. Sum middle elements of valid updates.

**Answer: 5964**

### Part 2
For invalid updates, sort using the rules as a comparator and sum middle elements.

```go
sort.Slice(sorted, func(i, j int) bool {
    return rules[[2]int{sorted[i], sorted[j]}]
})
```

**Answer: 4719**

## Complexity

- **Part 1**: O(u * p^2) where u = updates, p = pages per update
- **Part 2**: O(u * p log p) for sorting invalid updates
