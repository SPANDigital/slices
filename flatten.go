package slices

func Flatten[T any](s [][]T) (flattened []T) {
	for _, oneValue := range s {
		flattened = append(flattened, oneValue...)
	}
	return
}

func Flatten3[T any](s [][][]T) (flattened []T) {
	for _, oneValue := range s {
		for _, twoValue := range oneValue {
			flattened = append(flattened, twoValue...)
		}
	}
	return
}

func Flatten4[T any](s [][][][]T) (flattened []T) {
	for _, oneValue := range s {
		for _, twoValue := range oneValue {
			for _, threeValue := range twoValue {
				flattened = append(flattened, threeValue...)
			}
		}
	}
	return
}
