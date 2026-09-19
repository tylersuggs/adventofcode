# Day 02: Invalid Product IDs

## Problem

Given ranges of product IDs (e.g., `10327-17387`), identify and sum invalid IDs. An invalid ID is one where the digits form a repeating pattern.

### Part 1
An ID is invalid if it consists of a sequence repeated exactly twice:
- `55` = "5" repeated twice ✓
- `6464` = "64" repeated twice ✓
- `123123` = "123" repeated twice ✓

### Part 2
An ID is invalid if it consists of a sequence repeated at least twice:
- `123123` = "123" repeated 2 times ✓
- `123123123` = "123" repeated 3 times ✓
- `1111111` = "1" repeated 7 times ✓

## Solution

### Part 1
Check if the number has even length and first half equals second half:
```go
if len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:] {
    return true
}
```
**Answer: 30323879646**

### Part 2
Try all possible pattern lengths and check if repeating that pattern creates the full string:
```go
for patternLen := 1; patternLen <= n/2; patternLen++ {
    if n%patternLen == 0 {
        pattern := s[:patternLen]
        if strings.Repeat(pattern, n/patternLen) == s {
            return true
        }
    }
}
```
**Answer: 43872163557**

## Pitfalls Encountered

1. **Exact vs. at least twice**: Part 1 required exactly 2 repetitions, while Part 2 required 2 or more. The distinction is subtle but important.

2. **Pattern length iteration**: Must check all divisors of the string length, not just even lengths.

## Example Analysis

Given range: `10-130`

### Part 1 Invalid IDs
- 11 = "1" repeated 2x ✓
- 22 = "2" repeated 2x ✓
- 33 = "3" repeated 2x ✓
- ...
- 99 = "9" repeated 2x ✓
- 1212 = "12" repeated 2x ✓
- 1313 = "13" repeated 2x ✓

Sum: 11+22+33+44+55+66+77+88+99+1212+1313 = 1728

### Part 2 Additional Invalid IDs
- 111 = "1" repeated 3x ✓
- All from Part 1 ✓

## Input Format

The input is a comma-separated list of ranges:
```
10327-17387,74025-113072,79725385-79874177,...
```

Total numbers to check: Millions across ~30 ranges

## Complexity

- **Time**: O(n × m) where n is total numbers in ranges, m is average number length
- **Space**: O(1) per number checked
- Optimization: Only check numbers with 2+ digits

## Key Insights

- String pattern matching can be done efficiently with `strings.Repeat()`
- Part 2's answer is larger because it includes all Part 1 results plus additional patterns
- The difference (13548283911) represents IDs with 3+ repetitions
- Most numbers are valid - invalid IDs are relatively rare
- Pattern length must divide total length evenly
