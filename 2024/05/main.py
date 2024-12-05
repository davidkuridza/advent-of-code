import functools
import sys

with open(sys.argv[1]) as f:
    rules, updates = f.read().split("\n\n")
    updates = [update.split(",") for update in updates.split()]


def compare(a: str, b: str) -> int:
    if f"{a}|{b}" in rules:
        return -1
    if f"{b}|{a}" in rules:
        return 1
    return 0


def part_one() -> int:
    sum = 0

    for update in updates:
        ordered = sorted(update, key=functools.cmp_to_key(compare))
        if ordered == update:
            sum += int(ordered[len(ordered) // 2])

    return sum


def part_two() -> int:
    sum = 0

    for update in updates:
        ordered = sorted(update, key=functools.cmp_to_key(compare))
        if ordered != update:
            sum += int(ordered[len(ordered) // 2])

    return sum


print("[part 1]:", part_one())
print("[part 2]:", part_two())
