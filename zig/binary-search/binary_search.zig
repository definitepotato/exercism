// Take a look at the tests, you might have to change the function arguments
const std = @import("std");

pub fn binarySearch(comptime T: type, target: usize, items: ?[]const T) ?usize {
    var buf = items.?;
    var left: usize = 0;

    while (buf.len > 0) {
        const idx = buf.len / 2;
        if (target == buf[idx]) return left + idx;

        if (target < buf[idx]) {
            buf = buf[0..idx];
        } else {
            buf = buf[idx + 1 ..];
            left += idx + 1;
        }
    }

    return null;
}
