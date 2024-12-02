import os.path

lists = "one.txt"
listOne = []
listTwo = []
diff = []
similarity = []

if not os.path.isfile(lists):
    print("List don't exist.")
else:
    with open(lists) as f:
        for l in f:
            listOne.append(l.split()[0])
            listTwo.append(l.split()[1])

listOne.sort()
listTwo.sort()

for i in range(len(listOne)):
    diff.append(abs(int(listOne[i]) - int(listTwo[i])))
    similarity.append(int(listOne[i]) * listTwo.count(listOne[i]))

print(sum(similarity))
print(sum(diff))