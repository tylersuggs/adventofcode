# Day 07: Bridge Repair

## Problem

Determine which calibration equations can be made true by inserting operators between numbers. Operators are evaluated left-to-right, not by precedence.

### Part 1
Insert `+` and `*` operators between numbers to match the test value. Sum the test values of all solvable equations.

### Part 2
Additionally allow a concatenation operator (`||`) that joins digits, e.g. `12 || 345 = 12345`.

## Solution

### Part 1
Recursive search trying `+` and `*` at each position, with early termination when the running total exceeds the target.

**Answer: 2437272016585**

### Part 2
Same recursive approach but also trying concatenation at each position.

**Answer: 162987117690649**

## Complexity

- **Time**: O(n * 3^m) where n = equations, m = numbers per equation (worst case)
- Early pruning keeps actual runtime well below worst case
