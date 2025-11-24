use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    let lower = candidate.trim().to_lowercase();
    if lower.is_empty() {
        return true;
    }

    let mut map = HashSet::new();

    for c in lower.chars() {
        if c.is_alphabetic() {
            if map.contains(&c) {
                return false;
            }
        }
        map.insert(c);
    }

    true
}
