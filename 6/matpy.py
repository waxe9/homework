def get_transposed_matrix(original_matrix):  # Функция транспонирования матрицы
    # Создаем новую матрицу, где строки становятся столбцами и наоборот
    return [[original_matrix[column][row] for column in range(len(original_matrix))] for row in
            range(len(original_matrix))]


def calculate_matrix_trace(square_matrix):  # Функция вычисления следа матрицы
    return sum(square_matrix[index][index] for index in range(len(square_matrix)))


def calculate_matrix_determinant(three_by_three_matrix):  # Функция вычисления определителя для матрицы 3x3
    # Разбиваем матрицу на отдельные элементы для удобства
    top_left, top_middle, top_right = three_by_three_matrix[0]
    middle_left, center, middle_right = three_by_three_matrix[1]
    bottom_left, bottom_middle, bottom_right = three_by_three_matrix[2]

    # Вычисляем определитель по формуле Саррюса (правило треугольника)
    return (top_left * center * bottom_right +
            top_middle * middle_right * bottom_left +
            top_right * middle_left * bottom_middle -
            top_right * center * bottom_left -
            top_middle * middle_left * bottom_right -
            top_left * middle_right * bottom_middle)


# Исходная матрица 3x3
input_matrix = [
    [1, 2, 3],
    [4, 5, 2],
    [7, 1, 9]
]

print("Исходная матрица:")
for matrix_row in input_matrix:
    print(matrix_row)

print("Транспонированная матрица:")
transposed_result = get_transposed_matrix(input_matrix)
for transposed_row in transposed_result:
    print(transposed_row)

print("След матрицы:", calculate_matrix_trace(input_matrix))
print("Определитель матрицы:", calculate_matrix_determinant(input_matrix))