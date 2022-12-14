package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 50 stars avant le 25/12
func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	calories := []int{0}
	calIndex := 0
	for i := 0; i < len(lines); i++ {
		var j = i

		calories = append(calories, 0)
		for ; j < len(lines) && lines[j] != ""; j++ {
			cal, err := strconv.Atoi(lines[j])
			if err != nil {
				log.Fatal(err)
			}
			calories[calIndex] += cal
		}
		i = j
		calIndex++
	}

	sort.SliceStable(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	fmt.Println("most calories:", calories[0]+calories[1]+calories[2])
}
