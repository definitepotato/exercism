const std = @import("std");
const pow = std.math.pow;

const number = struct {
    n: u128,

    pub fn number_of_digits(self: number) u128 {
        var num = self.n;
        var count: u128 = 0;

        while (num != 0) : (count += 1) {
            num /= 10;
        }

        return count;
    }

    pub fn is_armstrong(self: number) bool {
        var num = self.n;
        var sum: u128 = 0;
        const exponent = self.number_of_digits();

        while (num != 0) : (num /= 10) {
            sum += pow(u128, num % 10, exponent);
        }

        return sum == self.n;
    }
};

pub fn isArmstrongNumber(num: u128) bool {
    var n = number{
        .n = num,
    };

    return n.is_armstrong();
}
