numbers = []
while True:
    line = input()
    if not line:
        break
    numbers.append(int(line))

for i in range(len(numbers)):
    for j in range(i + 1, len(numbers)):
        if i != j and numbers[i] + numbers[j] == 2020:
            print(numbers[i] * numbers[j])

for i in range(len(numbers)):
    for j in range(i, len(numbers)):
        for k in range(j, len(numbers)):
            if numbers[i] + numbers[j] + numbers[k] == 2020:
                print(numbers[i] * numbers[j] * numbers[k])
