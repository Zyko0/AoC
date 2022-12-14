package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Equal    = 0
	Inferior = 1
	Superior = 2
)

type Data interface {
	Cmp(Data) int
	String() string
}

type Node int
type DataList []Data

func (n Node) Cmp(d Data) int {
	switch dt := d.(type) {
	case Node:
		switch {
		case n == dt:
			return Equal
		case n < dt:
			return Inferior
		case n > dt:
			return Superior
		}
	case DataList:
		return DataList{n}.Cmp(d)
	}
	return Equal
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n)
}

func (dl DataList) Cmp(d Data) int {
	switch v := d.(type) {
	case Node:
		return dl.Cmp(DataList{v})
	case DataList:
		for i := range dl {
			if i >= len(v) {
				return Superior
			}
			switch cmp := dl[i].Cmp(v[i]); cmp {
			case Equal:
				continue
			default:
				return cmp
			}
		}
		if len(dl) < len(v) {
			return Inferior
		}
	}

	return Equal
}

func (dl DataList) Unmarshal(set string) DataList {
	if set == "" {
		return dl
	}
	// Remove commas
	set = strings.TrimLeft(set, ",")
	// Check if number
	if set[0] >= '0' && set[0] <= '9' {
		v := 0
		length, _ := fmt.Sscanf(set, "%d", &v)
		dl = append(dl, Node(v))
		return dl.Unmarshal(set[length:])
	}
	// List
	s := set
	if set[0] == '[' {
		req, i := 1, 1
		for ; req > 0; i++ {
			switch s[i] {
			case '[':
				req++
			case ']':
				req--
			}
		}
		s = s[:i]
	}
	ss := s[1 : len(s)-1]
	if len(ss) > 0 && (ss[0] == '[' || strings.Count(ss, ",") > 0) {
		dl = append(dl, DataList{}.Unmarshal(ss))
	} else {
		ndl := DataList{}
		if ss != "" {
			v, _ := strconv.Atoi(ss)
			ndl = append(ndl, DataList{Node(v)})
		}
		dl = append(dl, ndl)
	}

	return dl.Unmarshal(strings.Replace(set, s, "", 1))
}

func (dl DataList) String() string {
	s := "["
	for i, n := range dl {
		s += n.String()
		if i < len(dl)-1 {
			s += ","
		}
	}
	return s + "]"
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	pairs := make([][2]DataList, 0)
	pairIndex := 0
	var p0, p1 DataList
	for _, l := range lines {
		if l == "" {
			continue
		}
		l = l[1 : len(l)-1]
		list := DataList{}.Unmarshal(l)
		switch pairIndex {
		case 0:
			p0 = list
		case 1:
			p1 = list
			pairs = append(pairs, [2]DataList{p0, p1})
		}
		pairIndex = (pairIndex + 1) % 2
	}
	// Part 1
	sum := 0
	for i, p := range pairs {
		if p[0].Cmp(p[1]) <= Inferior {
			sum += i + 1
		}
	}
	fmt.Println("Part1:", sum)
	// Part 2
	flatten := []DataList{}
	for i := range pairs {
		flatten = append(flatten, pairs[i][:]...)
	}
	flatten = append(flatten, DataList{DataList{Node(2)}})
	flatten = append(flatten, DataList{DataList{Node(6)}})
	sort.SliceStable(flatten, func(i, j int) bool {
		cmp := flatten[i].Cmp(flatten[j])
		return cmp == Inferior
	})
	i0, i1 := 0, 0
	for i := range flatten {
		switch flatten[i].String() {
		case "[[2]]":
			i0 = i + 1
		case "[[6]]":
			i1 = i + 1
		}
	}
	fmt.Println("Part2:", i0*i1)
}
