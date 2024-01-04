package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation string
	divisible int
	throw     map[bool]int
	inspected int
}

func (m *monkey) parse(input string) {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile("[0-9]+")
	for _, item := range re.FindAllString(lines[1], -1) {
		i, _ := strconv.Atoi(item)
		m.items = append(m.items, i)
	}

	m.operation = strings.TrimPrefix(lines[2], "  Operation: new = ")
	fmt.Sscanf(lines[3], "  Test: divisible by %d", &m.divisible)

	var toMonkeyTrue, toMonkeyFalse int
	fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &toMonkeyTrue)
	fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &toMonkeyFalse)
	m.throw = map[bool]int{
		true:  toMonkeyTrue,
		false: toMonkeyFalse,
	}
}

func (m *monkey) inspect(item int, lcm int) (int, int) {
	operation := strings.ReplaceAll(m.operation, "old", strconv.Itoa(item))
	var x int
	var op string
	fmt.Sscanf(operation, "%d %s %d", &item, &op, &x)

	if op == "+" {
		item += x
	} else if op == "*" {
		item *= x
	}

	if lcm == 0 { // part 1
		item /= 3
	} else { // part 2
		item %= lcm
	}

	m.inspected++

	divisble := item%m.divisible == 0
	return item, m.throw[divisble]
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lcm := 1
	monkeysPartOne := []*monkey{}
	monkeysPartTwo := []*monkey{}
	for _, raw := range strings.Split(string(b), "\n\n") {
		m := monkey{}
		m.parse(raw)
		lcm *= m.divisible
		monkeysPartOne = append(monkeysPartOne, &m)
		m2 := m
		monkeysPartTwo = append(monkeysPartTwo, &m2)
	}

	fmt.Println("[part 1] monkey business:", business(monkeysPartOne, 20, 0))
	fmt.Println("[part 2] monkey business:", business(monkeysPartTwo, 10000, lcm))
}

func business(monkeys []*monkey, rounds int, lcm int) int {
	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				item, to := monkey.inspect(item, lcm)
				monkeys[to].items = append(monkeys[to].items, item)
				monkey.items = monkey.items[1:]
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[j].inspected < monkeys[i].inspected
	})

	return monkeys[0].inspected * monkeys[1].inspected
}
