// Package strain provides functions to filter slices based on a predicate.
package strain

// Keep returns elements for which pred == true.
func Keep[T any](data []T, pred func(T) bool) []T {
	return filter(data, pred)
}

// Discard returns elements for which pred == false.
func Discard[T any](data []T, pred func(T) bool) []T {
	return filter(data, func(x T) bool { return !pred(x) })
}

// filter is an internal helper that keeps elements passing the predicate.
func filter[T any](data []T, pred func(T) bool) []T {
	out := make([]T, 0, len(data))
	for _, v := range data {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}
