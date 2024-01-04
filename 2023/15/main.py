import sys

with open(sys.argv[1]) as f:
    sequence = [seq for seq in f.read().strip().split(",")]


def hash(string: str) -> int:
    s = 0
    for ch in string:
        s += ord(ch)
        s *= 17
        s %= 256
    return s


def solve_one(sequence: list[str]) -> int:
    return sum([hash(seq) for seq in sequence])


def solve_two(sequence: list[str]) -> int:
    boxes = [{} for _ in range(256)]
    for seq in sequence:
        if "-" in seq:
            label, _ = seq.split("-")
            i = hash(label)
            boxes[i].pop(label, None)
        if "=" in seq:
            label, length = seq.split("=")
            i = hash(label)
            boxes[i][label] = int(length)

    power = 0
    for i, box in enumerate(boxes):
        for j, k in enumerate(box):
            power += (i + 1) * (j + 1) * box[k]

    return power


print("[part 1]:", solve_one(sequence))
print("[part 2]:", solve_two(sequence))
