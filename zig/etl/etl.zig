const std = @import("std");
const mem = std.mem;

pub fn transform(allocator: mem.Allocator, legacy: std.AutoHashMap(i5, []const u8)) mem.Allocator.Error!std.AutoHashMap(u8, i5) {
    var map = std.AutoHashMap(u8, i5).init(allocator);

    var it = legacy.iterator();
    while (it.next()) |item| {
        const score = item.key_ptr.*;

        for (item.value_ptr.*) |value| {
            try map.put(std.ascii.toLower(value), score);
        }
    }

    return map;
}
