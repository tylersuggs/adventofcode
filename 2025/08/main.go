package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z float64
}

type UF struct {
	parent []int
	rank   []int
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UF{parent: parent, rank: rank}
}

func (uf *UF) Find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UF) Union(a, b int) {
	ra, rb := uf.Find(a), uf.Find(b)
	if ra == rb {
		return
	}
	if uf.rank[ra] < uf.rank[rb] {
		ra, rb = rb, ra
	}
	uf.parent[rb] = ra
	if uf.rank[ra] == uf.rank[rb] {
		uf.rank[ra]++
	}
}

type Edge struct {
	i, j int
	d    float64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var points []Point

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		z, _ := strconv.ParseFloat(parts[2], 64)
		points = append(points, Point{x, y, z})
	}

	n := len(points)

	// Build all pairwise edges sorted by Euclidean distance
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			d := math.Sqrt(dx*dx + dy*dy + dz*dz)
			edges = append(edges, Edge{i, j, d})
		}
	}
	sort.Slice(edges, func(a, b int) bool { return edges[a].d < edges[b].d })

	// Part 1: Make the 1000 shortest connections, form connected components,
	// report product of three largest component sizes.
	uf := NewUF(n)
	connections := 1000
	for i := 0; i < connections; i++ {
		uf.Union(edges[i].i, edges[i].j)
	}

	sizes := make(map[int]int)
	for i := 0; i < n; i++ {
		sizes[uf.Find(i)]++
	}
	sizeList := make([]int, 0, len(sizes))
	for _, s := range sizes {
		sizeList = append(sizeList, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizeList)))

	fmt.Println("Part 1:", sizeList[0]*sizeList[1]*sizeList[2])

	// Part 2: Keep connecting closest unconnected pairs until all in one circuit.
	// Multiply X coordinates of the last two junction boxes connected.
	uf2 := NewUF(n)
	var lastI, lastJ int
	for _, e := range edges {
		if uf2.Find(e.i) != uf2.Find(e.j) {
			uf2.Union(e.i, e.j)
			lastI = e.i
			lastJ = e.j
		}
	}
	fmt.Println("Part 2:", int(points[lastI].x)*int(points[lastJ].x))
}
