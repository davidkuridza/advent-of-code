package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Almanac struct {
	seeds []int
	maps  [][]Map
}

type Map struct {
	Destination int
	Source      int
	Range       int
}

func toInt(input string) int {
	n, _ := strconv.Atoi(input)
	return n
}

func parse(r io.Reader) Almanac {
	almanac := Almanac{
		seeds: []int{},
		maps:  [][]Map{},
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds: ") {
			for _, s := range strings.Split(line, " ")[1:] {
				almanac.seeds = append(almanac.seeds, toInt(s))
			}
			continue
		}

		if strings.HasSuffix(line, "map:") {
			almanac.maps = append(almanac.maps, []Map{})
			continue
		}

		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		almanac.maps[len(almanac.maps)-1] = append(almanac.maps[len(almanac.maps)-1], Map{
			Destination: toInt(parts[0]),
			Source:      toInt(parts[1]),
			Range:       toInt(parts[2]),
		})
	}

	return almanac
}

var (
	seeds = []int{}
	maps  = [][]Map{}
)

func main() {
	f, err := os.Open("input.txt.test")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	almanac := parse(f)
	seeds = almanac.seeds
	maps = almanac.maps

	locationOne := math.MaxInt
	for _, seed := range seeds {
		locationOne = min(locationOne, location(seed))
	}

	fmt.Println("[part 1] location:", locationOne, locationOne == 35)
	fmt.Println("[part 1] location:", locationOne, locationOne == 806029445)

	locationTwo := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		end := seeds[i] + seeds[i+1]
		for s := seeds[i]; s < end; s++ {
			locationTwo = min(locationTwo, location(s))
		}
	}

	fmt.Println("[part 2] location:", locationTwo, locationTwo == 46)
	fmt.Println("[part 2] location:", locationTwo, locationTwo == 59370572)
}

func location(seed int) int {
	for _, m := range maps {
		seed = closest(seed, m)
	}
	return seed
}

func closest(seed int, m []Map) int {
	for _, d := range m {
		if d.Source <= seed && seed < d.Source+d.Range {
			return d.Destination + seed - d.Source
		}
	}
	return seed
}
