package slices

func GroupBy[S ~[]V, V any, K comparable](s S, extract func(v V) K) map[K][]V {
	out := make(map[K][]V)
	for _, item := range s {
		e := extract(item)
		out[e] = append(out[e], item)
	}
	return out
}
