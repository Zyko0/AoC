package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Pos struct {
	X, Y int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func (p *Pos) Touching(p2 *Pos) bool {
	return abs(p2.X-p.X) <= 1 && abs(p2.Y-p.Y) <= 1
}

func NewPos() *Pos {
	return &Pos{
		X: 0,
		Y: 0,
	}
}

var marks = map[Pos]struct{}{}

func maxAbsDist(a, b int) int {
	if abs(a) > abs(b) {
		return a
	}
	return b
}

func RunMove(dir string, count int, head, tail *Pos, mark bool) {
	dx, dy := 0, 0
	switch dir {
	case "U":
		dy = -1
	case "D":
		dy = 1
	case "R":
		dx = 1
	case "L":
		dx = -1
	}
	for i := 0; i < count; i++ {
		head.X, head.Y = head.X+dx, head.Y+dy
		if !tail.Touching(head) {
			distX, distY := head.X-tail.X, head.Y-tail.Y
			if tail.X != head.X && tail.Y != head.Y {
				switch maxAbsDist(distX, distY) {
				case distX, distY:
					tail.X += sign(distX)
					tail.Y += sign(distY)
				case distX:
					tail.Y += sign(distY)
				case distY:
					tail.X += sign(distX)
				}
			}
			if !tail.Touching(head) {
				if head.X == tail.X {
					tail.Y += sign(distY)
				}
				if head.Y == tail.Y {
					tail.X += sign(distX)
				}
			}
		}
		if mark {
			marks[*tail] = struct{}{}
		}
	}
}

func RunMoves2(dir string, count int, knots []*Pos) {
	for i := 0; i < count; i++ {
		for j := len(knots) - 1; j > 0; j-- {
			var d, c = "none", 1
			if j == len(knots)-1 {
				d, c = dir, 1
			}
			RunMove(d, c, knots[j], knots[j-1], j-1 == 0)
		}
	}
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	head, tail := NewPos(), NewPos()
	for _, l := range lines {
		var dir string
		var count int

		fmt.Sscanf(l, "%s %d", &dir, &count)
		RunMove(dir, count, head, tail, true)
	}
	fmt.Println("Part1:", len(marks))

	marks = map[Pos]struct{}{}
	knots := make([]*Pos, 10)
	for i := 0; i < len(knots); i++ {
		knots[i] = NewPos()
	}
	for _, l := range lines {
		var dir string
		var count int

		fmt.Sscanf(l, "%s %d", &dir, &count)
		RunMoves2(dir, count, knots)
	}
	fmt.Println("Part2:", len(marks))
}
