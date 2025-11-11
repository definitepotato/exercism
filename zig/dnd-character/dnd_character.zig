const std = @import("std");
var prng = std.rand.DefaultPrng.init(0);
const rand = prng.random();

pub fn modifier(score: i8) i8 {
    return @divFloor(score - 10, 2);
}

pub fn ability() i8 {
    var dice_results = [4]i8{
        rand.intRangeAtMost(i8, 1, 6),
        rand.intRangeAtMost(i8, 1, 6),
        rand.intRangeAtMost(i8, 1, 6),
        rand.intRangeAtMost(i8, 1, 6),
    };

    std.mem.sort(i8, &dice_results, {}, std.sort.desc(i8));
    return dice_results[0] + dice_results[1] + dice_results[2];
}

pub const Character = struct {
    strength: i8,
    dexterity: i8,
    constitution: i8,
    intelligence: i8,
    wisdom: i8,
    charisma: i8,
    hitpoints: i8,

    pub fn init() Character {
        const constitution = ability();
        return Character{
            .strength = ability(),
            .dexterity = ability(),
            .constitution = constitution,
            .intelligence = ability(),
            .wisdom = ability(),
            .charisma = ability(),
            .hitpoints = 10 + modifier(constitution),
        };
    }
};
