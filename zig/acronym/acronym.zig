const std = @import("std");
const mem = std.mem;

pub fn abbreviate(allocator: mem.Allocator, words: []const u8) mem.Allocator.Error![]u8 {
    var arr = std.array_list.Managed(u8).init(allocator);

    var lines = mem.tokenizeAny(u8, words, " -");

    while (lines.next()) |word| {
        for (word) |c| {
            if (std.ascii.isAlphabetic(c)) {
                try arr.append(std.ascii.toUpper(c));
                break;
            }
        }
    }

    return arr.toOwnedSlice();
}
