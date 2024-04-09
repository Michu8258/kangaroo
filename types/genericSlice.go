package types

type GenericSlice[T interface{}] []T

func (slice GenericSlice[T]) Where(predicate func(T) bool) GenericSlice[T] {
	result := GenericSlice[T]{}
	for _, v := range slice {
		isMatching := predicate(v)
		if isMatching {
			result = append(result, v)
		}
	}

	return result
}

func (slice GenericSlice[T]) FirstOrDefault(defaultValue T, predicate func(T) bool) T {
	for _, v := range slice {
		isMatching := predicate(v)
		if isMatching {
			return v
		}
	}

	return defaultValue
}

func (slice GenericSlice[T]) Any(predicate func(T) bool) bool {
	for _, v := range slice {
		isMatching := predicate(v)
		if isMatching {
			return true
		}
	}

	return false
}
