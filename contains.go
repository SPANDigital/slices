package slices

import "slices"

// Contains reports whether v is present in s.
//
// Deprecated: Use slices.Contains from the standard library instead.
// This function will be removed in v1.0.0.
//
// Migration:
//
//	import "slices"
//	slices.Contains(mySlice, value)
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return slices.Contains(s, v)
}
