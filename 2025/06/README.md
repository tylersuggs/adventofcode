# Day 06: Column Arithmetic

## Problem

Given 4 rows of numbers and a 5th row of operations (`*` or `+`), perform calculations on grouped numbers.

### Part 1
Split each row by whitespace to get numbers. Numbers at the same position across rows form a column. Apply the operation for each column to its 4 numbers, then sum all column results.

### Part 2
Reinterpret the data as character-aligned columns. Each vertical character column (top to bottom, rows 1-4) forms a single number, with the most significant digit at the top. Groups are determined by the position of operation characters in row 5. Apply each group's operation to all vertical numbers in that group, then sum.

## Solution

### Part 1
Simple whitespace-split parsing:
```go
fields := strings.Fields(line)
// Match by position across rows
```
**Answer: 6957525317641**

### Part 2
Character-level column parsing:
```go
// For each character column in a group, read digits vertically
for col := start; col < end; col++ {
    for row := 0; row < 4; row++ {
        // collect digits top-to-bottom
    }
    // digits form a number (most significant at top)
}
```
**Answer: 13215665360076**

## Example Walkthrough (Part 2)

```
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
```

Reading character columns vertically:

| Group | Op | Column Numbers | Result |
|-------|----|---------------|--------|
| 1 (pos 0-3) | * | 1, 24, 356 | 8,544 |
| 2 (pos 4-7) | + | 369, 248, 8 | 625 |
| 3 (pos 8-11) | * | 32, 581, 175 | 3,253,600 |
| 4 (pos 12-14) | + | 623, 431, 4 | 1,058 |

Column 0 vertically: `1`, ` `, ` ` → **1**
Column 1 vertically: `2`, `4`, ` ` → **24**
Column 2 vertically: `3`, `5`, `6` → **356**

Total: 8,544 + 625 + 3,253,600 + 1,058 = **3,263,827**

## Complexity

### Part 1
- **Time**: O(n) where n is number of columns
- **Space**: O(n) for parsed numbers

### Part 2
- **Time**: O(c × r) where c = total character columns, r = 4 rows
- **Space**: O(c) for numbers per group

## Key Insights

- Part 1 and Part 2 use the same data but interpret it completely differently
- Part 2 requires character-level alignment rather than whitespace-delimited parsing
- Vertical numbers are formed by reading digits top-to-bottom in each character column
- Group boundaries are determined by operation character positions in the last row
- The scanner buffer size needs to be increased for very long input lines
