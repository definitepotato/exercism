const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;

pub fn isIsogram(str: []const u8) bool {
    const allocator = std.heap.page_allocator;

    var char_map = std.AutoHashMap(u8, void).init(allocator);
    defer char_map.deinit();

    for (str) |c| {
        if (std.ascii.isAlphabetic(c)) {
            const ch = std.ascii.toLower(c);

            if (char_map.get(ch)) |_| {
                return false;
            }

            char_map.put(ch, {}) catch unreachable;
        }
    }
    return true;
}
