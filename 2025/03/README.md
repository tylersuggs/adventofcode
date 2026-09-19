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

## Example Walkthrough

### Part 1: `818181911112111`

All pairs and their joltages:
- (pos 0, pos 1): 8,1 → 81
- (pos 0, pos 6): 8,9 → 89
- (pos 6, pos 7): 9,1 → 91
- (pos 6, pos 11): 9,2 → **92** ✓ (maximum)
- ...and many more

**Best selection**: Position 6 (digit 9) and position 11 (digit 2) → 92

### Part 2 Greedy Example: `987654321111111` (pick 12)

```
Step | Search Range  | Best Digit | Position | Reasoning
-----|---------------|------------|----------|---------------------------
1    | 0-3          | 9          | 0        | Need 11 more, can start at 0-3
2    | 1-4          | 8          | 1        | Need 10 more, can look 1-4
3    | 2-5          | 7          | 2        | Need 9 more, can look 2-5
4    | 3-6          | 6          | 3        | Need 8 more, can look 3-6
5    | 4-7          | 5          | 4        | ...continuing pattern
...  | ...          | ...        | ...      |
12   | 11-14        | 1          | 11       | Last digit needed

Result: 987654321111 (taking first 12 digits)
```

## Why Greedy Works

The value of a multi-digit number is:
```
d₁×10^(n-1) + d₂×10^(n-2) + ... + dₙ×10^0
```

Since 10^(n-1) >> 10^(n-2), maximizing d₁ is most important, then d₂, etc. This makes the greedy approach optimal.

## Complexity

### Part 1
- **Time**: O(n²) where n is line length (~100)
- **Space**: O(n) for digit array
- Total: 200 lines × 100² = ~2,000,000 comparisons

### Part 2
- **Time**: O(12 × n) = O(n) where n is line length
- **Space**: O(1) for selected digits
- Much faster than Part 1 despite larger output!

## Input Format

200 lines of ~100 digits each:
```
4123535244222342322334342233754335452333242522124322242423331132232242422443224231234323332243364522
2633322623185223292633232342336241353472323432337144252444247533312232433423423228632337331231223633
...
```

## Key Insights

- For maximizing concatenated numbers, leftmost positions matter most
- Greedy algorithms work when optimal substructure exists (maximize digit by digit)
- O(n²) for Part 1 is acceptable for ~100 character strings
- The greedy approach in Part 2 is optimal: always pick the best available digit while leaving room for remaining selections
- Part 2 is actually faster than Part 1 algorithmically (O(n) vs O(n²))
- Can't just take the 12 largest digits - order matters! "92" > "29"
