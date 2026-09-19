# Day 03: Mull It Over

## Problem

Scan corrupted memory for valid multiplication instructions.

### Part 1
Find all valid `mul(X,Y)` instructions (where X and Y are 1-3 digit numbers) and sum their products. Ignore malformed instructions with extra characters or spacing.

### Part 2
Additionally handle `do()` and `don't()` instructions. `mul` is enabled at start; `don't()` disables future `mul`s; `do()` re-enables them. Only sum enabled multiplications.

## Solution

### Part 1
Use regex `mul\((\d{1,3}),(\d{1,3})\)` to find all valid instructions.

**Answer: 171183089**

### Part 2
Extend regex to also match `do()` and `don't()`, then process matches in order tracking enabled state.

**Answer: 63866497**

## Complexity

- **Time**: O(n) where n = input length
- **Space**: O(1) beyond the regex matches
