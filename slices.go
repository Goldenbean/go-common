package common

// SliceContains : slice contains substr
func SliceContains(slice []string, substr string) bool {
	for _, n := range slice {
		if substr == n {
			return true
		}
	}
	return false
}
