package slices

func GroupByLen[S ~[]V, V any](input S, length int) (output [][]V) {
	if length < 1 {
		panic("length must be greater than or equal to 1")
	}
	if len(input)%length != 0 {
		panic("input length must be a multiple of length")
	}
	var inner []V
	for idx, v := range input {
		if idx%length == 0 {
			inner = nil
		}
		inner = append(inner, v)
		if idx%length == length-1 {
			output = append(output, inner)
		}
	}
	return
}
