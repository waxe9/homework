from collections import deque  # Используем deque для эффективной работы с очередью


def load_labyrinth_data(file_path):
    """
    Загружает данные лабиринта из файла
    Возвращает двумерный список (матрицу) лабиринта
    """
    with open(file_path, "r") as input_file:
        labyrinth_map = [list(line.strip()) for line in input_file]
    return labyrinth_map


def locate_entrance_and_exit(labyrinth_map):
    """
    Находит координаты входа (S) и выхода (E) в лабиринте
    Возвращает кортеж ((row_start, col_start), (row_end, col_end))
    """
    entrance_coords = exit_coords = None
    # Перебираем все ячейки лабиринта
    for row_idx in range(len(labyrinth_map)):
        for col_idx in range(len(labyrinth_map[row_idx])):
            if labyrinth_map[row_idx][col_idx] == 'S':
                entrance_coords = (row_idx, col_idx)
            elif labyrinth_map[row_idx][col_idx] == 'E':
                exit_coords = (row_idx, col_idx)
    return entrance_coords, exit_coords


def find_shortest_path(labyrinth_map, start_point, end_point):
    """
    Поиск кратчайшего пути в лабиринте с использованием BFS
    Возвращает список координат пути или None, если путь не найден
    """
    path_queue = deque()  # Очередь для хранения путей
    path_queue.append((start_point, [start_point]))
    visited_cells = set()  # Множество посещенных клеток
    visited_cells.add(start_point)

    while path_queue:
        current_cell, current_path = path_queue.popleft()
        curr_row, curr_col = current_cell

        # Если достигли конечной точки, возвращаем путь
        if current_cell == end_point:
            return current_path

        # Проверяем все возможные направления движения
        for row_step, col_step in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            new_row = curr_row + row_step
            new_col = curr_col + col_step

            # Проверяем, что новые координаты допустимы
            if (0 <= new_row < len(labyrinth_map)) and \
                    (0 <= new_col < len(labyrinth_map[0])) and \
                    labyrinth_map[new_row][new_col] != '#' and \
                    (new_row, new_col) not in visited_cells:
                visited_cells.add((new_row, new_col))
            # Добавляем новую точку в очередь с обновленным путем
            path_queue.append(((new_row, new_col), current_path +[(new_row, new_col)]))

    return None  # Если путь не найден


def visualize_solution_path(labyrinth_map, solution_path):
    # Отмечает путь в лабиринте символами '*' и возвращает количество шагов
    step_counter = -1  # Счетчик шагов (начальная точка не считается шагом)
    for (row, col) in solution_path:
        step_counter += 1
        if labyrinth_map[row][col] not in ('S', 'E'):
            labyrinth_map[row][col] = '*'
    return step_counter


def convert_labyrinth_to_text(labyrinth_map):
    # Преобразует лабиринт в текстовое представление
    return '\n'.join(''.join(row) for row in labyrinth_map)


def save_solution_results(output_path, labyrinth_map, total_steps):
    # Сохраняет результаты решения в файл
    with open(output_path, 'w') as output_file:
        output_file.write(f"Steps: {total_steps}\n")
        output_file.write(convert_labyrinth_to_text(labyrinth_map) + '\n')


def solve_labyrinth(input_file_path, output_file_path="output.txt"):
    # Основная функция для решения лабиринта
    # Загрузка лабиринта
    labyrinth_map = load_labyrinth_data(input_file_path)

    # Поиск начальной и конечной точек
    start_point, end_point = locate_entrance_and_exit(labyrinth_map)
    if not start_point or not end_point:
        print("Не найдена начальная или конечная точка.")
        return

    # Поиск кратчайшего пути
    solution_path = find_shortest_path(labyrinth_map, start_point, end_point)

    if solution_path:
        # Если путь найден - визуализируем и сохраняем
        total_steps = visualize_solution_path(labyrinth_map, solution_path)
        map_text = convert_labyrinth_to_text(labyrinth_map)
        print(f"Количество шагов: {total_steps}")
        print(map_text)
        save_solution_results(output_file_path, labyrinth_map, total_steps)
    else:
        # Если путь не найден
        print("Путь не найден.")
        with open(output_file_path, 'w') as output_file:
            output_file.write("Путь не найден.\n")


if __name__ == "__main__":
    solve_labyrinth("input.txt")  # Запуск программы