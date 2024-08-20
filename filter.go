package slices

func Filter[S ~[]V, V any](s S, predicate func(V V) bool) []V {
	out := make([]V, 0, len(s))
	for _, v := range s {
		if predicate(v) {
			out = append(out, v)
		}
	}
	return out
}
