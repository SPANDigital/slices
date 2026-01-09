package slices

// GroupByLen splits input into chunks of the specified length.
//
// Deprecated: Use slices.Chunk from the standard library instead.
// This function will be removed in v1.0.0.
// Note that slices.Chunk returns an iterator, not a materialized slice.
//
// Migration:
//
//	import "slices"
//	// To collect into a slice of slices:
//	var result [][]V
//	for chunk := range slices.Chunk(mySlice, length) {
//	    result = append(result, chunk)
//	}
func GroupByLen[S ~[]V, V any](input S, length int) (output [][]V) {
	if length < 1 {
		panic("length must be greater than or equal to 1")
	}
	var inner []V
	for idx, v := range input {
		if idx%length == 0 {
			inner = nil
		}
		inner = append(inner, v)
		if idx%length == length-1 {
			output = append(output, inner)
			inner = nil
		}
	}
	if len(inner) > 0 {
		output = append(output, inner)
	}
	return
}
