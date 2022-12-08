package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Tree struct {
	Height  int
	Visible bool
}

type Forest [][]*Tree

func (f Forest) Width() int {
	return len(f[0])
}

func (f Forest) Height() int {
	return len(f)
}

func (f Forest) Initialize() {
	for y := range f {
		for x, t := range f[y] {
			v := true
			for tx := x - 1; tx >= 0; tx-- {
				v = v && t.Height > f[y][tx].Height
			}
			t.Visible = t.Visible || v
			v = true
			for tx := x + 1; tx < f.Width(); tx++ {
				v = v && t.Height > f[y][tx].Height
			}
			t.Visible = t.Visible || v
			v = true
			for ty := y - 1; ty >= 0; ty-- {
				v = v && t.Height > f[ty][x].Height
			}
			t.Visible = t.Visible || v
			v = true
			for ty := y + 1; ty < f.Height(); ty++ {
				v = v && t.Height > f[ty][x].Height
			}
			t.Visible = t.Visible || v
		}
	}
}

func (f Forest) CountVisible() int {
	count := 0
	for y := range f {
		for _, t := range f[y] {
			if t.Visible {
				count++
			}
		}
	}

	return count
}

func (f Forest) ComputeScenicAt(x, y int) int {
	scenic := 1
	dist := 0
	t := f[y][x]

	for tx := x - 1; tx >= 0; tx-- {
		dist++
		if t.Height <= f[y][tx].Height {
			break
		}
	}
	scenic, dist = scenic*dist, 0
	for tx := x + 1; tx < f.Width(); tx++ {
		dist++
		if t.Height <= f[y][tx].Height {
			break
		}
	}
	scenic, dist = scenic*dist, 0
	for ty := y - 1; ty >= 0; ty-- {
		dist++
		if t.Height <= f[ty][x].Height {
			break
		}
	}
	scenic, dist = scenic*dist, 0
	for ty := y + 1; ty < f.Height(); ty++ {
		dist++
		if t.Height <= f[ty][x].Height {
			break
		}
	}
	scenic, dist = scenic*dist, 0

	return scenic
}

func (f Forest) GetHighestScenic() int {
	highest := 0

	for y := range f {
		for x := range f[y] {
			if s := f.ComputeScenicAt(x, y); s > highest {
				highest = s
			}
		}
	}

	return highest
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	forest := Forest{}
	for y, l := range lines {
		forest = append(forest, make([]*Tree, len(l)))
		for x, r := range l {
			height := int(r - 48)
			forest[y][x] = &Tree{
				Height: height,
			}
		}
	}
	forest.Initialize()

	fmt.Println("Part1:", forest.CountVisible())
	fmt.Println("Part2:", forest.GetHighestScenic())
}
