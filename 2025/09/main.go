package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct{ x, y int }
type HEdge struct{ y, xmin, xmax int }
type VEdge struct{ x, ymin, ymax int }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pts []Point
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		pts = append(pts, Point{x, y})
	}

	n := len(pts)

	// Part 1: largest rectangle from any pair of red tiles as opposing corners
	best1 := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := pts[i].x - pts[j].x
			if dx < 0 {
				dx = -dx
			}
			dy := pts[i].y - pts[j].y
			if dy < 0 {
				dy = -dy
			}
			area := (dx + 1) * (dy + 1)
			if area > best1 {
				best1 = area
			}
		}
	}
	fmt.Println("Part 1:", best1)

	// Build polygon edges (consecutive red tiles connected by straight lines)
	var hedges []HEdge
	var vedges []VEdge
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		p1, p2 := pts[i], pts[j]
		if p1.y == p2.y {
			xmin, xmax := p1.x, p2.x
			if xmin > xmax {
				xmin, xmax = xmax, xmin
			}
			hedges = append(hedges, HEdge{p1.y, xmin, xmax})
		} else {
			ymin, ymax := p1.y, p2.y
			if ymin > ymax {
				ymin, ymax = ymax, ymin
			}
			vedges = append(vedges, VEdge{p1.x, ymin, ymax})
		}
	}

	// Point-in-polygon using horizontal ray cast (count vertical edge crossings)
	pointInPoly := func(px, py float64) bool {
		count := 0
		for _, e := range vedges {
			ex := float64(e.x)
			eyMin := float64(e.ymin)
			eyMax := float64(e.ymax)
			if ex > px && eyMin <= py && py < eyMax {
				count++
			}
		}
		return count%2 == 1
	}

	// Check if axis-aligned rectangle is entirely inside the polygon
	rectInPoly := func(xmin, xmax, ymin, ymax int) bool {
		// No horizontal polygon edge may cross the open interior
		for _, e := range hedges {
			if ymin < e.y && e.y < ymax {
				if max(e.xmin, xmin) < min(e.xmax, xmax) {
					return false
				}
			}
		}
		// No vertical polygon edge may cross the open interior
		for _, e := range vedges {
			if xmin < e.x && e.x < xmax {
				if max(e.ymin, ymin) < min(e.ymax, ymax) {
					return false
				}
			}
		}
		// Verify the rectangle is inside (not outside) the polygon
		cx := (float64(xmin) + float64(xmax)) / 2.0
		cy := (float64(ymin) + float64(ymax)) / 2.0
		// Offset to ensure non-integer coordinates (polygon has integer coords)
		if cx == float64(int(cx)) {
			cx += 0.1
		}
		if cy == float64(int(cy)) {
			cy += 0.1
		}
		return pointInPoly(cx, cy)
	}

	// Part 2: largest rectangle inside polygon with red tiles at opposing corners
	best2 := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := pts[i].x - pts[j].x
			if dx < 0 {
				dx = -dx
			}
			dy := pts[i].y - pts[j].y
			if dy < 0 {
				dy = -dy
			}
			if dx == 0 || dy == 0 {
				continue
			}
			area := (dx + 1) * (dy + 1)
			if area <= best2 {
				continue // prune: can't beat current best
			}
			xmin := min(pts[i].x, pts[j].x)
			xmax := max(pts[i].x, pts[j].x)
			ymin := min(pts[i].y, pts[j].y)
			ymax := max(pts[i].y, pts[j].y)
			if rectInPoly(xmin, xmax, ymin, ymax) {
				best2 = area
			}
		}
	}
	fmt.Println("Part 2:", best2)
}
