# Day 07: Tachyon Beam Splitting

## Problem

A tachyon beam is emitted downward from position `S` at the top of a triangular grid. When the beam hits a splitter (`^`), it stops and two new beams form on either side (left and right). If two beams occupy the same position, they may be treated differently depending on the interpretation.

### Part 1
Count the total number of times a beam hits a splitter (number of splits).

### Part 2
Apply the many-worlds interpretation: beams at the same position are NOT merged but represent separate timelines. Each split doubles the timelines at that position. Report the total number of active timelines at the bottom.

## Solution

### Part 1
Track beam positions as a set. For each row, check if any beam hits a `^`. Count splits and replace the beam with two new beams at adjacent positions. Merge duplicates.

**Answer: 1640**

### Part 2
Track beam COUNTS at each position using a map. When N timelines at a position hit a splitter, N timelines go left and N go right. No merging - counts accumulate.

```go
// Instead of map[int]bool, use map[int]int64
if grid[r][col] == '^' {
    newBeams[col-1] += count  // N timelines go left
    newBeams[col+1] += count  // N timelines go right
} else {
    newBeams[col] += count    // N timelines continue
}
```

**Answer: 40999072541589**

## Grid Structure

- 141 rows, ~141 columns
- `S` at center of row 1 (position 70)
- Even rows are empty (all dots)
- Odd rows contain `^` splitters in an expanding triangular pattern
- Pattern resembles a Rule 90 cellular automaton / Sierpinski triangle

## Example

```
......S......
.............
......^......
.............
.....^.^.....
.............
....^...^....
```

- Beam starts at S (position 6)
- Hits ^ at position 6 → splits to 5 and 7 (1 split)
- Beam at 5 hits ^ at 5 → splits to 4 and 6 (1 split)
- Beam at 7 hits ^ at 7 → splits to 6 and 8 (1 split)
- Part 1: beams at 4, 6 (merged), 8 → 3 splits so far
- Part 2: beams at 4(1), 6(2), 8(1) → 4 timelines

## Complexity

### Part 1
- **Time**: O(rows × beams) where beams grows linearly
- **Space**: O(beams) for the position set

### Part 2
- **Time**: O(rows × positions) - same traversal
- **Space**: O(positions) for the count map
- Numbers grow exponentially - int64 required

## Key Insights

- Part 1 vs Part 2 is set (boolean) vs multiset (counting)
- The triangular pattern naturally creates exponential timeline growth
- Empty rows between splitter rows don't affect the simulation (beams just pass through)
- The pattern is deterministic - same input always produces same splits
