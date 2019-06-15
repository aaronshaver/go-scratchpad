package derangement

// Tests for derangement. Only works with true sets (collection of distinct objects)
func isDeranged(slice1, slice2 []string) bool {
	for i := range slice1 {
		if slice1[i] == slice2[i] {
			return false
		}
	}

	return true
}
