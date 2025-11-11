// Please implement the `ComputationError.IllegalArgument` error.

const std = @import("std");

pub const ComputationError = error{
    IllegalArgument,
};

pub fn steps(number: usize) anyerror!usize {
    if (number == 0) return ComputationError.IllegalArgument;

    var n = number;
    var n_steps: usize = 0;
    while (n > 0) {
        if (n == 1) return n_steps;

        if (@mod(n, 2) == 0) {
            n /= 2;
            n_steps += 1;
            continue;
        }

        if (@mod(n, 2) > 0) {
            n *= 3;
            n += 1;
            n_steps += 1;
            continue;
        }
    }

    return n_steps;
}
