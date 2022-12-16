package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type Sensor struct {
	X, Y            int
	ClosestDistance int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Manhattan(x0, y0, x1, y1 int) int {
	return abs(x0-x1) + abs(y0-y1)
}

var sensors = []*Sensor{}

func CountAtY(y, min, max int) int {
	count := -1
	for x := min; x < max; x++ {
		for _, s := range sensors {
			if Manhattan(s.X, s.Y, x, y) <= s.ClosestDistance {
				count++
				break
			}
		}
	}
	return count
}

const MinP2, MaxP2 = 0, 4000000

func IsDistress(x, y int) bool {
	for _, s := range sensors {
		if Manhattan(x, y, s.X, s.Y) <= s.ClosestDistance {
			return false
		}
	}
	return true
}

func FindDistressBeaconXY() (int, int) {
	sort.SliceStable(sensors, func(i, j int) bool {
		return sensors[i].Y > sensors[j].Y
	})
	for _, s := range sensors {
		for y := s.Y - s.ClosestDistance - 1; y <= s.Y+s.ClosestDistance+1; y++ {
			if y < MinP2 {
				continue
			}
			if y > MaxP2 {
				break
			}
			if v := abs(y - s.Y); v < s.ClosestDistance {
				dx := (s.ClosestDistance - v + 1)
				x0, x1 := s.X-dx, s.X+dx
				if x0 >= MinP2 && x0 <= MaxP2 {
					if IsDistress(x0, y) {
						return x0, y
					}
				}
				if x0 != x1 && x1 >= MinP2 && x1 <= MaxP2 {
					if IsDistress(x1, y) {
						return x1, y
					}
				}
			}
		}
	}

	return -1, -1
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var minX, maxX, bx, by, gd int
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		s := &Sensor{}
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.X, &s.Y, &bx, &by)
		s.ClosestDistance = Manhattan(s.X, s.Y, bx, by)
		sensors = append(sensors, s)
		gd = max(gd, s.ClosestDistance)
		minX = min(s.X, minX)
		maxX = max(s.X, maxX)
	}

	now := time.Now()
	fmt.Println("Part1", CountAtY(2000000, minX-gd, maxX+gd), time.Since(now))
	now = time.Now()
	x, y := FindDistressBeaconXY()
	fmt.Println("Part2", x, y, "signal:", int64(x)*4000000+int64(y), time.Since(now))
}
