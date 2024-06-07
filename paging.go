package slices

import "math"

func NumPages[V any](s []V, pageSize int) int {
	return int(math.Ceil(float64(len(s)) / float64(pageSize)))
}

func Page[V any](s []V, pageSize int, pageIndex int) []V {
	return s[pageIndex*pageSize : min(len(s), (pageIndex+1)*pageSize)]
}
