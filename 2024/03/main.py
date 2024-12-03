import re
import sys

with open(sys.argv[1]) as f:
    line = f.read()


def solve_one(input: str) -> int:
    instructions = re.findall(r"mul\((\d+),(\d+)\)", input)
    return sum([int(a) * int(b) for a, b in instructions])


def solve_two(input: str) -> int:
    safe = True
    sum = 0
    pattern = re.compile(r"mul\((\d+),(\d+)\)|don't\(\)|do\(\)")

    for m in re.finditer(pattern, input):
        if m.group(0) == "do()":
            safe = True
        elif m.group(0) == "don't()":
            safe = False
        elif safe:
            sum += int(m.group(1)) * int(m.group(2))

    return sum


print("[part 1]:", solve_one(line))
print("[part 2]:", solve_two(line))
