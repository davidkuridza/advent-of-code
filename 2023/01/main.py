import re


with open("input.txt") as f:
    lines = f.readlines()
    lines = [line.strip() for line in lines]


def digit(line):
    matches = re.findall(r"\d", line)
    return int(matches[0] + matches[-1])


replace = {
    "one": "one1one",
    "two": "two2two",
    "three": "three3three",
    "four": "four4four",
    "five": "five5five",
    "six": "six6six",
    "seven": "seven7seven",
    "eight": "eight8eight",
    "nine": "nine9nine",
}

sumOne = 0
sumTwo = 0
for line in lines:
    # part 1
    sumOne += digit(line)

    # part 2
    # replace to avoid overlapping values, for example:
    # eighthree -> should be 83, not 8hree
    for k, v in replace.items():
        line = line.replace(k, v)

    sumTwo += digit(line)


print("[part 1] sum:", sumOne)
print("[part 2] sum:", sumTwo)
