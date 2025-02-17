package slices

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
