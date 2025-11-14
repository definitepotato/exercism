const std = @import("std");
const allocator = std.heap.page_allocator;

const Reply = enum {
    Whatever,
    Sure,
    Chill,
    CalmDown,
    Fine,
};

fn isSilent(s: []const u8) bool {
    return s.len == 0;
}

fn isShouting(s: []const u8) bool {
    const s_upper = std.ascii.allocUpperString(allocator, s) catch unreachable;
    defer allocator.free(s_upper);
    const all_upper = std.mem.eql(u8, s_upper, s);

    var has_letters = false;
    for (s) |c| {
        if (std.ascii.isAlphabetic(c)) has_letters = true;
    }

    return all_upper and has_letters;
}

fn isQuestion(s: []const u8) bool {
    return s[s.len - 1] == '?';
}

fn parseResponse(s: []const u8) Reply {
    if (isSilent(s)) return Reply.Fine; // silent
    if (isShouting(s) and isQuestion(s)) return Reply.CalmDown; // exasperated
    if (isShouting(s)) return Reply.Chill; // shouting
    if (isQuestion(s)) return Reply.Sure; // question
    return Reply.Whatever;
}

pub fn response(s: []const u8) []const u8 {
    const trimmed = std.mem.trim(u8, s, " \t\r\n");
    const res = parseResponse(trimmed);

    switch (res) {
        .Sure => return "Sure.",
        .Chill => return "Whoa, chill out!",
        .CalmDown => return "Calm down, I know what I'm doing!",
        .Fine => return "Fine. Be that way!",
        .Whatever => return "Whatever.",
    }
}
