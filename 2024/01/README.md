# Day 01: Historian Hysteria

## Problem

Two groups of historians have created separate lists of location IDs that don't match up. Reconcile them.

### Part 1
Sort both lists, pair smallest-to-smallest, and sum the absolute differences between each pair.

### Part 2
Calculate a similarity score: for each number in the left list, multiply it by how many times it appears in the right list, then sum those products.

## Solution

### Part 1
Sort both lists independently, then iterate through pairs summing `|left[i] - right[i]|`.

**Answer: 1938424**

### Part 2
Build a frequency map of the right list, then for each left value multiply by its count in the map.

**Answer: 22014209**

## Complexity

- **Part 1**: O(n log n) for sorting
- **Part 2**: O(n) with a hash map for counts
