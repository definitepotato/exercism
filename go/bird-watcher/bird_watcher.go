package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, day := range birdsPerDay {
		total += day
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	weekCount := 1
	dayCount := 1

	for _, day := range birdsPerDay {
		if weekCount == week {
			total += day
		}

		if dayCount == 7 {
			dayCount = 0
			weekCount += 1
		}
		dayCount += 1
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx, day := range birdsPerDay {
		if idx%2 == 0 {
			birdsPerDay[idx] = day + 1
		}
	}

	return birdsPerDay
}
