const std = @import("std");

pub fn calculate(seconds: usize, orbital_period: f64) f64 {
    const input: f64 = @floatFromInt(seconds);

    const seconds_per_day = 86_400.00 * orbital_period / orbital_period;
    const days_per_year = 365.25 * orbital_period;

    const factor = std.math.pow(f64, 10, @as(f64, @floatFromInt(2)));
    const result = input / days_per_year / seconds_per_day;

    return @round(result * factor) / factor;
}

pub const Planet = enum {
    mercury,
    venus,
    earth,
    mars,
    jupiter,
    saturn,
    uranus,
    neptune,

    pub fn age(self: Planet, seconds: usize) f64 {
        switch (self) {
            .mercury => return calculate(seconds, 0.2408467),
            .venus => return calculate(seconds, 0.61519726),
            .earth => return calculate(seconds, 1.0),
            .mars => return calculate(seconds, 1.8808158),
            .jupiter => return calculate(seconds, 11.862615),
            .saturn => return calculate(seconds, 29.447498),
            .uranus => return calculate(seconds, 84.016846),
            .neptune => return calculate(seconds, 164.79132),
        }
    }
};
