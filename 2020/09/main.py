import sys
from collections import deque
from typing import Sequence


def check_sum(buffer: deque, val: int) -> bool:
    return val in [a + b for a in buffer for b in buffer]


def invalid_number(nums: list[int]) -> int | None:
    period = 25
    buffer = deque(maxlen=period)
    for num in nums:
        if len(buffer) == period and not check_sum(buffer, num):
            return num
        buffer.append(num)


def find_range(nums: list[int], target: int) -> Sequence[int] | None:
    for left, _ in enumerate(nums):
        for right, _ in enumerate(nums):
            if sum(nums[left:right]) == target:
                return nums[left:right]


def solve():
    nums = [int(line.strip()) for line in sys.stdin]
    inv = invalid_number(nums)
    if not inv:
        return
    print(inv)
    r = find_range(nums, inv)
    if not r:
        return
    print(min(r) + max(r))


solve()
