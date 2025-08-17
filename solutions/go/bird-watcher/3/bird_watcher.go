package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
    for _, t := range birdsPerDay {
        total += t
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    start := (week - 1) * 7
    if start < 0 || start >= len(birdsPerDay) {
        return 0
    }
    end := start + 7
    if end > len(birdsPerDay) {
        end = len(birdsPerDay)
    }
    return TotalBirdCount(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	 for i := 0; i < len(birdsPerDay); i+=2 {
        birdsPerDay[i]++
     }
    return birdsPerDay
}
