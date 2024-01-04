package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type point struct {
	x, y int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gridOne := make(map[point]struct{}, 0)
	gridTwo := make(map[point]struct{}, 0)
	height := 0

	re := regexp.MustCompile("([0-9]+,[0-9]+)")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for i := 0; i < len(matches)-1; i++ {
			var from, to point
			fmt.Sscanf(matches[i][0], "%d,%d", &from.x, &from.y)
			fmt.Sscanf(matches[i+1][0], "%d,%d", &to.x, &to.y)
			if from.x == to.x {
				y := min(from.y, to.y)
				for y := y; y <= max(from.y, to.y); y++ {
					gridOne[point{x: from.x, y: y}] = struct{}{}
					gridTwo[point{x: from.x, y: y}] = struct{}{}
				}
				height = max(height, y)
			}
			if from.y == to.y {
				for x := min(from.x, to.x); x <= max(from.x, to.x); x++ {
					gridOne[point{x: x, y: from.y}] = struct{}{}
					gridTwo[point{x: x, y: from.y}] = struct{}{}
				}
				height = max(height, from.y)
			}
		}
	}

	fmt.Println("[part 1] grains:", drop(gridOne, height, false))
	fmt.Println("[part 2] grains:", drop(gridTwo, height, true))
}

func drop(grid map[point]struct{}, height int, partTwo bool) int {
	directions := []point{
		{x: 0, y: 1},  // down
		{x: -1, y: 1}, // diagonal left
		{x: 1, y: 1},  // diagonal right
	}
	grains := 0

	for {
		grain := point{x: 500, y: 0}

		if partTwo {
			if _, blocked := grid[grain]; blocked {
				return grains
			}
		}
		for {
			if grain.y+1 >= height+2 {
				if !partTwo {
					return grains
				}
				grid[grain] = struct{}{}
				grains++
				break
			}

			blocked := true
			for _, direction := range directions {
				candidate := point{x: grain.x + direction.x, y: grain.y + direction.y}
				if _, ok := grid[candidate]; !ok {
					grain = candidate
					blocked = false
					break
				}
			}

			if blocked {
				grid[grain] = struct{}{}
				grains++
				break
			}
		}
	}
}
