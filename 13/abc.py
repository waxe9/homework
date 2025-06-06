from collections import Counter

with open('A.txt', 'r') as input_file_1:
    data_first = list(map(int, input_file_1.read().split()))

with open('B.txt', 'r') as input_file_2:
    data_second = list(map(int, input_file_2.read().split()))

frequency_first = Counter(data_first)
frequency_second = Counter(data_second)  
shared_numbers = set(data_first) & set(data_second)
output_list = []

for number in shared_numbers:
    highest_frequency = max(frequency_first[number], frequency_second[number])
    output_list.extend([number] * highest_frequency)


sorted_output = sorted(output_list)

with open('C.txt', 'w') as result_file:
    result_file.write('\n'.join(map(str, sorted_output)))