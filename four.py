# Help from: https://zhaqenl.github.io/en/wordsearch/
import os.path

input = "four.txt"
search = []
count = 0


def coord_char(coord, matrix):
    row_index, column_index = coord

    return matrix[row_index][column_index]


def make_word(coord_matrix, matrix):
    return ''.join([coord_char(coord, matrix)
                   for coord in coord_matrix])


def find_base_match(char, matrix):
    base_matches = [(row_index, column_index)
                    for row_index, row in enumerate(matrix)
                    for column_index, column in enumerate(row)
                    if char == column]

    return base_matches


def matched_neighbors(coord, char, matrix, row_length, col_length):
    row_num, col_num = coord
    neighbors_coords = [(row, column)
                        for row in (row_num - 1, row_num + 1)
                        for column in (col_num - 1, col_num + 1)
                        if row_length > row >= 0
                        and col_length > column >= 0
                        and coord_char((row, column), matrix) == char
                        and not (row, column) == coord]

    return neighbors_coords


def complete_line(base_coord, targ_coord, row_length, col_len):
    line = [base_coord, targ_coord]
    diff1, diff2 = targ_coord[0] - base_coord[0], targ_coord[1] - base_coord[1]

    line += [(line[0][0] - diff1, line[0][1] - diff2)]

    if 0 <= line[-1][0] < row_length and 0 <= line[-1][1] < col_len:
        return line

    return []


def complete_match(word, matrix, base_match, row_len, col_len):
    new = (complete_line(base, n, row_len, col_len)
           for base in base_match
           for n in matched_neighbors(base, word[1], matrix, row_len, col_len))

    return [ero for ero in new
            if make_word(ero, matrix) == word]


if not os.path.isfile(input):
    print("Can't search for Words.")
else:
    with open(input) as f:
        for line in f:
            search.append(line.strip())


row_len, column_len = len(search), len(search[0])
base_matches = find_base_match('A', search)
matches = complete_match('ASM', search, base_matches, row_len, column_len)

for word in matches:
    for otherword in matches:
        if word is not otherword:
            if word[0] == otherword[0]:
                count += 1

print(int(count/2))
