# Day 01: Rotary Dial Puzzle

## Problem

A rotary dial puzzle where you track position changes on a dial numbered 0-99. The dial starts at position 50 and responds to commands like "L44" (rotate left 44 clicks) or "R35" (rotate right 35 clicks).

### Part 1
Count how many times the dial lands exactly on position 0 after completing a move.

### Part 2
Count how many times the dial points at position 0 during ANY individual click (not just at the end of moves).

## Solution

### Part 1
- Track position with modulo arithmetic: `(position ± clicks) % 100`
- Count when final position equals 0 after each move
- **Answer: 1055**

### Part 2
- Count every time position becomes 0 during movement
- For right moves: `(position + clicks) / 100` gives number of times we hit 0
- For left moves:
  - If starting at position 0: `clicks / 100` (don't count starting position)
  - Otherwise: `(clicks + 100 - position) / 100`
- **Answer: 6386**

## Pitfalls Encountered

1. **Initial wrapping bug**: Used simple `if position < 0 { position += 100 }` instead of proper modulo arithmetic. This failed for large movements like R998 that wrap multiple times.

2. **Part 2 starting position**: When the dial starts at position 0 and moves away (e.g., L5 or R14), we shouldn't count the starting position as "hitting" 0. Required special case handling.

3. **Formula complexity**: The formula for counting zero crossings during left moves required careful thought about when we actually pass through position 0 versus just starting there.

## Key Insights

- Modulo arithmetic is essential for circular movement
- Counting occurrences during movement is different from counting final positions
- Edge cases matter: starting at the target position requires special handling
