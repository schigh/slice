// Code generated by github.com/schigh/slice/base/base.go. DO NOT EDIT.
package slice

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestInt32_IndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int32
		needle   int32
		expected int
	}{
		{
			name:     "only item",
			slc:      []int32{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slc:      []int32{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slc:      []int32{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int32(tt.slc).IndexOf(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestInt32_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int32
		needle   int32
		expected bool
	}{
		{
			name:     "present",
			slc:      []int32{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slc:      []int32{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int32(tt.slc).Contains(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt32_SortAsc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int32
		expected []int32
	}{
		{
			name:     "empty",
			slc:      []int32{},
			expected: []int32{},
		},
		{
			name:     "already sorted",
			slc:      []int32{0, 1, 2, 3, 4, 5},
			expected: []int32{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slc:      []int32{5, 4, 3, 2, 1, 0},
			expected: []int32{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slc:      []int32{3, 1, 4, 5, 0, 2},
			expected: []int32{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int32(tt.slc).SortAsc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt32_SortDesc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int32
		expected []int32
	}{
		{
			name:     "empty",
			slc:      []int32{},
			expected: []int32{},
		},
		{
			name:     "already sorted",
			slc:      []int32{5, 4, 3, 2, 1, 0},
			expected: []int32{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slc:      []int32{0, 1, 2, 3, 4, 5},
			expected: []int32{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slc:      []int32{3, 1, 4, 5, 0, 2},
			expected: []int32{5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int32(tt.slc).SortDesc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt32_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int32
		expected []int32
	}{
		{
			name:     "unaffected",
			slc:      []int32{0, 1, 2, 3, 4, 5},
			expected: []int32{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slc:      []int32{5, 0, 1, 2, 3, 4, 5},
			expected: []int32{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slc:      []int32{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []int32{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int32(tt.slc).Unique().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt32_Reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int32
		expected []int32
	}{
		{
			name:     "even length",
			slc:      []int32{0, 1, 2, 3, 4, 5},
			expected: []int32{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slc:      []int32{0, 1, 2, 3, 4, 5, 6},
			expected: []int32{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int32(tt.slc).Reverse().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt32_Filter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		slc        []int32
		expected   []int32
		filterFunc func(int32) bool
	}{
		{
			name:       "gt 10",
			slc:        []int32{1, 2, 5, 11, 13, 15},
			expected:   []int32{11, 13, 15},
			filterFunc: func(n int32) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slc:        []int32{1, 2, 6, 11, 12, 15, 17},
			expected:   []int32{6, 12, 15},
			filterFunc: func(n int32) bool { return n%3 == 0 },
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int32(tt.slc).Filter(tt.filterFunc).Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt32_Each(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(int32(0))

	tests := []struct {
		name     string
		slc      []int32
		expected int32
		eachFunc func(int32)
	}{
		{
			name:     "add n",
			slc:      []int32{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n int32) {
				v := rabbit.Load().(int32)
				v = v + n
				rabbit.Store(v)
			},
		},
		{
			name:     "subtract n",
			slc:      []int32{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n int32) {
				v := rabbit.Load().(int32)
				v = v - n
				rabbit.Store(v)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			Int32(tt.slc).Each(tt.eachFunc)
			v := rabbit.Load().(int32)
			if tt.expected != v {
				t.Errorf("expected %v, got %v", tt.expected, rabbit)
			}
		})
	}
}

func TestInt32_TryEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(int32(0))

	myErr := errors.New("i am an error")
	tests := []struct {
		name      string
		slc       []int32
		expected  int
		expected2 error
		before    func()
		eachFunc  func(int32) error
	}{
		{
			name:      "add n",
			slc:       []int32{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n int32) error {
				v := rabbit.Load().(int32)
				v = v + n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name:      "subtract n",
			slc:       []int32{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n int32) error {
				v := rabbit.Load().(int32)
				v = v - n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name:      "errors",
			slc:       []int32{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: myErr,
			before:    func() { rabbit.Store(int32(0)) },
			eachFunc: func(n int32) error {
				v := rabbit.Load().(int32)
				if n > 5 {
					return myErr
				}
				v = v + n
				rabbit.Store(v)
				return nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before()
			}
			e, i := Int32(tt.slc).TryEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func TestInt32_IfEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(int32(0))

	tests := []struct {
		name      string
		slc       []int32
		expected  int
		expected2 bool
		before    func()
		err       error
		eachFunc  func(int32) bool
	}{
		{
			name:      "all return true",
			slc:       []int32{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n int32) bool {
				v := rabbit.Load().(int32)
				v = v + n
				rabbit.Store(v)
				return true
			},
		},
		{
			name:      "subtract n",
			slc:       []int32{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n int32) bool {
				v := rabbit.Load().(int32)
				v = v - n
				rabbit.Store(v)
				return true
			},
		},
		{
			name:      "breaking",
			slc:       []int32{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: false,
			before:    func() { rabbit.Store(int32(0)) },
			eachFunc: func(n int32) bool {
				v := rabbit.Load().(int32)
				if n > 5 {
					return false
				}
				v = v + n
				rabbit.Store(v)
				return true
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before()
			}
			e, i := Int32(tt.slc).IfEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func TestInt32_Map(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int32
		expected []int32
		mapFunc  func(int32) int32
	}{
		{
			name:     "add 3",
			slc:      []int32{1, 2, 5, 11, 13, 15},
			expected: []int32{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n int32) int32 { return n + 3 },
		},
		{
			name:     "set mod 2",
			slc:      []int32{1, 2, 6, 8, 12, 15, 17},
			expected: []int32{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n int32) int32 { return int32(n % 2) },
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Int32(tt.slc).Map(tt.mapFunc)
			if !reflect.DeepEqual(tt.expected, tt.slc) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func TestInt32_Chunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		size     int
		slc      []int32
		expected [][]int32
	}{
		{
			name:     "chunks of 2 no remainder",
			size:     2,
			slc:      []int32{1, 2, 5, 11, 13, 15},
			expected: [][]int32{[]int32{1, 2}, []int32{5, 11}, []int32{13, 15}},
		},
		{
			name:     "chunks of 2 with remainder",
			size:     2,
			slc:      []int32{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int32{[]int32{1, 2}, []int32{5, 11}, []int32{13, 15}, []int32{17}},
		},
		{
			name:     "chunks of 100",
			size:     100,
			slc:      []int32{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int32{[]int32{1, 2, 5, 11, 13, 15, 17}},
		},
		{
			name:     "chunks of 4",
			size:     4,
			slc:      []int32{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int32{[]int32{1, 2, 5, 11}, []int32{13, 15, 17}},
		},
		{
			name:     "chunks of 5",
			size:     5,
			slc:      []int32{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int32{[]int32{1, 2, 5, 11, 13}, []int32{15, 17}},
		},
		{
			name:     "empty slice",
			size:     5,
			slc:      []int32{},
			expected: [][]int32{},
		},
		{
			name:     "invalid chunk size",
			size:     -1,
			slc:      []int32{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int32{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := Int32(tt.slc).Chunk(tt.size)
			if !reflect.DeepEqual(tt.expected, out) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func BenchmarkInt32_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).IndexOf(1)
			}
		})
	}
}

func BenchmarkInt32_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).Contains(1)
			}
		})
	}
}

func BenchmarkInt32_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).SortAsc()
			}
		})
	}
}

func BenchmarkInt32_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).SortDesc()
			}
		})
	}
}

func BenchmarkInt32_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).Reverse()
			}
		})
	}
}

func BenchmarkInt32_Filter(b *testing.B) {
	benchFunc := func(n int32) bool {
		return true
	}
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).Filter(benchFunc)
			}
		})
	}
}

func BenchmarkInt32_Each(b *testing.B) {
	benchFunc := func(n int32) { n = 0 }
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).Each(benchFunc)
			}
		})
	}
}

func BenchmarkInt32_Map(b *testing.B) {
	benchFunc := func(n int32) int32 {
		return n
	}
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).Map(benchFunc)
			}
		})
	}
}

func BenchmarkInt32_Unique(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int32
	}{
		{
			name: "10 elements",
			slc:  genInt32(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt32(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt32(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt32(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt32(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32(bm.slc).Unique()
			}
		})
	}
}
