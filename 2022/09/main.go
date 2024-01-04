package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

type point struct {
	x, y int
}

func (p *point) move(to point, knot bool) {
	if knot {
		p.x += sign(to.x)
		p.y += sign(to.y)
	} else {
		p.x += to.x
		p.y += to.y
	}
}

func (p *point) knot(to point) point {
	return point{x: to.x - p.x, y: to.y - p.y}
}

var directions = map[string]point{
	"U": {x: 0, y: 1},
	"D": {x: 0, y: -1},
	"L": {x: -1, y: 0},
	"R": {x: 1, y: 0},
}

type move struct {
	direction string
	steps     int
}

func (m *move) point() point {
	return directions[m.direction]
}

func solve(moves []move, length int) int {
	positions := make([]point, length)
	visited := map[point]bool{}

	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			positions[0].move(m.point(), false)
			for j := 1; j < length; j++ {
				knot := positions[j].knot(positions[j-1])
				if abs(knot.x) > 1 || abs(knot.y) > 1 {
					positions[j].move(knot, true)
				}
			}
			visited[positions[length-1]] = true
		}
	}

	return len(visited)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	moves := []move{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var m move
		fmt.Sscanf(scanner.Text(), "%s %d", &m.direction, &m.steps)
		moves = append(moves, m)
	}

	fmt.Println("[part 1] positions:", solve(moves, 2))
	fmt.Println("[part 2] positions:", solve(moves, 10))
}
