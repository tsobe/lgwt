package arrays

func Find[T any](collection []T, f func(x T) bool) (value T, found bool) {
	for _, val := range collection {
		if f(val) {
			return val, true
		}
	}
	return
}
