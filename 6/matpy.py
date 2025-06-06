def transpose(matrix):
    return [[matrix[j][i] for j in range(len(matrix))] for i in range(len(matrix))]

def trace(matrix):
    return sum(matrix[i][i] for i in range(len(matrix)))

def determinant(matrix):
    a, b, c = matrix[0]
    d, e, f = matrix[1]
    g, h, i = matrix[2]
    return a*e*i + b*f*g + c*d*h - c*e*g - b*d*i - a*f*h

matrix = [
    [1, 2, 3],
    [4, 5, 2],
    [7, 1, 9]
]

print("Исходная матрица:")
for row in matrix:
    print(row)

print("Транспонированная матрица:")
transposed = transpose(matrix)
for row in transposed:
    print(row)

print("След матрицы:", trace(matrix))
print("Определитель матрицы:", determinant(matrix))