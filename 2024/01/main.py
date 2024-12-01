import sys


with open(sys.argv[1]) as f:
    lines = f.readlines()
    lines = [line.strip() for line in lines]

left = []
right = []


for line in lines:
    l, r = line.split()
    left.append(int(l))
    right.append(int(r))


def solve_one(left: list[int], right: list[int]) -> int:
    left.sort()
    right.sort()
    return sum([abs(v[0] - v[1]) for v in zip(left, right)])


def solve_two(left: list[int], right: list[int]) -> int:
    return sum(v * right.count(v) for v in left)


print("[part 1]:", solve_one(left, right))
print("[part 2]:", solve_two(left, right))
