package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	ORDERED = iota
	UNORDERED
	EQUAL
)

type packet struct {
	divider bool
	packets []packet
	value   int
}

func (pd *packet) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] == '[' {
		var packets []packet
		if err := json.Unmarshal(data, &packets); err != nil {
			return err
		}
		*pd = packet{packets: packets, value: -1}
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	*pd = packet{value: i}
	return nil
}

func parse(line string) *packet {
	v := (*packet)(nil)
	if err := json.Unmarshal([]byte(line), &v); err != nil {
		panic(err)
	}
	return v
}

func cmp(left, right int) int {
	if left == right {
		return EQUAL
	}
	if left < right {
		return ORDERED
	}
	return UNORDERED
}

func compare(left, right *packet) int {
	if left.value != -1 {
		left.packets = []packet{{value: left.value}}
	}
	if right.value != -1 {
		right.packets = []packet{{value: right.value}}
	}

	if left.value != -1 && right.value != -1 {
		return cmp(left.value, right.value)
	}

	var i, j int
	for i < len(left.packets) && j < len(right.packets) {
		if c := compare(&left.packets[i], &right.packets[i]); c != EQUAL {
			return c
		}
		i++
		j++
	}

	return cmp(len(left.packets), len(right.packets))
}

func partOne(packets [][]*packet) int {
	indices := 0
	for i, pairs := range packets {
		if compare(pairs[0], pairs[1]) == ORDERED {
			indices += i + 1
		}
	}
	return indices
}

func partTwo(packets []*packet) int {
	dividers := []string{"[[2]]", "[[6]]"}
	for _, divider := range dividers {
		p := parse(divider)
		p.divider = true
		packets = append(packets, p)
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == ORDERED
	})

	indices := 1
	for i, p := range packets {
		if p.divider {
			indices *= (i + 1)
		}
	}
	return indices
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	packetsPartOne := [][]*packet{}
	packetsPartTwo := []*packet{}

	lines := strings.Split(string(b), "\n\n")
	for _, line := range lines {
		pairs := strings.Split(line, "\n")
		p1 := parse(pairs[0])
		p2 := parse(pairs[1])
		packetsPartOne = append(packetsPartOne, []*packet{p1, p2})
		packetsPartTwo = append(packetsPartTwo, []*packet{p1, p2}...)
	}

	fmt.Println("[part 1] indices:", partOne(packetsPartOne))
	fmt.Println("[part 2] indices:", partTwo(packetsPartTwo))
}
