package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	buffer = make([]*Cell, 0, 4)
	grid   [][]*Cell
	sn, gn *Cell
)

type Cell struct {
	X, Y   int
	Height rune
}

func (c *Cell) GetNeighbors() []*Cell {
	x, y := c.X, c.Y
	allowed := c.Height + 1
	buffer = buffer[:0]
	if x > 0 && grid[y][x-1].Height <= allowed {
		buffer = append(buffer, grid[y][x-1])
	}
	if x < len(grid[0])-1 && grid[y][x+1].Height <= allowed {
		buffer = append(buffer, grid[y][x+1])
	}
	if y > 0 && grid[y-1][x].Height <= allowed {
		buffer = append(buffer, grid[y-1][x])
	}
	if y < len(grid)-1 && grid[y+1][x].Height <= allowed {
		buffer = append(buffer, grid[y+1][x])
	}
	return buffer
}

func BFSPathLen(root, end *Cell) int {
	q := []*Cell{root}
	exploredToParents := map[*Cell]*Cell{root: nil}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v == end {
			cnt := 0
			for v != root {
				v = exploredToParents[v]
				cnt++
			}
			return cnt
		}
		for _, n := range v.GetNeighbors() {
			if _, ok := exploredToParents[n]; !ok {
				exploredToParents[n] = v
				q = append(q, n)
			}
		}
	}

	return 0
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	grid = make([][]*Cell, len(lines))
	for y, l := range lines {
		grid[y] = make([]*Cell, len(l))
		for x, r := range l {
			n := &Cell{
				Height: r,
				X:      x,
				Y:      y,
			}
			switch r {
			case 'S':
				n.Height = 'a' // Consider starting as 0 height
				sn = n
			case 'E':
				n.Height = 'z' // Consider end as max height
				gn = n
			}
			grid[y][x] = n
		}
	}
	shortest := BFSPathLen(sn, gn)
	fmt.Println("Part1:", shortest)
	for y := range grid {
		for x := range grid[y] {
			if n := grid[y][x]; n.Height == 'a' {
				if length := BFSPathLen(n, gn); length > 0 && length < shortest {
					shortest = length
				}
			}
		}
	}
	fmt.Println("Part2:", shortest)
}
