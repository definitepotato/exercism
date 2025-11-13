import math

def is_armstrong_number(number):
    if number == 0: return True
    
    sum = 0
    og_num = number
    len_num = len(str(number))

    while number > 0:
        digit = number % 10
        sum = sum + (digit ** len_num)
        number = math.floor(number / 10)

    return sum == og_num

