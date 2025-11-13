pub const BufferError = error{BufferOverflow};

pub fn CircularBuffer(comptime T: type, comptime capacity: usize) type {
    return struct {
        buffer: [capacity]T = undefined,
        count: usize = 0, // number of items in the buffer
        start: usize = 0, // index of the oldest item
        end: usize = 0, // index of where the next write will go

        const Self = @This();

        pub fn init() Self {
            return .{};
        }

        /// "Discards" all items in the buffer.
        pub fn clear(self: *Self) void {
            self.count = 0;
            self.start = 0;
            self.end = 0;
        }

        /// "Pop" the oldest item from the buffer.
        pub fn read(self: *Self) ?T {
            if (self.count == 0) return null; // empty buffer returns null
            const val = self.buffer[self.start]; // pop oldest item
            self.start = (self.start + 1) % self.buffer.len; // advance next write to next position, ensures wrapping using mod
            self.count -= 1; // decrease number of items since we popped one
            return val;
        }

        /// Insert `item` into the buffer.
        pub fn write(self: *Self, item: T) BufferError!void {
            if (self.count == self.buffer.len) return BufferError.BufferOverflow; // return error if buffer is full
            self.overwrite(item); // otherwise we insert
        }

        /// Insert `item` into the buffer, replacing the oldest item if necessary.
        pub fn overwrite(self: *Self, item: T) void {
            self.buffer[self.end] = item; // store item where the next write will go
            self.end = (self.end + 1) % self.buffer.len; // advance where the next write will go, ensures wrapping using mod
            if (self.count != self.buffer.len) { // if the buffer isn't full yet
                self.count += 1; // increase number of items in the buffer
            } else { // otherwise buffer is full
                self.start = (self.start + 1) % self.buffer.len; // advance index of oldest item (this "drops" it)
            }
        }
    };
}
