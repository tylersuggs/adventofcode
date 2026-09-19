# Day 10: Machine Configuration

Each machine has a row of indicator lights and a set of buttons. Each button
affects a specific subset of the indicators, and you need to figure out the
minimum number of button presses to reach the desired configuration.

## Input Format

Each line describes one machine:
- `[..#.#]` — target indicator pattern (`#` = on, `.` = off)
- `(0,2,3)` — a button that affects indicators 0, 2, and 3
- `{4,0,7,2,1}` — joltage targets for each counter (Part 2)

## Part 1: XOR Toggles

Each button press **toggles** (XOR) its associated indicators. Find the minimum
number of button presses to match the target pattern.

**Approach:** Brute-force enumeration of all 2^m subsets of buttons (each button
is pressed at most once since XOR is self-inverse). Early exit when the current
subset size exceeds the best known solution.

**Complexity:** O(machines * 2^m) where m is the number of buttons per machine.

**Answer:** `500`

## Part 2: Additive Joltage Counters

Each button press **adds 1** to its associated counters (no toggling). Find the
minimum total number of button presses across all buttons such that every
counter reaches its exact joltage target.

**Approach:** Model as a system of linear equations Ax = t where A is the
incidence matrix, x is the vector of press counts, and t is the target vector.
Solve via Gaussian elimination (reduced row echelon form) using **exact rational
arithmetic** to avoid floating-point errors. Free variables are enumerated
recursively with pruning (non-negativity constraints on all variables, upper
bounds derived from the current partial solution and remaining degrees of
freedom).

**Complexity:** O(n * m^2) for RREF per machine, plus enumeration of free
variable assignments (bounded by the structure of the problem).

**Answer:** `19763`
