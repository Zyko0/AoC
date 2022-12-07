package main

import (
	"fmt"
	"log"
	"os"
)

func uniqueRunes(s string) bool {
	for i, r0 := range s {
		for j, r1 := range s {
			if i != j && r0 == r1 {
				return false
			}
		}
	}

	return true
}

func FindFirstPacketIndex(s string, packetSize int) int {
	for i := range s {
		if uniqueRunes(s[i : i+packetSize]) {
			return i + packetSize
		}
	}
	return -1
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part1:", FindFirstPacketIndex(string(b), 4))
	fmt.Println("Part2:", FindFirstPacketIndex(string(b), 14))
}
