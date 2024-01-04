package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type device struct {
	cycle    int
	x        int
	strength int
	next     int
	display  []int
}

func (d *device) tick(increment int) {
	d.display = append(d.display, d.x)

	d.cycle++
	if d.next <= d.cycle {
		d.strength += d.cycle * d.x
		d.next += 40
	}
	d.x += increment
}

func (d *device) render() {
	for cycle, x := range d.display {
		pixel := cycle % 40
		if x-1 <= pixel && pixel <= x+1 {
			// FULL BLOCK
			// https://www.fileformat.info/info/unicode/char/2588/index.htm
			fmt.Print("\u2588")
		} else {
			fmt.Print(" ")
		}
		if (cycle+1)%40 == 0 {
			fmt.Println()
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	device := device{
		x:    1,
		next: 20,
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		device.tick(0)
		var inc int
		if _, err := fmt.Sscanf(scanner.Text(), "addx %d", &inc); err == nil {
			device.tick(inc)
		}
	}

	fmt.Println("[part 1] strength:", device.strength)
	fmt.Printf("[part 2] display:\n\n")

	device.render()
}
