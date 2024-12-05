import os.path
import re

input = "five.txt"
sum = 0
correctUpdates = []
incorrectUpdates = []
rules = {}

if not os.path.isfile(input):
    print("Print Order don't exist.")
else:
    with open(input) as f:
        text = f.read()

rulesPattern = re.compile(r"\d{2}\|\d{2}")

orders = rulesPattern.findall(text)
updates = [update.split(",")
           for update in text.split("\n\n", 1)[1].splitlines()]

for order in orders:
    k, v = order.split("|")[0], order.split("|")[1]
    if k not in rules:
        rules[k] = [v]
    else:
        rules[k].append(v)

for update in updates:
    correct = True
    for i, n in enumerate(update):
        if i == 0:
            pass
        else:
            for j in range(i):
                if n in rules and update[j] in rules.get(n):
                    correct = False
                    update[i], update[j] = update[j], update[i]

    if correct:
        correctUpdates.append(update)
    else:
        incorrectUpdates.append(update)


# for update in correctUpdates:
#     sum += int(update[int(len(update)/2)])

for update in incorrectUpdates:
    sum += int(update[int(len(update)/2)])

print(sum)
