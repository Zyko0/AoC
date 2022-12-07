package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type CratePile []rune

type CratePileList []CratePile

func (cl CratePileList) Move(amount, from, to int) {
	from, to = from-1, to-1

	for i := 0; i < amount; i++ {
		item := cl[from][0]
		cl[from] = cl[from][1:]
		cl[to] = append(CratePile{item}, cl[to]...)
	}
}

func (cl CratePileList) Move2(amount, from, to int) {
	from, to = from-1, to-1

	items := cl[from][0:amount]
	cl[from] = cl[from][amount:]

	v := append(CratePile{}, items...)
	v = append(v, cl[to]...)
	cl[to] = v
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	// Crate piles0
	piles0 := make(CratePileList, 9)
	piles1 := make(CratePileList, 9)
	i := 0
	init := false
	regx := regexp.MustCompile(`(\s{4}|[A-Z])`)
	for ; !init && i < len(lines); i++ {
		bl := []byte(lines[i])
		if regx.Match(bl) {
			results := regx.FindAll(bl, 9)
			for i, r := range results {
				if strings.ReplaceAll(string(r), " ", "") != "" {
					piles0[i] = append(piles0[i], rune(r[0]))
					piles1[i] = append(piles1[i], rune(r[0]))
				}
			}
		}
		if strings.ReplaceAll(lines[i], " ", "") == "123456789" {
			for ; !init && i < len(lines); i++ {
				if strings.HasPrefix(lines[i], "move") {
					init = true
					break
				}
			}
		}
	}
	i-- // i has been incremented one more time

	// Treat moves
	for ; i < len(lines); i++ {
		var amount, src, dst int
		fmt.Sscanf(lines[i], "move %d from %d to %d", &amount, &src, &dst)
		piles0.Move(amount, src, dst)
		piles1.Move2(amount, src, dst)
	}
	// Part 1
	topCrates0 := ""
	for _, p := range piles0 {
		topCrates0 += string(p[0])
	}
	// Part 2
	topCrates1 := ""
	for _, p := range piles1 {
		topCrates1 += string(p[0])
	}

	fmt.Println("Part 1:", topCrates0)
	fmt.Println("Part 2:", topCrates1)

	for _, p := range piles1 {
		for _, cc := range p {
			fmt.Print(string(cc), ",")
		}
		fmt.Println()
	}
}
