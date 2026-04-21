import sys
from string import ascii_lowercase
from typing import List


def cnt_any(group: List[str]):
    answers = set()
    for person in group:
        answers = answers.union(person)
    return len(answers)


def cnt_all(group: List[str]):
    answers = set(ascii_lowercase)
    for person in group:
        answers = answers.intersection(person)
    return len(answers)


def solve():
    groups = [[]]
    for line in sys.stdin:
        line = line.strip()
        if line:
            groups[-1].append(line)
        else:
            groups.append([])
    sum_any = 0
    sum_all = 0
    for group in groups:
        sum_any += cnt_any(group)
        sum_all += cnt_all(group)
    print(sum_any)
    print(sum_all)


solve()
