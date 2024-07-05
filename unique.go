package slices

func Unique[T comparable](s []T) []T {
	set := make(map[T]struct{})
	for _, v := range s {
		set[v] = struct{}{}
	}
	var result []T
	for k := range set {
		result = append(result, k)
	}
	return result
}
