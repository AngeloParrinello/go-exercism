package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for i := 0; i < len(birdsPerDay); i++ {
		if birdsPerDay[i] < 0 {
			panic("Bird count cannot be negative")
		}
		total += birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := startIndex + 7
	if startIndex < 0 || endIndex > len(birdsPerDay) {
		panic("Week index out of range")
	}
	total := 0
	for i := startIndex; i < endIndex; i++ {
		if birdsPerDay[i] < 0 {
			panic("Bird count cannot be negative")
		}
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
