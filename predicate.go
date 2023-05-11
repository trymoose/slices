package slices

func IsEmpty[T any](a []T) bool    { return len(a) == 0 }
func IsNotEmpty[T any](a []T) bool { return !IsEmpty(a) }
