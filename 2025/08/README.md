# Day 8: Circuit Assembly

## Problem

Given 1000 junction boxes with 3D spatial coordinates (X,Y,Z), connect them by Euclidean distance to form circuits.

### Part 1
Make the 1000 shortest connections (pairs of boxes sorted by straight-line distance). Some connections may be redundant (both boxes already in the same circuit). Report the product of the sizes of the three largest circuits.

**Answer: 26400** (33 x 32 x 25)

### Part 2
Continue connecting closest unconnected pairs until all boxes form a single circuit. Report the product of the X coordinates of the last two junction boxes connected.

**Answer: 8199963486**

## Solution Approach

### Part 1: Sorted Edge Processing
1. Compute all ~500K pairwise Euclidean distances
2. Sort edges by distance
3. Process the first 1000 edges using Union-Find, connecting boxes into components
4. Find three largest components, multiply their sizes

### Part 2: Kruskal's MST Completion
1. Continue processing all edges in distance order (Kruskal's algorithm)
2. Track the last edge that actually merges two different components (the 999th merge)
3. Multiply the X coordinates of that edge's endpoints

## Example

With 20 junction boxes and 10 connections:
- Shortest edge connects `(162,817,812)` and `(425,690,689)` (d=316.9)
- Some edges are redundant (endpoints already in same circuit)
- After 10 connections: circuits of size [5, 4, 2, 2, 1, 1, 1, 1, 1, 1, 1] -> 5 x 4 x 2 = **40**
- Last MST edge connects `(216,146,977)` and `(117,168,530)` -> 216 x 117 = **25272**

## Pitfalls
- **Redundant connections count**: The 1000 connections include edges where both endpoints are already in the same circuit — these still count toward the limit but don't change the component structure
- **Floating point distance**: Use `math.Sqrt` for true Euclidean distance, not squared distance (edge ordering must be correct)
- **The connection count is puzzle-specific**: The number of connections (1000) is given in the puzzle text, not derived from the input size

## Complexity
- **Time**: O(N^2 log N) for building and sorting ~N^2/2 edges, plus O(N * alpha(N)) for Union-Find operations
- **Space**: O(N^2) for storing all pairwise edges
