import os.path
import re

input = "three.txt"
sum = 0
enabled = True

if not os.path.isfile(input):
    print("Where's the Input?")
else:
    with open(input) as f:
        memory = f.read()

pattern = re.compile(r"mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)")
filter = pattern.findall(memory)

for operation in filter:
    if operation == "don't()":
        enabled = False
    elif operation == "do()":
        enabled = True
    else:
        if enabled:
            operands = list(map(int, re.findall(r"\d{1,3}", operation)))
            product = 1
            for num in operands:
                product *= num
            sum += product

print(sum)
