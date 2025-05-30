# Импортируем класс Counter из модуля collections для подсчета частот элементов
from collections import Counter

# Открываем первый файл для чтения и загружаем числа
with open('A.txt', 'r') as input_file_1:
    # Читаем весь файл, разбиваем на отдельные строки/числа и преобразуем в int
    data_first = list(map(int, input_file_1.read().split()))

# Аналогично открываем и обрабатываем второй файл
with open('B.txt', 'r') as input_file_2:
    data_second = list(map(int, input_file_2.read().split()))

# Создаем словари с частотой встречаемости каждого числа в обоих файлах
frequency_first = Counter(data_first)  # Частоты чисел из первого файла
frequency_second = Counter(data_second)  # Частоты чисел из второго файла

# Находим пересечение множеств - числа, присутствующие в обоих файлах
shared_numbers = set(data_first) & set(data_second)

# Инициализируем список для хранения результата
output_list = []

# Для каждого общего числа определяем, сколько раз его нужно добавить в результат
for number in shared_numbers:
    # Берем максимальное количество вхождений из двух файлов
    highest_frequency = max(frequency_first[number], frequency_second[number])
    # Добавляем число в результирующий список нужное количество раз
    output_list.extend([number] * highest_frequency)

# Сортируем итоговый список по возрастанию
sorted_output = sorted(output_list)

# Записываем результат в файл, каждое число на новой строке
with open('C.txt', 'w') as result_file:
    # Преобразуем числа в строки и объединяем через перенос строки
    result_file.write('\n'.join(map(str, sorted_output)))