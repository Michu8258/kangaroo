package helpers

func FirstOrDefault[T interface{}](slice []T, fallback T, predicate func(T) bool) T {
	for i := 0; i < len(slice); i++ {
		isMatching := predicate(slice[i])
		if isMatching {
			return slice[i]
		}
	}

	return fallback
}

func Where[T interface{}](slice []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range slice {
		isMatching := predicate(v)
		if isMatching {
			result = append(result, v)
		}
	}

	return result
}

func Any[T interface{}](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		isMatching := predicate(v)
		if isMatching {
			return true
		}
	}

	return false
}
