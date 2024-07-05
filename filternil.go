package slices

func FilterNil[T any](in []*T) (out []*T) {
	out = make([]*T, 0, len(in))
	for _, v := range in {
		if v != nil {
			out = append(out, v)
		}
	}
	return
}
