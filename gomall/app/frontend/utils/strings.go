package utils

func InArray[T int | int32 | int64 | float32 | float64 | string](needle T, haystack []T) bool {
	for _, k := range haystack {
		if needle == k {
			return true
		}
	}
	return false
}
