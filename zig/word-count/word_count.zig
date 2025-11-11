const std = @import("std");
const mem = std.mem;

/// Returns the counts of the words in `s`.
/// Caller owns the returned memory.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {
    var map = std.StringHashMap(u32).init(allocator);
    errdefer {
        var keys = map.keyIterator();
        while (keys.next()) |ptr| allocator.free(ptr.*);
        map.deinit();
    }

    var tok = mem.tokenizeAny(u8, s, " ,.:;!@$%^&\n");
    while (tok.next()) |item| {
        const cleaned_word = std.mem.trim(u8, item, "'\"");
        if (cleaned_word.len == 0) continue;

        const cleaned_and_lowered = try std.ascii.allocLowerString(allocator, cleaned_word);
        errdefer allocator.free(cleaned_and_lowered);

        const key = try map.getOrPut(cleaned_and_lowered);
        if (!key.found_existing) {
            key.value_ptr.* = 1;
            continue;
        }
        key.value_ptr.* += 1;
        allocator.free(cleaned_and_lowered);
    }

    return map;
}
