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

## Key Insights

- String pattern matching can be done efficiently with `strings.Repeat()`
- Part 2's answer is larger because it includes all Part 1 results plus additional patterns
- The difference (13548283911) represents IDs with 3+ repetitions
