package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Win  = 6
	Draw = 3
	Lose = 0
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

var (
	inToKind = map[string]int{
		"A": Rock,
		"X": Rock,
		"B": Paper,
		"Y": Paper,
		"C": Scissors,
		"Z": Scissors,
	}

	inToState = map[string]int{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}
)

func getScore(kindE, kindMe int) (uint64, int) {
	var state int

	score := kindMe
	switch {
	// Draw
	case kindMe == kindE:
		state = Draw
	case kindMe == Rock && kindE == Paper:
		state = Lose
	case kindMe == Rock && kindE == Scissors:
		state = Win
	case kindMe == Paper && kindE == Rock:
		state = Win
	case kindMe == Paper && kindE == Scissors:
		state = Lose
	case kindMe == Scissors && kindE == Paper:
		state = Win
	case kindMe == Scissors && kindE == Rock:
		state = Lose
	}

	return uint64(score + state), state
}

func GetScore(inE, inMe string) uint64 {
	kindE, kindMe := inToKind[inE], inToKind[inMe]
	score, _ := getScore(kindE, kindMe)

	return score
}

func GetScorePart2(inE, inState string) uint64 {
	kindE, state := inToKind[inE], inToState[inState]

	var kindMe int
	for kindMe = Rock; kindMe <= Scissors; kindMe++ {
		if score, st := getScore(kindE, kindMe); st == state {
			return score
		}
	}

	log.Fatal("oops")
	return 0
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	totalScore := uint64(0)
	for _, l := range lines {
		spl := strings.Split(l, " ")
		// Part 1
		/*inE, inMe := strings.TrimSpace(spl[0]), strings.TrimSpace(spl[1])
		totalScore += GetScore(inE, inMe)*/

		// Part 2
		inE, inState := strings.TrimSpace(spl[0]), strings.TrimSpace(spl[1])
		totalScore += GetScorePart2(inE, inState)
	}

	fmt.Println("total score:", totalScore)
}
