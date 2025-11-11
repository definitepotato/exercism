const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;

pub fn eggCount(number: usize) usize {
    if (number == 0) return 0;

    var eggs: usize = 0;

    var i = number;
    while (i > 0) : (i = @divTrunc(i, 2)) {
        const rem = @mod(i, 2);
        eggs += rem;
    }

    return eggs;
}
