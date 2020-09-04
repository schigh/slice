// Code generated by github.com/schigh/slice/base/base.go. DO NOT EDIT.

package slice

import "sort"

// Int is the alias of []int
type Int []int

// Value returns the aliased []int
func (slc Int) Value() []int {
	return []int(slc)
}

// IndexOf returns the first index of needle, or -1 if not found
func (slc Int) IndexOf(needle int) int {
	for i := 0; i < len(slc); i++ {
		if needle == slc[i] {
			return i
		}
	}

	return NotInSlice
}

// Contains returns true if the slice contains needle
func (slc Int) Contains(needle int) bool {
	return slc.IndexOf(needle) != NotInSlice
}

// SortAsc will sort the slice in ascending order
func (slc Int) SortAsc() Int {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i] < slc[j]
	})

	return slc
}

// SortDesc will sort the slice in descending order
func (slc Int) SortDesc() Int {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[j] < slc[i]
	})

	return slc
}

// Reverse will reverse the order of the slice
func (slc Int) Reverse() Int {
	l := len(slc)
	for i, j := 0, l-1; i < l/2; i++ {
		slc[i], slc[j] = slc[j], slc[i]
		j--
	}

	return slc
}

// Unique filters out duplicate int values
func (slc Int) Unique() Int {
	u := make([]int, 0, len(slc))
	m := make(map[int]bool)

	for i := 0; i < len(slc); i++ {
		if _, ok := m[slc[i]]; !ok {
			m[slc[i]] = true
			u = append(u, slc[i])
		}
	}

	return Int(u)
}

// Filter will return all int values that evaluate true in the user-supplied function
func (slc Int) Filter(f func(int) bool) Int {
	out := make([]int, 0, len(slc))
	for i := 0; i < len(slc); i++ {
		if f(slc[i]) {
			out = append(out, slc[i])
		}
	}

	return Int(out)
}

// Each will apply a function to each int in the slice.
// This function will iterate over the slice completely.  No
// items in the slice should be mutated by this operation.
func (slc Int) Each(f func(int)) {
	for i := 0; i < len(slc); i++ {
		f(slc[i])
	}
}

// TryEach will apply a function to each int in the slice.
// If the function returns an error, the iteration will stop and return
// the index of the element that caused the function to return the error.
// The second returned value will be first error returned from the
// supplied function, and nil otherwise.
// No items in the slice should be mutated by this operation.
func (slc Int) TryEach(f func(int) error) (int, error) {
	for i := 0; i < len(slc); i++ {
		if err := f(slc[i]); err != nil {
			return i, err
		}
	}

	return NotInSlice, nil
}

// IfEach will apply a function to each int in the slice.
// If the function returns false, the iteration will stop and return
// the index of the element that caused the function to return false.
// The second returned value will be true if all members of the slice
// cause the provided function to return true, and false otherwise.
// No items in the slice should be mutated by this operation.
func (slc Int) IfEach(f func(int) bool) (int, bool) {
	for i := 0; i < len(slc); i++ {
		if !f(slc[i]) {
			return i, false
		}
	}

	return NotInSlice, true
}

// Map will apply a function to each int in the slice and replace the previous value
func (slc Int) Map(f func(int) int) {
	for i := 0; i < len(slc); i++ {
		slc[i] = f(slc[i])
	}
}

// Chunk will divide the slice of int into smaller slices defined by chunk length
func (slc Int) Chunk(size int) [][]int {
	l := len(slc)
	if l == 0 || size <= 0 {
		return make([][]int, 0)
	}

	floor := l / size
	out := make([][]int, 0, floor+1)
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