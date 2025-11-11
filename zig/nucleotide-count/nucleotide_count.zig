const std = @import("std");

pub const NucleotideError = error{Invalid};

pub const Counts = struct {
    a: u32,
    c: u32,
    g: u32,
    t: u32,
};

pub fn countNucleotides(s: []const u8) NucleotideError!Counts {
    var c = Counts{ .a = 0, .c = 0, .g = 0, .t = 0 };

    for (s) |nuc| {
        if (nuc != 'A' and nuc != 'C' and nuc != 'G' and nuc != 'T') {
            return NucleotideError.Invalid;
        }

        if (nuc == 'A') c.a += 1;
        if (nuc == 'C') c.c += 1;
        if (nuc == 'G') c.g += 1;
        if (nuc == 'T') c.t += 1;
    }

    return c;
}
