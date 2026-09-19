# Day 9: Red Tile Rectangles

## Problem

Given coordinates of red tiles that form vertices of a rectilinear polygon, find the largest axis-aligned rectangle with two red tiles at opposing corners.

### Part 1
Find the largest rectangle where any two red tiles are opposing corners. No constraints on what's inside.

**Answer: 4749838800**

### Part 2
The red tiles are connected by green tiles forming a rectilinear polygon, with the interior also green. The rectangle must now fit entirely within the polygon (only red/green tiles allowed).

**Answer: 1624057680**

## Solution Approach

### Part 1: Brute Force All Pairs
For each pair of red tiles, compute area as `(|dx|+1) * (|dy|+1)`. With 495 tiles, ~122K pairs is trivial.

### Part 2: Rectangle-in-Polygon Check
For each candidate pair, verify the rectangle fits inside the rectilinear polygon:
1. **Edge intersection test**: No polygon edge (horizontal or vertical) may cross the open interior of the rectangle
2. **Point-in-polygon test**: The center of the rectangle must be inside the polygon (ray casting against vertical edges)

If no edge crosses the interior, all interior points are on the same side of the polygon boundary, so a single point-in-polygon check suffices.

## Key Insight
For a simply-connected rectilinear polygon and an axis-aligned rectangle:
- The rectangle is inside the polygon iff no polygon edge intersects its open interior AND its center is inside the polygon
- Pruning by area (skip pairs that can't beat the current best) keeps runtime fast

## Pitfalls
- **Inclusive area**: Area counts tiles, so it's `(dx+1)*(dy+1)`, not `dx*dy`
- **Boundary handling**: Polygon edges ON the rectangle boundary are fine; only edges crossing the open interior disqualify
- **Point-in-polygon with integer coords**: Offset the test point slightly to avoid landing on polygon edges

## Complexity
- **Part 1**: O(N^2) where N = number of red tiles
- **Part 2**: O(N^2 * E) where E = number of polygon edges (~N), with pruning
