package slices

import "math"

func NumPages[V any](s []V, pageSize int) int {
	return int(math.Ceil(float64(len(s)) / float64(pageSize)))
}
