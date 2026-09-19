# Day 05: Fresh Ingredient IDs

## Problem

Given a set of ranges representing acceptable (fresh) ingredient IDs, and a list of individual ingredient IDs.

### Part 1
Count how many of the individual ingredient IDs are fresh (fall within at least one range).

### Part 2
Count the total number of unique fresh ingredient IDs across all ranges combined.

## Solution

### Part 1
For each ingredient ID, check if it falls within any range. Break early on first match.

```go
for _, id := range ids {
    for _, r := range ranges {
        if id >= r.start && id <= r.end {
            count++
            break
        }
    }
}
```

**Answer: 735** (out of 999 ingredients)

### Part 2
Merge overlapping ranges, then sum the sizes:
1. Sort ranges by start value
2. Merge overlapping or adjacent ranges
3. For each merged range: `count += end - start + 1`

```go
sort.Slice(ranges, func(i, j int) bool {
    return ranges[i].start < ranges[j].start
})
// Merge and sum...
```

**Answer: 344306344403172**

## Input Format

Two sections separated by a blank line:
```
20362219004570-27230899748695    (ranges)
416093383751820-420589239208898
...
                                 (blank line)
545521298904543                  (individual IDs)
70149313006282
...
```

- 185 ranges (some overlapping, some single-value where start = end)
- 999 individual ingredient IDs
- Values up to ~560 trillion (requires int64)

## Complexity

### Part 1
- **Time**: O(n × m) where n = 999 IDs, m = 185 ranges
- **Space**: O(1)

### Part 2
- **Time**: O(m log m) for sorting ranges, O(m) for merging
- **Space**: O(1) beyond the sorted range list
- Cannot iterate individual IDs (up to 560 trillion) - must compute mathematically

## Key Insights

- Ranges can overlap, so naive counting would double-count
- Merging sorted intervals is a classic technique for computing coverage
- Adjacent ranges (end + 1 = next start) should also be merged
- int64 is required - these numbers exceed int32 range
- Part 2 ignores the individual IDs entirely
