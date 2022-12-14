package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Arg struct {
	Time  int
	Value int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	args := []*Arg{}
	x, lastCycle, sum1, rowX := 1, 0, 0, 0
	for cycle := 0; cycle < 240; cycle++ {
		var cmd = "noop"
		var n int

		if cycle < len(lines) {
			fmt.Sscanf(lines[cycle], "%s %d", &cmd, &n)
		}
		lastCycle++
		if cmd != "noop" {
			lastCycle++
			args = append(args, &Arg{
				Time:  lastCycle + 1,
				Value: n,
			})
		}
		for _, a := range args {
			if cycle == a.Time {
				x += a.Value
			}
		}

		if (cycle-20)%40 == 0 {
			sum1 += cycle * x
		}
		if cycle == 0 {
			continue
		}
		if abs(x-rowX) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		rowX++
		if rowX == 40 {
			fmt.Println()
			rowX = 0
		}
	}

	fmt.Println("\nPart1:", sum1)
}
