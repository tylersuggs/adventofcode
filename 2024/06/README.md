# Day 06: Guard Gallivant

## Problem

Predict a guard's patrol route through a lab map. The guard turns right 90 degrees when blocked, otherwise moves forward.

### Part 1
Count how many distinct positions the guard visits before leaving the mapped area.

### Part 2
Find how many positions you could place a single new obstruction to trap the guard in a loop.

## Solution

### Part 1
Simulate the guard's movement, tracking visited positions until they exit the grid.

**Answer: 5162**

### Part 2
For each position on the guard's original path, temporarily place an obstruction and check if the guard enters a loop (revisits the same position+direction state).

**Answer: 1909**

## Complexity

- **Part 1**: O(r * c) for simulation
- **Part 2**: O(n * r * c) where n = positions on original path
