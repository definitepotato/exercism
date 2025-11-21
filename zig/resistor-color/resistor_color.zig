pub const ColorBand = enum {
    black,
    brown,
    red,
    orange,
    yellow,
    green,
    blue,
    violet,
    grey,
    white,
};

pub fn colorCode(color: ColorBand) usize {
    switch (color) {
        .black => return 0,
        .brown => return 1,
        .red => return 2,
        .orange => return 3,
        .yellow => return 4,
        .green => return 5,
        .blue => return 6,
        .violet => return 7,
        .grey => return 8,
        .white => return 9,
    }
}

pub fn colors() []const ColorBand {
    return &[_]ColorBand{
        .black,
        .brown,
        .red,
        .orange,
        .yellow,
        .green,
        .blue,
        .violet,
        .grey,
        .white,
    };
}
