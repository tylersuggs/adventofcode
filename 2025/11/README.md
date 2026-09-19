# Day 11: Device Paths

A directed acyclic graph of devices is given, where each device forwards signals
to one or more downstream devices. The goal is to count valid paths through the
network.

## Input Format

Each line describes a device and its connections:

```
you: abc def
abc: def out
def: out
```

The device name appears before the colon, and the space-separated list of
downstream targets appears after.

## Part 1: Count All Paths

Count all distinct paths from `you` to `out` in the DAG.

**Approach:** Memoized DFS (dynamic programming on the DAG). For each node,
recursively sum the number of paths from each successor to `out`, caching
results to avoid redundant computation.

**Complexity:** O(V + E) where V is the number of devices and E is the number of
edges.

**Answer:** `753`

## Part 2: Constrained Path Counting

Count all paths from `svr` to `out` that visit **both** `dac` and `fft` (in any
order).

**Approach:** State-space DP using a `(node, bitmask)` pair where the 2-bit mask
tracks which of the two required nodes (`dac` = bit 0, `fft` = bit 1) have been
visited so far. A path is only counted when it reaches `out` with both bits set
(mask == 3).

**Complexity:** O(4 * (V + E)) — the DAG is traversed once per possible mask
value (4 states per node).

**Answer:** `450854305019580`
