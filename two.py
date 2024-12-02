import os.path

reports = "two.txt"
safeCount = 0


def is_safe(report):
    temp = []
    safe = True
    direction = None

    for index, item in enumerate(report):
        if index == 0:
            temp.append(item)
        else:
            if temp[-1] > item:
                if direction is None:
                    direction = "descending"
                elif direction == "ascending":
                    safe = False

            elif temp[-1] < item:
                if direction is None:
                    direction = "ascending"
                elif direction == "descending":
                    safe = False

            temp.append(item)

            if abs(item - temp[index-1]) > 3:
                safe = False
            elif abs(item - temp[index-1]) == 0:
                safe = False

    return safe


if not os.path.isfile(reports):
    print("Reports aren't here!")
else:
    with open(reports) as f:
        for line in f:
            report = list(map(int, line.split()))
            if is_safe(report):
                safeCount += 1
            else:
                for i in range(len(report)):
                    dampener = list(report)
                    dampener.pop(i)
                    if is_safe(dampener):
                        safeCount += 1
                        break

print(safeCount)
