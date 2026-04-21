from collections import Counter
import sys


def part1(jolts: list[int]):
    diffs = [jolts[i] - jolts[i - 1] for i in range(1, len(jolts))]
    cnts = Counter(diffs)
    print(cnts[1] * cnts[3])


def part2(jolts: list[int]):
    dp = [0] * len(jolts)
    dp[0] = 1
    for i in range(1, len(jolts)):
        for j in range(i - 1, -1, -1):
            if jolts[i] - jolts[j] > 3:
                break
            dp[i] += dp[j]
    print(dp[-1])


def solve():
    jolts = sorted([int(line.strip()) for line in sys.stdin])
    jolts = [0, *jolts, jolts[-1] + 3]
    part1(jolts)
    part2(jolts)


solve()
