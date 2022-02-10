package calclib

// SumUntilMax sums up until max and returns the number of loops
func SumUntilMax(max int) int {
	sum := 0
	for i := 0; ; i++ {
		sum += i
		if sum >= max {
			return i
		}
	}
}
