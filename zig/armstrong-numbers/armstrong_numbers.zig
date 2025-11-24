const std = @import("std");
const pow = std.math.pow;

pub fn isArmstrongNumber(num: u128) bool {
    if (num == 0) return true;

    var len_num: u128 = 0;
    var acc = num;
    // get len of num
    while (acc > 0) : (acc /= 10) {
        len_num += 1;
    }

    var sum: u128 = 0;
    var my_num = num;
    while (my_num > 0) {
        sum += pow(u128, my_num % 10, len_num);
        my_num /= 10;
    }

    return sum == num;
}
