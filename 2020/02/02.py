import sys


def parse_rule(rule):
    interval, char = rule.split()
    mn, mx = map(int, interval.split("-"))
    return mn, mx, char


def is_valid_pwd(s):
    rule, pwd = s.split(": ")
    mn, mx, char = parse_rule(rule)
    return mn <= pwd.count(char) <= mx


def is_valid_pwd2(s):
    rule, pwd = s.split(": ")
    a, b, char = parse_rule(rule)
    xorcnt = 0
    if pwd[a - 1] == char:
        xorcnt += 1
    if pwd[b - 1] == char:
        xorcnt += 1
    return xorcnt == 1


cnt = 0
cnt2 = 0
for line in sys.stdin:
    line = line.strip()
    if is_valid_pwd(line):
        cnt += 1
    if is_valid_pwd2(line):
        cnt2 += 1
print(cnt)
print(cnt2)
