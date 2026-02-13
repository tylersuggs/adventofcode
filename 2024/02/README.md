# Day 02: Red-Nosed Reports

## Problem

Analyze safety reports from a nuclear reactor. Each report is a list of numbers (levels).

### Part 1
A report is safe if:
1. Levels are all increasing or all decreasing
2. Adjacent levels differ by at least 1 and at most 3

Count the number of safe reports.

### Part 2
A report also counts as safe if removing any single level makes it satisfy the Part 1 rules (the "Problem Dampener").

## Solution

### Part 1
Check each report for monotonicity and valid step sizes.

**Answer: 379**

### Part 2
For each unsafe report, try removing each level one at a time and re-check.

**Answer: 430**

## Complexity

- **Part 1**: O(n * m) where n = reports, m = levels per report
- **Part 2**: O(n * m^2) due to trying each removal
