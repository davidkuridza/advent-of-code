import collections
import sys

draws = [
    [1, 1, 1, 1, 1],  # 5 different
    [1, 1, 1, 2],  # 3 different, 2 same
    [1, 2, 2],  # 1 different, 2x2 same
    [1, 1, 3],  # 2x1 different, 3 same
    [2, 3],  # 2 same, 3 same
    [1, 4],  # 1 different, 4 same
    [5],  # 5 same
]

strength_one = list("23456789TJQKA")  # ['2', '3', ...]
strength_two = list("J23456789TQKA")  # ['J', '2', ...]


def type(hand: str) -> int:
    c = collections.Counter(hand)
    draw = sorted(c.values())
    return draws.index(draw)


def score(hand: str, strength: list[str], type: int) -> int:
    score = 0
    for i, v in enumerate(hand, start=1):
        score += (13 ** (5 - i)) * strength.index(v)

    return score + (13**5) * type


def winnings(hands: list) -> int:
    hands.sort(key=lambda x: x["score"])
    sum = 0
    for i, hand in enumerate(hands):
        sum += hand["bid"] * (i + 1)

    return sum


with open(sys.argv[1]) as f:
    lines = f.readlines()
    lines = [line.strip() for line in lines]


hands_one = []
hands_two = []
for line in lines:
    hand, bid = line.split(" ")
    # part one
    hands_one.append(
        {
            "bid": int(bid),
            "score": score(hand, strength_one, type(hand)),
        }
    )
    # part two
    highest_type = max([type(hand.replace("J", card)) for card in strength_two])
    hands_two.append(
        {
            "bid": int(bid),
            "score": score(hand, strength_two, highest_type),
        }
    )


print("[part 1] winnings:", winnings(hands_one))
print("[part 2] winnings:", winnings(hands_two))
