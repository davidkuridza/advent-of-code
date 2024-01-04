package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y  int
	value rune
}

type grid struct {
	points [][]point
	width  int
	height int
}

func (g *grid) neighbours(p point) []point {
	n := []point{}
	if 0 <= p.x-1 { // left: {-1, 0}
		n = append(n, g.points[p.y][p.x-1])
	}
	if p.x+1 < g.width { // right: {1, 0}
		n = append(n, g.points[p.y][p.x+1])
	}
	if 0 <= p.y-1 { // up: {0, -1}
		n = append(n, g.points[p.y-1][p.x])
	}
	if p.y+1 < g.height { // down: {0, 1}
		n = append(n, g.points[p.y+1][p.x])
	}
	return n
}

func (g *grid) path(from, to point) int {
	var (
		distance = map[point]int{}
		visited  = map[point]bool{}
		queue    = []point{from}
	)

	for 0 < len(queue) {
		current := queue[0]
		queue = queue[1:]

		if ok := visited[current]; ok {
			continue
		}
		visited[current] = true

		for _, neighbour := range g.neighbours(current) {
			if current == to {
				return distance[current]
			}
			if neighbour.value-current.value > 1 {
				continue
			}
			distance[neighbour] = distance[current] + 1
			queue = append(queue, neighbour)
		}
	}

	return -1
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	grid := grid{
		points: make([][]point, len(lines)),
		width:  len(lines[0]),
		height: len(lines),
	}

	var start []point
	var end point

	for row, line := range lines {
		grid.points[row] = make([]point, len(line))
		for col, r := range line {
			p := point{x: col, y: row, value: r}
			switch r {
			case 'S':
				p.value = 'a'
				start = append(start, p)
			case 'E':
				p.value = 'z'
				end = p
			case 'a':
				start = append(start, p)
			}
			grid.points[row][col] = p
		}
	}

	steps := grid.path(start[0], end)
	fmt.Println("[part 1] steps:", steps)

	least := steps
	for _, p := range start {
		steps := grid.path(p, end)
		if -1 < steps && steps < least {
			least = steps
		}
	}

	fmt.Println("[part 2] steps:", least)
}
