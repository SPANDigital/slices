package slices

func Filter[S ~[]V, V any](s S, filter func(V V) bool) []V {
	out := make([]V, 0, len(s))
	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}
	return out
}
