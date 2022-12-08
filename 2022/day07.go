// This program implements the solution for https://adventofcode.com/2022/day/7.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day07_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dirs := map[string]*dir{}
	var curPath string
	var curDir *dir
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Store a map of directory paths to a dir struct that contains what files
		// and subdirs are in that dir. Then later on we can do DFS from the root
		// dir to count sizes.
		//
		// E.g., after parsing the input we'll end up with our dirs map to be
		// something like:
		//
		//   "/": {files: ["a.dat", "b.dat"], dirs: ["subdir1", "subdir2"]}
		//   "subdir1": {files: ["c.bat"]}
		//
		// Also note this is probably not the cleanest or most efficient way to do
		// this.
		switch {
		case strings.HasPrefix(line, "$ cd"):
			// Save the current directory, we're done with it.
			if _, ok := dirs[curPath]; !ok && curPath != "" {
				dirs[curPath] = curDir
			}
			curDir = &dir{}

			intoDir := strings.Split(line, " ")[2]
			switch intoDir {
			case "/":
				curPath = intoDir
			case "..":
				pathParts := strings.Split(curPath, "/")
				curPath = strings.Join(pathParts[:len(pathParts)-1], "/")
				if curPath == "" {
					curPath += "/"
				}
			default:
				if curPath == "/" {
					curPath += intoDir
					break
				}
				curPath += "/" + intoDir
			}
		case strings.HasPrefix(line, "$ ls"):
			// Do nothing on this command.
		default:
			lineParts := strings.Split(line, " ")
			size, err := strconv.Atoi(lineParts[0])

			// This is a directory.
			if err != nil {
				curDir.dirs = append(curDir.dirs, lineParts[1])
				break
			}

			// This is a file.
			curDir.files = append(curDir.files, fil{
				name: lineParts[1],
				size: size,
			})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if _, ok := dirs[curPath]; !ok {
		dirs[curPath] = curDir
	}

	smallDirsTotal := 0
	usedSpace := size("/", dirs, &smallDirsTotal)
	fmt.Printf("Part 1: %d\n", smallDirsTotal)

	unusedSpace := 70000000 - usedSpace
	toDelSize := dirToDel(30000000-unusedSpace, dirs)
	fmt.Printf("Part 2: %d\n", toDelSize)
}

type fil struct {
	name string
	size int
}

type dir struct {
	files []fil
	dirs  []string
	size  int
}

// size recursively visits all the directories in the dirs map and calculates
// each dir's size. Also keeps track of the total size of all small dirs (those
// under 100000), for part 1.
func size(path string, dirs map[string]*dir, smallDirsTotal *int) int {
	dir := dirs[path]

	dirSize := 0
	for _, f := range dir.files {
		dirSize += f.size
	}

	for _, subDir := range dir.dirs {
		subPath := path + subDir
		if path != "/" {
			subPath = path + "/" + subDir
		}
		dirSize += size(subPath, dirs, smallDirsTotal)
	}

	if dirSize <= 100000 {
		*smallDirsTotal += dirSize
	}

	dir.size = dirSize

	return dirSize
}

// dirToDel finds the smallest dir to delete to give us over 30000000 of free
// space.
func dirToDel(needSpace int, dirs map[string]*dir) int {
	toDelSize := math.MaxInt
	for _, dir := range dirs {
		if dir.size >= needSpace {
			if dir.size < toDelSize {
				toDelSize = dir.size
			}
		}
	}
	return toDelSize
}
