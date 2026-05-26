package util

import (
	"cmp"
	"iter"
)

// TransformSlice applies the converter to each element in the input slice and returns a new slice.
func TransformSlice[T any, R any](in []T, converter func(T) R) []R {
	_ = "STUB: not implemented"
	return nil
}

// CanonicalMapIter returns an iterator that yields map entries in sorted key order.
// This ensures deterministic iteration over maps, which is useful for generating
// consistent output (e.g., DDL statements) regardless of Go's random map iteration order.
func CanonicalMapIter[T any](m map[string]T) iter.Seq2[string, T] {
	_ = "STUB: not implemented"
	return nil
}

// SortedCopy returns a sorted copy of the input slice without modifying the original.
func SortedCopy[T cmp.Ordered](in []T) []T { _ = "STUB: not implemented"; return nil }
