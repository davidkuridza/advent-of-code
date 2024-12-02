import sys


with open(sys.argv[1]) as f:
    lines = f.readlines()
    lines = [[int(n) for n in line.strip().split()] for line in lines]


def safe(values: list[int]) -> bool:
    return all(
        # between 0-4 and all increasing/decreasing
        0 < abs(a - b) < 4 and (0 < b - a) == (values[0] < values[1])
        for a, b in zip(values, values[1:])
    )


def solve_one(values: list[int]) -> bool:
    return safe(values)


def solve_two(values: list[int]) -> bool:
    if safe(values):
        return True

    for i in range(len(values)):
        removed = values.copy()
        removed.pop(i)
        if safe(removed):
            return True

    return False


sum_one = 0
sum_two = 0

for line in lines:
    sum_one += int(solve_one(line))
    sum_two += int(solve_two(line))

print("[part 1]:", sum_one)
print("[part 2]:", sum_two)
