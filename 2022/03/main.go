package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func itemToPrio(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	}
	if r >= 'A' && r <= 'Z' {
		return int(r-64) + 26
	}
	return -1
}

func getPrioItemDuplicate(comp0, comp1 string) int {
	for _, r0 := range comp0 {
		for _, r1 := range comp1 {
			// Found the duplicate
			if r0 == r1 {
				return itemToPrio(r0)
			}
		}
	}
	// No duplicate
	return 0
}

func getPrioGroupBadge(grp0, grp1, grp2 string) int {
	for _, r0 := range grp0 {
		for _, r1 := range grp1 {
			if r0 == r1 {
				for _, r2 := range grp2 {
					// Found the badge
					if r0 == r2 {
						return itemToPrio(r0)
					}
				}
			}
		}
	}
	log.Fatal("no badge??")
	return -1
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	sumPrio := 0
	// Part 1
	/*
		for _, l := range lines {
				comp0 := l[:len(l)/2]
				comp1 := l[len(l)/2:]
				prio := getPrioItemDuplicate(comp0, comp1)
				sumPrio += prio
		}
	*/
	// Part 2
	for i := 0; i < len(lines); i += 3 {
		grp0, grp1, grp2 := lines[i], lines[i+1], lines[i+2]
		prio := getPrioGroupBadge(grp0, grp1, grp2)
		sumPrio += prio
	}

	fmt.Println("ok:", sumPrio)
}
