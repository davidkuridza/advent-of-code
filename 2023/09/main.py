import sys


def solve(line: list) -> int:
    if not any(line):
        return 0

    next = [b - a for a, b in zip(line, line[1:])]
    return line[-1] + solve(next)


with open(sys.argv[1]) as f:
    lines = [line.strip() for line in f.readlines()]
    numbers = [[int(i) for i in line.split()] for line in lines]


part_one = sum(solve(n) for n in numbers)
print("[part 1] sum:", part_one)

part_two = sum(solve(n[::-1]) for n in numbers)
print("[part 2] sum:", part_two)
