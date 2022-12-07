package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	TypeFolder = 0
	TypeFile   = 1
)

type File struct {
	Name     string
	Type     int
	Size     uint64
	Parent   *File
	Children map[string]*File
}

func (f *File) GetSize() uint64 {
	if f.Type == TypeFile {
		return f.Size
	}

	size := uint64(0)
	for _, child := range f.Children {
		size += child.GetSize()
	}

	return size
}

type Shell struct {
	Root    *File
	Current *File
}

func (s *Shell) ExecuteCommand(lines []string, offset int) {
	cmdArgs := strings.Split(lines[offset], " ")
	switch cmdArgs[1] {
	case "cd":
		switch cmdArgs[2] {
		case "..":
			if s.Current.Parent != nil {
				s.Current = s.Current.Parent
			}
		case "/":
			s.Current = s.Root
		default:
			s.Current = s.Current.Children[cmdArgs[2]]
		}
	case "ls":
		for i := offset + 1; i < len(lines) && lines[i][0] != '$'; i++ {
			var f *File
			infoArgs := strings.Split(lines[i], " ")
			switch infoArgs[0] {
			case "dir":
				f = &File{
					Name:     infoArgs[1],
					Type:     TypeFolder,
					Size:     0,
					Parent:   s.Current,
					Children: map[string]*File{},
				}
			default:
				size, _ := strconv.Atoi(infoArgs[0])
				f = &File{
					Name:     infoArgs[1],
					Type:     TypeFile,
					Size:     uint64(size),
					Parent:   s.Current,
					Children: map[string]*File{},
				}
			}
			s.Current.Children[infoArgs[1]] = f
		}
	default:
		panic("oops command couldn't be parsed")
	}
}

func FindPart1(current *File, minSize, maxSize uint64) uint64 {
	sum := uint64(0)
	size := current.GetSize()
	if current.Type == TypeFolder && size >= minSize && size <= maxSize {
		sum += size
	}
	for _, f := range current.Children {
		if f.Type == TypeFolder {
			sum += FindPart1(f, minSize, maxSize)
		}
	}

	return sum
}

func FindAll(current *File, minSize, maxSize uint64) []*File {
	potentialFolders := []*File{}

	size := current.GetSize()
	if current.Type == TypeFolder && size >= minSize && size <= maxSize {
		potentialFolders = append(potentialFolders, current)
	}
	for _, f := range current.Children {
		if f.Type == TypeFolder {
			potentialFolders = append(potentialFolders, FindAll(f, minSize, maxSize)...)
		}
	}

	return potentialFolders
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	root := &File{
		Name:     "/",
		Type:     TypeFolder,
		Size:     0,
		Parent:   nil,
		Children: map[string]*File{},
	}
	shell := &Shell{
		Root:    root,
		Current: root,
	}

	lines := strings.Split(string(b), "\n")
	for i := range lines {
		if strings.HasPrefix(lines[i], "$") {
			shell.ExecuteCommand(lines, i)
		}
	}

	size := FindPart1(shell.Root, 0, 100000)
	fmt.Println("Part1", size)

	requiredSize := 30000000 - (70000000 - shell.Root.GetSize())
	folders := FindAll(shell.Root, requiredSize, math.MaxUint64)
	sort.SliceStable(folders, func(i, j int) bool {
		return folders[i].GetSize() < folders[j].GetSize()
	})
	fmt.Println("requiredSize", requiredSize)
	fmt.Println("Part2", folders[0].GetSize())
}
