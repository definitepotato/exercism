def square(number):
    if number == 1: return 1
    if number == 0: raise ValueError("square must be between 1 and 64")
    if number < 0 or number > 64: raise ValueError("square must be between 1 and 64")

    n_grains = 1
    for square in range(number-1): n_grains *= 2
    return n_grains


def total():
    return square(64) * 2 - 1
