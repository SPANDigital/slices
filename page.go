package slices

func Page[V any](s []V, pageSize int, pageIndex int) []V {
	return s[pageIndex*pageSize : min(len(s), (pageIndex+1)*pageSize)]
}
