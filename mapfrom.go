package slices

func MapFrom[S ~[]I, I any, K comparable, V any](s S, extractKey func(i I) K, extractValue func(i I) V) map[K]V {
	out := make(map[K]V)
	for _, item := range s {
		out[extractKey(item)] = extractValue(item)
	}
	return out
}
