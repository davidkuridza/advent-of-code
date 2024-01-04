package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int

func (d direction) step() int {
	if d == up || d == left {
		return -1
	}
	return 1
}

const (
	up direction = iota
	down
	left
	right
)

type tree struct {
	x, y       int
	height     int
	visibility map[direction]bool
	distance   map[direction]int
}

func (t *tree) lookAround(grid [][]tree) {
	t.visibility = map[direction]bool{
		up:    true,
		down:  true,
		left:  true,
		right: true,
	}
	t.distance = map[direction]int{
		up:    0,
		down:  0,
		left:  0,
		right: 0,
	}

	t.rows(grid, up)
	t.rows(grid, down)
	t.cols(grid, left)
	t.cols(grid, right)
}

func (t *tree) visible() bool {
	return t.visibility[up] || t.visibility[down] || t.visibility[left] || t.visibility[right]
}

func (t *tree) scenic() int {
	return t.distance[up] * t.distance[down] * t.distance[left] * t.distance[right]
}

func (t *tree) rows(grid [][]tree, direction direction) {
	x, y, step := t.x, t.y, direction.step()
	for y += step; y >= 0 && y < len(grid); y += step {
		t.distance[direction]++
		if t.height <= grid[y][x].height {
			t.visibility[direction] = false
			break
		}
	}
}

func (t *tree) cols(grid [][]tree, direction direction) {
	x, y, step := t.x, t.y, direction.step()
	for x += step; x >= 0 && x < len(grid[y]); x += step {
		t.distance[direction]++
		if t.height <= grid[y][x].height {
			t.visibility[direction] = false
			break
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := [][]tree{}
	for scanner.Scan() {
		t := []tree{}
		for x, height := range scanner.Text() {
			t = append(t, tree{x: x, y: len(grid), height: int(height - '0')})
		}
		grid = append(grid, t)
	}

	visible := len(grid)*2 + (len(grid[0])-2)*2 // edges
	scenic := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			tree := grid[y][x]
			tree.lookAround(grid)
			if tree.visible() {
				visible++
			}
			if tree.scenic() > scenic {
				scenic = tree.scenic()
			}
		}
	}

	fmt.Println("[part 1] visible:", visible)
	fmt.Println("[part 2] scenic:", scenic)
}
