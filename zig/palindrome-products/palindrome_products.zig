const std = @import("std");
const mem = std.mem;

pub fn isPalindrome(n: usize) bool {
    var buf: [20]u8 = undefined;
    const int_as_str = std.fmt.bufPrint(&buf, "{}", .{n}) catch unreachable;

    var left: usize = 0;
    var right: usize = int_as_str.len - 1;

    while (left < right) {
        if (int_as_str[left] != int_as_str[right]) return false;
        left += 1;
        right -= 1;
    }

    return true;
}

pub const Pair = struct {
    first: u64,
    second: u64,
};

pub const Palindrome = struct {
    value: u64,
    factors: []Pair,
};

pub fn smallest(allocator: mem.Allocator, min: u64, max: u64) mem.Allocator.Error!?Palindrome {
    _ = allocator;
    _ = min;
    _ = max;
    @compileError("please implement the smallest function");
}

pub fn largest(allocator: mem.Allocator, min: u64, max: u64) mem.Allocator.Error!?Palindrome {
    _ = allocator;
    _ = min;
    _ = max;
    @compileError("please implement the largest function");
}
