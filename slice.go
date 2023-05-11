package slices

func New[T any](t ...T) []T { return t }

// Filter returns a new array containing only elements that meet fn(T) == true.
func Filter[T any](a []T, fn func(T) bool) []T {
	a2 := make([]T, 0, len(a))
	for _, v := range a {
		if fn(v) {
			a2 = append(a2, v)
		}
	}
	return a2[0:len(a2):len(a2)]
}

func FlatMap[T any](a [][]T) []T {
	l := 0
	for _, e := range a {
		l += len(e)
	}
	o := make([]T, 0, l)
	for _, a := range a {
		o = append(o, a...)
	}
	return o
}

// Reverse reverses input in place, returning input.
func Reverse[T any](input []T) []T {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

// Map maps a slice to another type.
// out will be a new slice.
// If fn == nil, nil is returned, otherwise len(out) == len(a).
func Map[I, O any](input []I, fn func(index int, elem I) O) (out []O) {
	if fn == nil {
		return nil
	}

	out = make([]O, len(input), cap(input))
	for i, e := range input {
		out[i] = fn(i, e)
	}
	return out
}

func ForEach[T any](input []T, fn func(index int, elem T)) {
	if fn != nil {
		for i, e := range input {
			fn(i, e)
		}
	}
}

// Copy copies a slice
func Copy[T any](input []T) (out []T) {
	out = make([]T, len(input))
	copy(out, input)
	return out
}
