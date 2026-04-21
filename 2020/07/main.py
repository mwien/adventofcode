import sys


def parse(line: str) -> tuple[str | None, list[tuple[str, int]]]:
    tokens = line.split()
    buffer = []

    container = None
    content = []
    for token in tokens:
        if token.startswith("bags") or token.startswith("bag"):
            if len(buffer) == 2:
                # this is the initial bag description
                container = " ".join(buffer)
            else:
                if buffer[-2] == "no" and buffer[-1] == "other":
                    continue
                # these are other bags
                color = " ".join(buffer[-2:])
                count = int(buffer[-3])
                content.append((color, count))
            buffer.clear()
        else:
            buffer.append(token)
    return container, content


graph = {}
for line in sys.stdin:
    line = line.strip()
    container, content = parse(line)
    graph[container] = content

cnt = 0
for key in graph:
    visited = set([key])
    stack = [key]
    while stack:
        u = stack.pop()
        if u == "shiny gold" and len(visited) > 1:
            cnt += 1
            break
        for v, _ in graph[u]:
            if v not in visited:
                visited.add(v)
                stack.append(v)
print(cnt)

cnt = 0
stack = [("shiny gold", 1)]
while stack:
    u, uc = stack.pop()
    for v, vc in graph[u]:
        stack.append((v, uc * vc))
        cnt += uc * vc
print(cnt)
