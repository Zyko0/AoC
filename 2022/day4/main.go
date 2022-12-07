package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

func (r *Range) Includes(r1 *Range) bool {
	return r.min <= r1.min && r.max >= r1.max
}

func CountOverlaps(r0, r1 *Range) int {
	count := 0
	for i0 := r0.min; i0 <= r0.max; i0++ {
		for i1 := r1.min; i1 <= r1.max; i1++ {
			if i0 == i1 {
				count++
			}
		}
	}

	return count
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	// Part 1
	count := 0
	countOverlaps := 0
	for _, l := range lines {
		ranges := strings.Split(l, ",")
		range0 := strings.Split(ranges[0], "-")
		range1 := strings.Split(ranges[1], "-")
		r00, _ := strconv.Atoi(range0[0])
		r01, _ := strconv.Atoi(range0[1])
		r10, _ := strconv.Atoi(range1[0])
		r11, _ := strconv.Atoi(range1[1])
		r0 := Range{
			min: r00,
			max: r01,
		}
		r1 := Range{
			min: r10,
			max: r11,
		}
		if r0.Includes(&r1) || r1.Includes(&r0) {
			count++
		}
		if CountOverlaps(&r0, &r1) > 0 {
			countOverlaps++
		}
	}

	fmt.Println("Part 1:", count)
	fmt.Println("Part 2:", countOverlaps)
}
