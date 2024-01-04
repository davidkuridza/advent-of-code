import re

with open("input.txt") as f:
    lines = f.readlines()
    lines = [line.strip() for line in lines]


def extract_numbers(input: str) -> dict[int]:
    return {int(n) for n in re.findall("\d+", input)}


part_one = 0
part_two = dict.fromkeys(range(0, len(lines)), 0)

for i, line in enumerate(lines):
    first, last = line.split("|")
    winning = extract_numbers(first.split(":")[1])
    numbers = extract_numbers(last)
    same = len(winning & numbers)
    if 0 < same:
        part_one += 2 ** (same - 1)

    part_two[i] += 1
    for j in range(same):
        part_two[i + j + 1] += part_two[i]

print("[part 1] points:", part_one)
print("[part 2] total: ", sum(part_two.values()))
