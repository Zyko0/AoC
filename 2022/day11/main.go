package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func ApplyRelief(n uint64) uint64 {
	return uint64(n / 3)
}
func ApplyNoop(n uint64) uint64 {
	return n
}

type ReliefFn func(uint64) uint64
type WorryOpFn func(uint64) uint64

type Monkey struct {
	Items        []uint64
	InspectOp    WorryOpFn
	ModTest      uint64
	MTrue        *Monkey
	MFalse       *Monkey
	InspectCount uint64
}

const monkeyCount = 8
const filename = "input.txt"

func (m *Monkey) InspectAndThrow(relief ReliefFn) {
	m.InspectCount += uint64(len(m.Items))
	for _, item := range m.Items {
		worry := m.InspectOp(item)
		worry = relief(worry)
		if worry%m.ModTest == 0 {
			m.MTrue.Items = append(m.MTrue.Items, worry)
		} else {
			m.MFalse.Items = append(m.MFalse.Items, worry)
		}
	}
}

var (
	regItems = regexp.MustCompile(`Starting items: (.*)`)
)

func main() {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	monkeys := make([]*Monkey, monkeyCount)
	lines := strings.Split(string(b), "\n")
	for i := range monkeys {
		monkeys[i] = &Monkey{
			InspectCount: 0,
		}
	}
	for i := 0; i < len(lines); i += 7 {
		if strings.ReplaceAll(lines[i], " ", "") == "" {
			continue
		}

		var index, value, mTrue, mFalse, modv int
		var op rune
		fmt.Sscanf(lines[i], "Monkey %d:", &index)
		spl := strings.Split(lines[i+1], ":")
		spl = strings.Split(spl[1], ",")
		for _, w := range spl {
			v, _ := strconv.Atoi(strings.TrimSpace(w))
			monkeys[index].Items = append(monkeys[index].Items, uint64(v))
		}
		fmt.Sscanf(strings.TrimSpace(lines[i+2]), "Operation: new = old %c %d", &op, &value)
		fmt.Sscanf(strings.TrimSpace(lines[i+3]), "Test: divisible by %d", &modv)
		fmt.Sscanf(strings.TrimSpace(lines[i+4]), "If true: throw to monkey %d", &mTrue)
		fmt.Sscanf(strings.TrimSpace(lines[i+5]), "If false: throw to monkey %d", &mFalse)
		vv := uint64(value)
		monkeys[index].InspectOp = func(n uint64) uint64 {
			v := vv
			if value == 0 {
				v = n
			}
			if op == '*' {
				return n * v
			}
			return n + v
		}
		monkeys[index].MTrue, monkeys[index].MFalse = monkeys[mTrue], monkeys[mFalse]
		monkeys[index].ModTest = uint64(modv)
	}
	// Saving original items
	monkeysItems := make([][]uint64, len(monkeys))
	for i := range monkeys {
		monkeysItems[i] = append(monkeysItems[i], monkeys[i].Items...)
	}
	// Part 1
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.InspectAndThrow(ApplyRelief)
			m.Items = m.Items[:0]
		}
	}
	monkeysP1 := append([]*Monkey{}, monkeys...)
	sort.SliceStable(monkeysP1, func(i, j int) bool {
		return monkeysP1[i].InspectCount > monkeysP1[j].InspectCount
	})
	fmt.Println("Part1:",
		monkeysP1[0].InspectCount*monkeysP1[1].InspectCount,
	)
	// Resetting monkeys
	for i := range monkeys {
		monkeys[i].Items = append(monkeys[i].Items[:0], monkeysItems[i]...)
		monkeys[i].InspectCount = 0
	}
	// Part 2
	mod := uint64(1)
	for _, m := range monkeys {
		mod *= m.ModTest
	}
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			for j := range m.Items {
				m.Items[j] = m.Items[j] % mod
			}
			m.InspectAndThrow(ApplyNoop)
			m.Items = m.Items[:0]
		}
	}
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].InspectCount > monkeys[j].InspectCount
	})
	fmt.Println("Part2:",
		monkeys[0].InspectCount*monkeys[1].InspectCount,
	)
}
