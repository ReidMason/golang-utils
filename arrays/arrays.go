package arrays

import "errors"

// The Filter function returns a filtered copy of the provided slice.
// Only elements that return true under the provided predicate will be contained in the returned slice
//
//	filtered = array.Filter(slice, fn func(x int32) bool { return x == 123 })
func Filter[T any](arr []T, fn func(x T) bool) []T {
	res := make([]T, 0)
	for _, x := range arr {
		if fn(x) {
			res = append(res, x)
		}
	}

	return res
}

// The Map function returns a copy of the provided slice with the supplied function run against all elements.
//
//	mapped = array.Map(slice, fn func(x int32) String { return fmt.Sprint(x) })
func Map[T, Y any](arr []T, fn func(x T) Y) []Y {
	res := make([]Y, 0, len(arr))
	for _, x := range arr {
		res = append(res, fn(x))
	}

	return res
}

// The FirstOrDefault function returns the first element or the provided default value
//
//	first := FirstOrDefault(slice, 0)
func FirstOrDefault[T any](arr []T, fn func(x T) bool, defaultValue T) T {
	for _, x := range arr {
		if fn(x) {
			return x
		}
	}

	return defaultValue
}

func Some[T any](arr []T, fn func(x T) bool) bool {
	for _, x := range arr {
		if fn(x) {
			return true
		}
	}

	return false
}

func Every[T any](arr []T, fn func(x T) bool) bool {
	for _, x := range arr {
		if !fn(x) {
			return false
		}
	}

	return true
}

func Find[T any](arr []T, fn func(x T) bool) (T, int, error) {
	for i, x := range arr {
		if fn(x) {
			return x, i, nil
		}
	}

	var val T
	return val, -1, errors.New("Element not found")
}

// Sort - Sort the array maybe using a lambda?

func qs[T any](arr []T, lo int, hi int, fn func(a, b T) bool) []T {
	if lo >= hi {
		return arr
	}

	pivotIndex, arr := partition(arr, lo, hi, fn)
	arr = qs(arr, lo, pivotIndex-1, fn)
	return qs(arr, pivotIndex+1, hi, fn)
}

func partition[T any](arr []T, lo int, hi int, fn func(a, b T) bool) (int, []T) {
	pivot := arr[hi]

	index := lo - 1
	for i := lo; i < hi; i++ {
		if fn(arr[i], pivot) {
			index++

			// Swap elements
			temp := arr[i]
			arr[i] = arr[index]
			arr[index] = temp
		}
	}

	index++
	// Swap the pivot to the last moved value location
	arr[hi] = arr[index]
	arr[index] = pivot

	return index, arr
}

// Sort a slice using the quicksort sorting method
//
//	QuickSort(slice, fn func(a, b int) bool { return a >= b})
func QuickSort[T any](arr []T, fn func(a, b T) bool) []T {
	return qs(arr, 0, len(arr)-1, fn)
}
