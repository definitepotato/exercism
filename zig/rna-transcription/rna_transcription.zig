const std = @import("std");
const mem = std.mem;

pub fn toRna(allocator: mem.Allocator, dna: []const u8) mem.Allocator.Error![]const u8 {
    const buffer = try allocator.alloc(u8, 1024);
    defer allocator.free(buffer); // 1: this runs before return so we impl #2.

    var writer: std.Io.Writer = .fixed(buffer);

    for (dna) |strand| {
        if (strand == 'G') writer.writeAll("C") catch unreachable;
        if (strand == 'C') writer.writeAll("G") catch unreachable;
        if (strand == 'T') writer.writeAll("A") catch unreachable;
        if (strand == 'A') writer.writeAll("U") catch unreachable;
    }

    // 2: we dupe before the .free(buffer) happens before return, the caller owns the allocation.
    const result = try allocator.dupe(u8, buffer[0..writer.end]);
    return result;
}
