// Code generated by github.com/schigh/slice/base/base.go. DO NOT EDIT.

package slice

import "sort"

// Int32 is the alias of []int32
type Int32 []int32

// Value returns the aliased []int32
func (slc Int32) Value() []int32 {
	return []int32(slc)
}

// IndexOf returns the first index of needle, or -1 if not found
func (slc Int32) IndexOf(needle int32) int {
	for i := 0; i < len(slc); i++ {
		if needle == slc[i] {
			return i
		}
	}

	return NotInSlice
}

// Contains returns true if the slice contains needle
func (slc Int32) Contains(needle int32) bool {
	return slc.IndexOf(needle) != NotInSlice
}

// SortAsc will sort the slice in ascending order
func (slc Int32) SortAsc() Int32 {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i] < slc[j]
	})

	return slc
}

// SortDesc will sort the slice in descending order
func (slc Int32) SortDesc() Int32 {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[j] < slc[i]
	})

	return slc
}

// Reverse will reverse the order of the slice
func (slc Int32) Reverse() Int32 {
	l := len(slc)
	for i, j := 0, l-1; i < l/2; i++ {
		slc[i], slc[j] = slc[j], slc[i]
		j--
	}

	return slc
}

// Unique filters out duplicate int32 values
func (slc Int32) Unique() Int32 {
	u := make([]int32, 0, len(slc))
	m := make(map[int32]bool)

	for i := 0; i < len(slc); i++ {
		if _, ok := m[slc[i]]; !ok {
			m[slc[i]] = true
			u = append(u, slc[i])
		}
	}

	return Int32(u)
}

// Filter will return all int32 values that evaluate true in the user-supplied function
func (slc Int32) Filter(f func(int32) bool) Int32 {
	out := make([]int32, 0, len(slc))
	for i := 0; i < len(slc); i++ {
		if f(slc[i]) {
			out = append(out, slc[i])
		}
	}

	return Int32(out)
}

// Each will apply a function to each int32 in the slice.
// This function will iterate over the slice completely.  No
// items in the slice should be mutated by this operation.
func (slc Int32) Each(f func(int32)) {
	for i := 0; i < len(slc); i++ {
		f(slc[i])
	}
}

// TryEach will apply a function to each int32 in the slice.
// If the function returns an error, the iteration will stop and return
// the index of the element that caused the function to return the error.
// The second returned value will be first error returned from the
// supplied function, and nil otherwise.
// No items in the slice should be mutated by this operation.
func (slc Int32) TryEach(f func(int32) error) (int, error) {
	for i := 0; i < len(slc); i++ {
		if err := f(slc[i]); err != nil {
			return i, err
		}
	}

	return NotInSlice, nil
}

// IfEach will apply a function to each int32 in the slice.
// If the function returns false, the iteration will stop and return
// the index of the element that caused the function to return false.
// The second returned value will be true if all members of the slice
// cause the provided function to return true, and false otherwise.
// No items in the slice should be mutated by this operation.
func (slc Int32) IfEach(f func(int32) bool) (int, bool) {
	for i := 0; i < len(slc); i++ {
		if !f(slc[i]) {
			return i, false
		}
	}

	return NotInSlice, true
}

// Map will apply a function to each int32 in the slice and replace the previous value
func (slc Int32) Map(f func(int32) int32) {
	for i := 0; i < len(slc); i++ {
		slc[i] = f(slc[i])
	}
}

// Chunk will divide the slice of int32 into smaller slicify defined by chunk length
func (slc Int32) Chunk(size int) [][]int32 {
	l := len(slc)
	if l == 0 || size <= 0 {
		return make([][]int32, 0)
	}

	floor := l / size
	out := make([][]int32, 0, floor+1)
	var k int

	for i := 0; i < floor; i++ {
		k = i*size + size
		out = append(out, slc[i*size:k])
	}
	if l > k {
		out = append(out, slc[k:])
	}

	return out
}
