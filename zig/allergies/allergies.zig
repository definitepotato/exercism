const std = @import("std");
const EnumSet = std.EnumSet;

pub const Allergen = enum {
    eggs,
    peanuts,
    shellfish,
    strawberries,
    tomatoes,
    chocolate,
    pollen,
    cats,
};

pub fn isAllergicTo(score: u8, allergen: Allergen) bool {
    const set = initAllergenSet(score);

    return set.contains(allergen);
}

pub fn initAllergenSet(score: usize) EnumSet(Allergen) {
    var allergens = EnumSet(Allergen).initEmpty();
    allergens.bits.mask = @truncate(score);

    return allergens;
}
