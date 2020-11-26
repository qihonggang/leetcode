package array

func findPoisonedDuration(timeSeries []int, duration int) int {
	var n = len(timeSeries)
	if n == 0 {
		return 0
	}
	var total int
	for i := 0; i < n - 1; i++ {
		total += Min(timeSeries[i + 1] - timeSeries[i], duration)
	}
	return total + duration
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
