# Day 04: Paper Roll Forklift Access

## Problem

A warehouse grid where `@` represents large rolls of paper and `.` represents empty space. Forklifts can only access a roll if there are fewer than 4 rolls in the 8 adjacent positions around it.

### Part 1
Count how many rolls can be accessed by a forklift.

### Part 2
Repeatedly remove all accessible rolls until no more can be reached. Report the total number of rolls removed.

## Solution

### Part 1
For each `@` in the grid, check all 8 neighbors. If fewer than 4 are also `@`, the roll is accessible.

```go
for dr := -1; dr <= 1; dr++ {
    for dc := -1; dc <= 1; dc++ {
        // count adjacent '@' cells
    }
}
if adjacentRolls < 4 { count++ }
```

**Answer: 1356**

### Part 2
Iteratively find and remove accessible rolls:
1. Scan entire grid for accessible rolls (fewer than 4 adjacent neighbors)
2. Remove all accessible rolls simultaneously
3. Repeat until no rolls can be removed in a pass
4. Sum all removed rolls across all rounds

As rolls are removed, previously blocked rolls may become accessible since their neighbor count decreases.

**Answer: 8713**

## Example Walkthrough

```
Round 1:        Round 2:        Round 3:
@ @ . @         . . . .         . . . .
@ @ @ @   →     . @ @ .   →     . . . .
@ @ @ @         . @ @ .         . . . .
. @ . @         . . . .         . . . .
```

- Round 1: Corner and edge rolls have < 4 neighbors, removed first
- Round 2: Inner rolls now exposed, neighbor counts drop below 4
- Round 3: Remaining rolls accessible, all removed

## Complexity

### Part 1
- **Time**: O(r × c) where r = rows, c = columns
- **Space**: O(1)

### Part 2
- **Time**: O(k × r × c) where k = number of rounds until stable
- **Space**: O(r × c) for mutable grid copy and removal list
- Input size: 137 × ~138 grid

## Key Insights

- Boundary handling: positions outside the grid are treated as empty (not rolls)
- Rolls must be removed simultaneously per round, not one at a time (changing the grid mid-scan would affect neighbor counts)
- The process is like erosion in image processing - outer layers peel away first
- Some tightly packed clusters of rolls may never be fully accessible if their interior cells always maintain 4+ neighbors
