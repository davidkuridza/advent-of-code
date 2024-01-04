package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type monkey struct {
	name    string
	value   int
	monkey1 string
	monkey2 string
	op      byte
}

func parse() map[string]*monkey {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	monkeys := map[string]*monkey{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		monkey := monkey{name: line[0:4]}
		if len(line) == 17 { // name: name1 op name2
			monkey.op = line[11]
			monkey.monkey1 = line[6:10]
			monkey.monkey2 = line[13:17]
		} else { // name: number
			monkey.value, _ = strconv.Atoi(line[6:])
		}
		monkeys[monkey.name] = &monkey
	}
	return monkeys
}

func calculate(op byte, a, b int) int {
	operations := map[byte]func(a, b int) int{
		'+': func(a, b int) int { return a + b },
		'-': func(a, b int) int { return a - b },
		'*': func(a, b int) int { return a * b },
		'/': func(a, b int) int { return a / b },
	}
	return operations[op](a, b)
}

func partOne() int {
	monkeys := parse()
	for monkeys["root"].value == 0 {
		for _, monkey := range monkeys {
			if 0 < monkey.value {
				continue
			}
			v1 := monkeys[monkey.monkey1].value
			v2 := monkeys[monkey.monkey2].value
			if 0 < v1 && 0 < v2 {
				monkey.value = calculate(monkey.op, v1, v2)
			}
		}
	}

	return monkeys["root"].value
}

func partTwo() int {
	monkeys := parse()
	monkeys["humn"].value = 0
	root := monkeys["root"]
	for monkeys[root.monkey1].value == 0 && monkeys[root.monkey2].value == 0 {
		for _, monkey := range monkeys {
			if 0 < monkey.value || monkey.name == "humn" {
				continue
			}
			v1 := monkeys[monkey.monkey1].value
			v2 := monkeys[monkey.monkey2].value
			if 0 < v1 && 0 < v2 {
				monkey.value = calculate(monkey.op, v1, v2)
			}
		}
	}

	var monkey *monkey
	if monkeys[root.monkey1].value == 0 {
		monkey = monkeys[root.monkey1]
		monkeys[root.monkey1].value = monkeys[root.monkey2].value
	} else {
		monkey = monkeys[root.monkey2]
		monkeys[root.monkey2].value = monkeys[root.monkey1].value
	}

	for monkeys["humn"].value == 0 {
		m1 := monkeys[monkey.monkey1]
		m2 := monkeys[monkey.monkey2]
		if m1.value == 0 {
			switch monkey.op {
			case '+':
				m1.value = monkey.value - m2.value
			case '-':
				m1.value = monkey.value + m2.value
			case '*':
				m1.value = monkey.value / m2.value
			case '/':
				m1.value = monkey.value * m2.value
			}
			monkey = m1
		}
		if m2.value == 0 {
			switch monkey.op {
			case '+':
				m2.value = monkey.value - m1.value
			case '-':
				m2.value = m1.value - monkey.value
			case '*':
				m2.value = monkey.value / m1.value
			case '/':
				m2.value = m1.value / monkey.value
			}
			monkey = m2
		}
	}

	return monkeys["humn"].value
}

func main() {
	fmt.Println("[part 1] number:", partOne())
	fmt.Println("[part 2] number:", partTwo())
}
