package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Air  = 0
	Rock = 1
	Sand = 2
)

type Tile struct {
	X, Y int
	Kind byte
}

func (t *Tile) Zero() bool {
	return t.X == 0 && t.Y == 0
}

func (t Tile) Add(x, y int) Tile {
	t.X += x
	t.Y += y
	return t
}

type Cave [][]Tile

func (c Cave) Simulate(w, h int) {
	for {
		static, abyss := false, false
		sx, sy := source.X, 0
		for !static {
			if sy >= h || sx <= 0 || sx >= w {
				abyss = true
				break
			}
			switch {
			case c[sy+1][sx].Kind == Air:
				sy++
			case sx > 0 && c[sy+1][sx-1].Kind == Air:
				sy, sx = sy+1, sx-1
			case sx < w && c[sy+1][sx+1].Kind == Air:
				sy, sx = sy+1, sx+1
			default:
				static = true
			}
		}
		if sy == 0 {
			c[sy][sx].Kind = Sand
			break
		}
		if abyss {
			break
		}
		if static {
			if c[sy][sx].Kind != Air {
				break
			}
			c[sy][sx].Kind = Sand
		}
	}
}

func (c Cave) Debug() {
	for _, row := range c {
		for _, t := range row {
			switch t.Kind {
			case Air:
				fmt.Print(".")
			case Rock:
				fmt.Print("#")
			case Sand:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

var (
	cave   = Cave{}
	source = Tile{X: 500}
)

func sub1(a, b int) int {
	v := a - b
	switch {
	case v > 0:
		return 1
	case v < 0:
		return -1
	}
	return 0
}

func main() {
	// Lazy for part 2, bruteforcing the floor line on height+1 does it :)
	// Just adding this to input.txt: `0,168 -> 1000,168`
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	rocks := []Tile{}
	for _, l := range lines {
		tuples := strings.Split(l, " -> ")
		last := Tile{}
		for _, t := range tuples {
			rock := Tile{}
			fmt.Sscanf(t, "%d,%d", &rock.X, &rock.Y)
			if last.Zero() {
				rocks = append(rocks, rock)
			} else {
				for nrock := last; nrock != rock; nrock = nrock.Add(sub1(rock.X, last.X), sub1(rock.Y, last.Y)) {
					rocks = append(rocks, nrock)
				}
				rocks = append(rocks, rock)
			}
			last = rock
		}
	}
	// Find dimensions
	var minX, minY, maxX, maxY = 9999, source.Y, -1, -1
	for _, r := range rocks {
		switch {
		case r.X < minX:
			minX = r.X
		case r.X > maxX:
			maxX = r.X
		case r.Y > maxY:
			maxY = r.Y
		}
	}
	// Fill cave
	w, h := maxX-minX+1, maxY-minY+1
	cave = make([][]Tile, h+1)
	for y := 0; y < h+1; y++ {
		cave[y] = make([]Tile, w+1)
		for x := 0; x < w; x++ {
			cave[y][x] = Tile{
				X: x, Y: y,
			}
		}
	}
	source.X, source.Y = source.X-minX+1, source.Y-minY
	// Fill rocks
	for _, r := range rocks {
		cave[r.Y-minY][r.X-minX+1].Kind = Rock
	}
	cave.Simulate(w, h)
	cnt := 0
	for _, row := range cave {
		for _, t := range row {
			if t.Kind == Sand {
				cnt++
			}
		}
	}
	fmt.Println("Count:", cnt)
}
