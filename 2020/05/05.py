import sys

seats = set()
for line in sys.stdin:
    line = line.strip()
    row = int("".join(map(lambda c: "1" if c == "B" else "0", line[:-3])), 2)
    col = int("".join(map(lambda c: "1" if c == "R" else "0", line[-3:])), 2)
    seats.add(row * 8 + col)

mn, mx = min(seats), max(seats)
print(mx)
for i in range(mn + 1, mx):
    if i not in seats:
        print(i)
        break
