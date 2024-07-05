package slices

func AppendNotNil[T any](in []*T, values ...*T) (out []*T) {
	return append(in, FilterNil(values)...)
}
