package calclib

// SumUntillMax sums up untill max and returns the number of loops
func SumUntillMax(max int) int {
	sum := 0
	for i := 0; ; i++ {
		sum += i
		if sum >= max {
			return i
		}
	}
}
