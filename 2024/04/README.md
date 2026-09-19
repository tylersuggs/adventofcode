# Day 04: Ceres Search

## Problem

Find patterns in a word search grid.

### Part 1
Count all occurrences of "XMAS" in any direction: horizontal, vertical, diagonal, forwards, or backwards.

### Part 2
Count X-MAS patterns: two "MAS" strings crossing diagonally at their center 'A', forming an X shape. Each "MAS" can read forwards or backwards.

## Solution

### Part 1
For each cell, check all 8 directions for a 4-character match against "XMAS".

**Answer: 2642**

### Part 2
For each 'A' in the grid (not on edges), check both diagonals for "MAS" or "SAM".

**Answer: 1974**

## Complexity

- **Part 1**: O(r * c * 8 * 4) = O(r * c)
- **Part 2**: O(r * c)
