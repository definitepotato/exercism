const std = @import("std");
const print = std.debug.print;

/// Writes a reversed copy of `s` to `buffer`.
pub fn reverse(buffer: []u8, s: []const u8) []u8 {
    if (s.len > 0) {
        const end: usize = s.len - 1;
        for (0..s.len) |idx| {
            buffer[idx] = s[end - idx];
        }
    }
    return buffer;
}
