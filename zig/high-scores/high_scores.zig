const std = @import("std");

pub const HighScores = struct {
    // This struct, as well as its fields and methods, needs to be implemented.
    scores: []const i32,

    pub fn init(scores: []const i32) HighScores {
        return .{
            .scores = scores,
        };
    }

    pub fn latest(self: *const HighScores) ?i32 {
        if (self.scores.len == 0) return null;

        return self.scores[self.scores.len - 1];
    }

    pub fn personalBest(self: *const HighScores) ?i32 {
        if (self.scores.len == 0) return null;

        var highest: i32 = 0;
        for (self.scores) |score| {
            if (score > highest) highest = score;
        }

        return highest;
    }

    pub fn personalTopThree(self: *const HighScores) []const i32 {
        const allocator = std.heap.page_allocator;
        var sorted_scores = allocator.alloc(i32, self.scores.len) catch unreachable;

        @memcpy(sorted_scores, self.scores);
        std.mem.sort(i32, sorted_scores, {}, std.sort.desc(i32));

        return sorted_scores[0..@min(3, self.scores.len)];
    }
};
