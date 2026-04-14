import sys


def countSlope(lines, right, down):
    cnt = 0
    row = 0
    col = 0
    while True:
        if row >= len(lines):
            break
        if lines[row][col % len(lines[row])] == "#":
            cnt += 1
        col += right
        row += down
    return cnt


lines = []
for line in sys.stdin:
    lines.append(line.strip())

print(countSlope(lines, 3, 1))

res = 1
slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
for slope in slopes:
    res *= countSlope(lines, *slope)
print(res)
