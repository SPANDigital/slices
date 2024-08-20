package slices

func Map[S ~[]V, V any, E any](s S, extract func(v V) E) []E {
	out := make([]E, len(s))
	for idx, v := range s {
		out[idx] = extract(v)
	}
	return out
}
