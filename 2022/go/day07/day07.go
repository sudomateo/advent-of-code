package day07

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (int, error) {
	d, err := buildTree(r)
	if err != nil {
		return 0, err
	}

	sizes := make([]int, 0)
	_ = calculateSize(d, &sizes)

	var totalSize int

	for _, size := range sizes {
		if size <= 100000 {
			totalSize += size
		}
	}

	return totalSize, nil
}

func Part2(r io.Reader) (int, error) {
	d, err := buildTree(r)
	if err != nil {
		return 0, err
	}

	sizeAvailable := 70000000
	freeNeeded := 30000000

	sizes := make([]int, 0)
	sizeUsed := calculateSize(d, &sizes)

	freeSpace := sizeAvailable - sizeUsed
	needToDelete := freeNeeded - freeSpace

	sort.SliceStable(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})

	var deletedSize int
	for _, size := range sizes {
		if size >= needToDelete {
			deletedSize = size
			break
		}
	}

	return deletedSize, nil
}

type dir struct {
	name   string
	parent *dir
	files  []file
	dirs   []*dir
}

type file struct {
	name string
	size int
}

func buildTree(r io.Reader) (*dir, error) {
	var rootDir *dir
	var currDir *dir

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, " ")

		switch fields[0] {
		case "$":
			cmd := fields[1]

			switch cmd {
			case "cd":
				path := fields[2]
				switch path {
				case "/":
					d := dir{
						name:   path,
						parent: nil,
						files:  make([]file, 0),
						dirs:   make([]*dir, 0),
					}
					rootDir = &d
					currDir = &d
				case "..":
					currDir = currDir.parent
				default:
					var d *dir
					for _, v := range currDir.dirs {
						if v.name == path {
							d = v
							break
						}
					}

					if d == nil {
						d = &dir{
							name:   path,
							parent: currDir,
							files:  make([]file, 0),
							dirs:   make([]*dir, 0),
						}
						currDir.dirs = append(currDir.dirs, d)
					}

					currDir = d
				}
			case "ls":
				continue
			default:
				return nil, fmt.Errorf("invalid command %s", cmd)
			}
		default:
			switch fields[0] {
			case "dir":
				var d *dir
				for _, v := range currDir.dirs {
					if v.name == fields[1] {
						d = v
						break
					}
				}

				if d == nil {
					d = &dir{
						name:   fields[1],
						parent: currDir,
						files:  make([]file, 0),
						dirs:   make([]*dir, 0),
					}
					currDir.dirs = append(currDir.dirs, d)
				}
			default:
				size, err := strconv.Atoi(fields[0])
				if err != nil {
					return nil, fmt.Errorf("could not parse size for %s: %w", line, err)
				}
				f := file{
					name: fields[1],
					size: size,
				}
				currDir.files = append(currDir.files, f)
			}
		}
	}

	return rootDir, nil
}

func calculateSize(d *dir, sizes *[]int) int {
	if d == nil {
		return 0
	}

	size := 0
	for _, f := range d.files {
		size += f.size
	}

	for _, v := range d.dirs {
		size += calculateSize(v, sizes)
	}

	*sizes = append(*sizes, size)

	return size
}

func printTree(d *dir, pad string) {
	if d == nil {
		return
	}

	fmt.Printf("%s- %s (dir)\n", pad, d.name)
	pad = pad + "  "

	for _, v := range d.dirs {
		printTree(v, pad)
	}

	for _, f := range d.files {
		fmt.Printf("%s- %s (file, size=%d)\n", pad, f.name, f.size)
	}
}
