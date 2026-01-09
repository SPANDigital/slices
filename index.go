package slices

import "slices"

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
//
// Deprecated: Use slices.Index from the standard library instead.
// This function will be removed in v1.0.0.
//
// Migration:
//
//	import "slices"
//	slices.Index(mySlice, value)
func Index[S ~[]E, E comparable](s S, v E) int {
	return slices.Index(s, v)
}
