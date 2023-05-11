package slices

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"sort"
)

func Sort[T any](a []T) {
	if len(a) > 0 {
		switch any(a[0]).(type) {
		case uint:
			_SortOrdered[uint](any(a).([]uint))
		case uint8:
			_SortOrdered[uint8](any(a).([]uint8))
		case int8:
			_SortOrdered[int8](any(a).([]int8))
		case uint16:
			_SortOrdered[uint16](any(a).([]uint16))
		case int16:
			_SortOrdered[int16](any(a).([]int16))
		case uint32:
			_SortOrdered[uint32](any(a).([]uint32))
		case int32:
			_SortOrdered[int32](any(a).([]int32))
		case uint64:
			_SortOrdered[uint64](any(a).([]uint64))
		case int64:
			_SortOrdered[int64](any(a).([]int64))
		case uintptr:
			_SortOrdered[uintptr](any(a).([]uintptr))
		case float32:
			_SortOrdered[float32](any(a).([]float32))
		case int:
			sort.Ints(any(a).([]int))
		case string:
			sort.Strings(any(a).([]string))
		case float64:
			sort.Float64s(any(a).([]float64))
		case Sortable[T]:
			sort.Slice(a, func(i, j int) bool { return any(a[i]).(Sortable[T]).Less(a[j]) })
		default:
			panic(fmt.Sprintf("unsortable type %T", a[0]))
		}
	}
}

type Sortable[T any] interface{ Less(T) bool }

func _SortOrdered[T constraints.Ordered](a []T) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
}
