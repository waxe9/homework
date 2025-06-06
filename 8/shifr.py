from string import ascii_uppercase

# Переименованные переменные:
original_text = (open("input.txt").readline()).upper()
shift_value = int(input())

def caesar_cipher(shift, text):
    """
    Функция реализует шифр Цезаря с заданным сдвигом.
    Каждая буква в тексте заменяется на букву, находящуюся на 'shift' позиций дальше в алфавите.
    """
    encrypted_result = ''
    alphabet = ascii_uppercase
    for char in text:
        encrypted_result += alphabet[(alphabet.index(char) + shift) % len(alphabet)]
    return encrypted_result

def atbash_cipher(text):
    """
    Функция реализует шифр Атбаш.
    Каждая буква заменяется на букву, симметричную ей в алфавите
    """
    alphabet = ascii_uppercase
    return text.translate(str.maketrans(alphabet + alphabet.upper(), alphabet[::-1] + alphabet.upper()[::-1]))

# Вывод результатов:
print(atbash_cipher(original_text))
print(caesar_cipher(shift_value, original_text))