const std = @import("std");

pub fn score(s: []const u8) u32 {
    var word_score: u32 = 0;

    for (s) |letter| {
        switch (std.ascii.toUpper(letter)) {
            'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T' => word_score += 1,
            'D', 'G' => word_score += 2,
            'B', 'C', 'M', 'P' => word_score += 3,
            'F', 'H', 'V', 'W', 'Y' => word_score += 4,
            'K' => word_score += 5,
            'J', 'X' => word_score += 8,
            'Q', 'Z' => word_score += 10,
            else => word_score += 0,
        }
    }

    return word_score;
}
