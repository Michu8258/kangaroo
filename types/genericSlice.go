package types

import (
	"slices"
)

type GenericSlice[T comparable] []T

func (slice GenericSlice[T]) Intersect(comparisionTarget GenericSlice[T]) GenericSlice[T] {
	commonItems := GenericSlice[T]{}

	for _, v := range slice {
		if slices.Contains(comparisionTarget, v) {
			commonItems = append(commonItems, v)
		}
	}

	for _, v := range comparisionTarget {
		if slices.Contains(slice, v) && !slices.Contains(commonItems, v) {
			commonItems = append(commonItems, v)
		}
	}

	return commonItems
}

func (slice GenericSlice[T]) EqualContent(comparisionTarget GenericSlice[T]) bool {
	if len(slice) != len(comparisionTarget) {
		return false
	}

	for _, value := range slice {
		if !slices.Contains(comparisionTarget, value) {
			return false
		}
	}

	for _, value := range comparisionTarget {
		if !slices.Contains(slice, value) {
			return false
		}
	}

	return true
}

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

func (slice GenericSlice[T]) All(predicate func(T) bool) bool {
	for _, v := range slice {
		isMatching := predicate(v)
		if !isMatching {
			return false
		}
	}

	return true
}
