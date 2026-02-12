# Day 03: Battery Joltage Optimization

## Problem

Each line represents a battery bank where each digit is a battery's joltage (0-9). Find the maximum joltage by selecting batteries and concatenating their digits in the order they appear.

### Part 1
Turn on exactly 2 batteries to maximize the joltage output.

Example: `987654321111111`
- Best pair: batteries at positions 0 and 1 (digits 9 and 8)
- Joltage: 98

### Part 2
Turn on exactly 12 batteries to maximize the joltage output.

## Solution

### Part 1
Check all pairs (i, j) where i < j and find the maximum:
```go
for i := 0; i < len(digits); i++ {
    for j := i + 1; j < len(digits); j++ {
        joltage := digits[i]*10 + digits[j]
        maxJoltage = max(maxJoltage, joltage)
    }
}
```
**Answer: 17158**

### Part 2
Use a greedy algorithm to select 12 batteries:
- At each step, pick the highest digit available
- Ensure enough positions remain for the remaining batteries needed
- Always maintain left-to-right order

```go
for i := 0; i < 12; i++ {
    remaining := 12 - i
    // Search from lastPos+1 to n-remaining
    // Pick position with maximum digit
}
```
**Answer: 170449335646486**

## Pitfalls Encountered

1. **Adjacent pairs assumption**: Initially thought we could only use adjacent batteries. Wrong! Can use ANY two batteries in order.

2. **Unique digits confusion**: Tried requiring the two digits to be distinct values (e.g., 9 and 8, not 9 and 9). This was incorrect.

3. **Order interpretation**: Went through several interpretations:
   - Descending order (greatest first) ✗
   - Smallest first ✗
   - Finally: order of appearance in string ✓

4. **Part 2 complexity**: Initially unclear how to pick 12 batteries optimally. Greedy algorithm works because we want to maximize the leftmost digits first.

## Key Insights

- For maximizing concatenated numbers, leftmost positions matter most
- Greedy algorithms work when optimal substructure exists (maximize digit by digit)
- O(n²) for Part 1 is acceptable for ~100 character strings
- The greedy approach in Part 2 is optimal: always pick the best available digit while leaving room for remaining selections
