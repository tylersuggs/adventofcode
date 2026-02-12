# Day 12: Presents Under Trees

Given a set of present shapes and a list of tree regions with required present
counts, determine how many regions have enough area to fit all their presents.

## Input Format

The input has two sections. First, numbered shape definitions drawn with `#`
characters on a grid:

```
0:
##
##
##

1:
.##
.#.
##.
#..
```

Second, a list of tree regions with dimensions and required counts of each shape:

```
10x8: 1 0 2 0 0 1
5x5: 0 1 0 0 1 0
```

Each region line specifies `WxH` followed by the number of each shape type
needed.

## Part 1: Area Arithmetic

Count how many regions have enough total area to fit all their required presents.
A region fits its presents when the sum of (count × shape area) for each shape
type is less than or equal to W × H.

The six shapes have areas of 6, 7, 7, 7, 5, and 7 cells respectively.

**Approach:** Parse each shape definition and count its `#` cells to compute its
area. For each region, multiply each shape's area by the required count, sum the
results, and compare against the region's total area (W × H).

**Complexity:** O(S + R) where S is the total size of shape definitions and R is
the number of regions.

**Answer:** `510`

## Part 2

Not released — Everybody Codes 2025 Day 12 only had one part available.
