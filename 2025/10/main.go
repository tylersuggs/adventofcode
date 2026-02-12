package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type rat struct{ n, d int64 }

func gcd64(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mkr(n, d int64) rat {
	if d < 0 {
		n, d = -n, -d
	}
	if n == 0 {
		return rat{0, 1}
	}
	g := gcd64(n, d)
	return rat{n / g, d / g}
}

func (a rat) add(b rat) rat { return mkr(a.n*b.d+b.n*a.d, a.d*b.d) }
func (a rat) sub(b rat) rat { return mkr(a.n*b.d-b.n*a.d, a.d*b.d) }
func (a rat) mul(b rat) rat { return mkr(a.n*b.n, a.d*b.d) }
func (a rat) div(b rat) rat { return mkr(a.n*b.d, a.d*b.n) }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	totalP1, totalP2 := 0, 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Parse indicator pattern from [...]
		bracketStart := strings.Index(line, "[")
		bracketEnd := strings.Index(line, "]")
		pattern := line[bracketStart+1 : bracketEnd]
		nCounters := len(pattern)

		xorTarget := 0
		for i, ch := range pattern {
			if ch == '#' {
				xorTarget |= 1 << i
			}
		}

		// Parse buttons from (...) groups
		rest := line[bracketEnd+1:]
		var buttonMasks []int
		var buttonSets [][]int
		for {
			ps := strings.Index(rest, "(")
			if ps == -1 {
				break
			}
			pe := strings.Index(rest, ")")
			btnStr := rest[ps+1 : pe]

			mask := 0
			var counters []int
			for _, s := range strings.Split(btnStr, ",") {
				num, _ := strconv.Atoi(strings.TrimSpace(s))
				mask |= 1 << num
				counters = append(counters, num)
			}
			buttonMasks = append(buttonMasks, mask)
			buttonSets = append(buttonSets, counters)
			rest = rest[pe+1:]
		}

		// Parse joltage targets from {...}
		braceStart := strings.Index(line, "{")
		braceEnd := strings.Index(line, "}")
		var targets []int
		for _, s := range strings.Split(line[braceStart+1:braceEnd], ",") {
			v, _ := strconv.Atoi(strings.TrimSpace(s))
			targets = append(targets, v)
		}

		// Part 1: XOR toggle, brute force all 2^m subsets
		m := len(buttonMasks)
		minP1 := m + 1
		for mask := 0; mask < (1 << m); mask++ {
			pc := bits.OnesCount(uint(mask))
			if pc >= minP1 {
				continue
			}
			state := 0
			for i := 0; i < m; i++ {
				if mask&(1<<i) != 0 {
					state ^= buttonMasks[i]
				}
			}
			if state == xorTarget {
				minP1 = pc
			}
		}
		totalP1 += minP1

		// Part 2: additive counters, minimize total presses
		totalP2 += solvePart2(nCounters, m, buttonSets, targets)
	}

	fmt.Println("Part 1:", totalP1)
	fmt.Println("Part 2:", totalP2)
}

// solvePart2 finds minimum sum(x) such that Ax = targets, x >= 0 integer
// where A[i][j] = 1 if button j affects counter i
func solvePart2(n, m int, buttons [][]int, targets []int) int {
	// Build augmented matrix [A | t] with rational entries
	aug := make([][]rat, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]rat, m+1)
		for j := 0; j <= m; j++ {
			aug[i][j] = rat{0, 1}
		}
		aug[i][m] = mkr(int64(targets[i]), 1)
	}
	for j, btn := range buttons {
		for _, c := range btn {
			aug[c][j] = rat{1, 1}
		}
	}

	// Gaussian elimination → RREF
	var pivotCols []int
	row := 0
	for col := 0; col < m && row < n; col++ {
		pr := -1
		for r := row; r < n; r++ {
			if aug[r][col].n != 0 {
				pr = r
				break
			}
		}
		if pr == -1 {
			continue
		}
		aug[row], aug[pr] = aug[pr], aug[row]

		pivot := aug[row][col]
		for j := 0; j <= m; j++ {
			aug[row][j] = aug[row][j].div(pivot)
		}

		for r := 0; r < n; r++ {
			if r == row || aug[r][col].n == 0 {
				continue
			}
			fac := aug[r][col]
			for j := 0; j <= m; j++ {
				aug[r][j] = aug[r][j].sub(fac.mul(aug[row][j]))
			}
		}

		pivotCols = append(pivotCols, col)
		row++
	}

	rank := len(pivotCols)

	// Check consistency: zero rows must have zero target
	for i := rank; i < n; i++ {
		if aug[i][m].n != 0 {
			return 1 << 30
		}
	}

	// Identify free variables
	pivotSet := make(map[int]bool)
	for _, c := range pivotCols {
		pivotSet[c] = true
	}
	var freeCols []int
	for j := 0; j < m; j++ {
		if !pivotSet[j] {
			freeCols = append(freeCols, j)
		}
	}
	nFree := len(freeCols)

	// Extract RREF coefficients: x_{pivot[i]} = rhs[i] - sum_j coeff[i][j] * f_j
	rhs := make([]rat, rank)
	coeff := make([][]rat, rank)
	for i := 0; i < rank; i++ {
		rhs[i] = aug[i][m]
		coeff[i] = make([]rat, nFree)
		for j := 0; j < nFree; j++ {
			coeff[i][j] = aug[i][freeCols[j]]
		}
	}

	// No free variables: unique solution
	if nFree == 0 {
		total := 0
		for i := 0; i < rank; i++ {
			if rhs[i].d != 1 || rhs[i].n < 0 {
				return 1 << 30
			}
			total += int(rhs[i].n)
		}
		return total
	}

	// Conservative upper bound for any free variable: sum of all targets
	globalUB := int64(0)
	for _, t := range targets {
		globalUB += int64(t)
	}

	// Enumerate free variables recursively with pruning
	// Pre-allocate residual arrays per recursion level
	resLevels := make([][]rat, nFree+1)
	for i := range resLevels {
		resLevels[i] = make([]rat, rank)
	}
	copy(resLevels[0], rhs)

	best := int64(1 << 60)

	var enumerate func(idx int, partialSum int64)
	enumerate = func(idx int, partialSum int64) {
		if partialSum >= best {
			return
		}

		residuals := resLevels[idx]

		if idx == nFree {
			total := partialSum
			for i := 0; i < rank; i++ {
				if residuals[i].d != 1 || residuals[i].n < 0 {
					return
				}
				total += residuals[i].n
			}
			if total < best {
				best = total
			}
			return
		}

		// Upper bound from non-negativity of pivot variables
		// Account for future free vars with negative coefficients that can increase residuals
		ub := globalUB
		for i := 0; i < rank; i++ {
			c := coeff[i][idx]
			if c.n > 0 {
				// Max possible help from future free vars (negative coefficients increase residual)
				help := mkr(0, 1)
				for j := idx + 1; j < nFree; j++ {
					if coeff[i][j].n < 0 {
						help = help.sub(coeff[i][j].mul(mkr(globalUB, 1)))
					}
				}
				relaxedRes := residuals[i].add(help)
				if relaxedRes.n < 0 {
					return // infeasible even with max future help
				}
				fl := relaxedRes.div(c)
				limit := fl.n / fl.d
				if limit < ub {
					ub = limit
				}
			}
		}
		if ub < 0 {
			return
		}

		// Incremental update: subtract coeff[i][idx] per unit of free var
		nextRes := resLevels[idx+1]
		copy(nextRes, residuals)

		for v := int64(0); v <= ub; v++ {
			if v > 0 {
				for i := 0; i < rank; i++ {
					nextRes[i] = nextRes[i].sub(coeff[i][idx])
				}
			}
			enumerate(idx+1, partialSum+v)
		}
	}

	enumerate(0, 0)
	return int(best)
}
