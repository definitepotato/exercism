pub const ChessboardError = error{
    IndexOutOfBounds,
};

pub fn square(index: usize) ChessboardError!u64 {
    if (index == 1) return 1;
    if (index == 0) return ChessboardError.IndexOutOfBounds;
    if (index < 0 or index > 64) return ChessboardError.IndexOutOfBounds;

    var grains: usize = 1;
    for (0..index - 1) |_| grains *= 2;
    return grains;
}

pub fn total() u64 {
    return ~@as(u64, 0); // ~ flips all bits, this gives you max u64 without std import.
}
