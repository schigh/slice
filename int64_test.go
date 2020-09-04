// Code generated by github.com/schigh/slice/base/base.go. DO NOT EDIT.
package slice

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestInt64_IndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int64
		needle   int64
		expected int
	}{
		{
			name:     "only item",
			slc:      []int64{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slc:      []int64{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slc:      []int64{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int64(tt.slc).IndexOf(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestInt64_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int64
		needle   int64
		expected bool
	}{
		{
			name:     "present",
			slc:      []int64{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slc:      []int64{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int64(tt.slc).Contains(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt64_SortAsc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int64
		expected []int64
	}{
		{
			name:     "empty",
			slc:      []int64{},
			expected: []int64{},
		},
		{
			name:     "already sorted",
			slc:      []int64{0, 1, 2, 3, 4, 5},
			expected: []int64{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slc:      []int64{5, 4, 3, 2, 1, 0},
			expected: []int64{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slc:      []int64{3, 1, 4, 5, 0, 2},
			expected: []int64{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int64(tt.slc).SortAsc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt64_SortDesc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int64
		expected []int64
	}{
		{
			name:     "empty",
			slc:      []int64{},
			expected: []int64{},
		},
		{
			name:     "already sorted",
			slc:      []int64{5, 4, 3, 2, 1, 0},
			expected: []int64{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slc:      []int64{0, 1, 2, 3, 4, 5},
			expected: []int64{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slc:      []int64{3, 1, 4, 5, 0, 2},
			expected: []int64{5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int64(tt.slc).SortDesc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt64_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int64
		expected []int64
	}{
		{
			name:     "unaffected",
			slc:      []int64{0, 1, 2, 3, 4, 5},
			expected: []int64{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slc:      []int64{5, 0, 1, 2, 3, 4, 5},
			expected: []int64{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slc:      []int64{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []int64{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int64(tt.slc).Unique().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt64_Reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int64
		expected []int64
	}{
		{
			name:     "even length",
			slc:      []int64{0, 1, 2, 3, 4, 5},
			expected: []int64{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slc:      []int64{0, 1, 2, 3, 4, 5, 6},
			expected: []int64{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int64(tt.slc).Reverse().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt64_Filter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		slc        []int64
		expected   []int64
		filterFunc func(int64) bool
	}{
		{
			name:       "gt 10",
			slc:        []int64{1, 2, 5, 11, 13, 15},
			expected:   []int64{11, 13, 15},
			filterFunc: func(n int64) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slc:        []int64{1, 2, 6, 11, 12, 15, 17},
			expected:   []int64{6, 12, 15},
			filterFunc: func(n int64) bool { return n%3 == 0 },
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Int64(tt.slc).Filter(tt.filterFunc).Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestInt64_Each(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(int64(0))

	tests := []struct {
		name     string
		slc      []int64
		expected int64
		eachFunc func(int64)
	}{
		{
			name:     "add n",
			slc:      []int64{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n int64) {
				v := rabbit.Load().(int64)
				v = v + n
				rabbit.Store(v)
			},
		},
		{
			name:     "subtract n",
			slc:      []int64{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n int64) {
				v := rabbit.Load().(int64)
				v = v - n
				rabbit.Store(v)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			Int64(tt.slc).Each(tt.eachFunc)
			v := rabbit.Load().(int64)
			if tt.expected != v {
				t.Errorf("expected %v, got %v", tt.expected, rabbit)
			}
		})
	}
}

func TestInt64_TryEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(int64(0))

	myErr := errors.New("i am an error")
	tests := []struct {
		name      string
		slc       []int64
		expected  int
		expected2 error
		before    func()
		eachFunc  func(int64) error
	}{
		{
			name:      "add n",
			slc:       []int64{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n int64) error {
				v := rabbit.Load().(int64)
				v = v + n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name:      "subtract n",
			slc:       []int64{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n int64) error {
				v := rabbit.Load().(int64)
				v = v - n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name:      "errors",
			slc:       []int64{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: myErr,
			before:    func() { rabbit.Store(int64(0)) },
			eachFunc: func(n int64) error {
				v := rabbit.Load().(int64)
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
			e, i := Int64(tt.slc).TryEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func TestInt64_IfEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(int64(0))

	tests := []struct {
		name      string
		slc       []int64
		expected  int
		expected2 bool
		before    func()
		err       error
		eachFunc  func(int64) bool
	}{
		{
			name:      "all return true",
			slc:       []int64{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n int64) bool {
				v := rabbit.Load().(int64)
				v = v + n
				rabbit.Store(v)
				return true
			},
		},
		{
			name:      "subtract n",
			slc:       []int64{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n int64) bool {
				v := rabbit.Load().(int64)
				v = v - n
				rabbit.Store(v)
				return true
			},
		},
		{
			name:      "breaking",
			slc:       []int64{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: false,
			before:    func() { rabbit.Store(int64(0)) },
			eachFunc: func(n int64) bool {
				v := rabbit.Load().(int64)
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
			e, i := Int64(tt.slc).IfEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func TestInt64_Map(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []int64
		expected []int64
		mapFunc  func(int64) int64
	}{
		{
			name:     "add 3",
			slc:      []int64{1, 2, 5, 11, 13, 15},
			expected: []int64{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n int64) int64 { return n + 3 },
		},
		{
			name:     "set mod 2",
			slc:      []int64{1, 2, 6, 8, 12, 15, 17},
			expected: []int64{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n int64) int64 { return int64(n % 2) },
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Int64(tt.slc).Map(tt.mapFunc)
			if !reflect.DeepEqual(tt.expected, tt.slc) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func TestInt64_Chunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		size     int
		slc      []int64
		expected [][]int64
	}{
		{
			name:     "chunks of 2 no remainder",
			size:     2,
			slc:      []int64{1, 2, 5, 11, 13, 15},
			expected: [][]int64{[]int64{1, 2}, []int64{5, 11}, []int64{13, 15}},
		},
		{
			name:     "chunks of 2 with remainder",
			size:     2,
			slc:      []int64{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int64{[]int64{1, 2}, []int64{5, 11}, []int64{13, 15}, []int64{17}},
		},
		{
			name:     "chunks of 100",
			size:     100,
			slc:      []int64{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int64{[]int64{1, 2, 5, 11, 13, 15, 17}},
		},
		{
			name:     "chunks of 4",
			size:     4,
			slc:      []int64{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int64{[]int64{1, 2, 5, 11}, []int64{13, 15, 17}},
		},
		{
			name:     "chunks of 5",
			size:     5,
			slc:      []int64{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int64{[]int64{1, 2, 5, 11, 13}, []int64{15, 17}},
		},
		{
			name:     "empty slice",
			size:     5,
			slc:      []int64{},
			expected: [][]int64{},
		},
		{
			name:     "invalid chunk size",
			size:     -1,
			slc:      []int64{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int64{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := Int64(tt.slc).Chunk(tt.size)
			if !reflect.DeepEqual(tt.expected, out) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func BenchmarkInt64_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).IndexOf(1)
			}
		})
	}
}

func BenchmarkInt64_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).Contains(1)
			}
		})
	}
}

func BenchmarkInt64_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).SortAsc()
			}
		})
	}
}

func BenchmarkInt64_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).SortDesc()
			}
		})
	}
}

func BenchmarkInt64_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).Reverse()
			}
		})
	}
}

func BenchmarkInt64_Filter(b *testing.B) {
	benchFunc := func(n int64) bool {
		return true
	}
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).Filter(benchFunc)
			}
		})
	}
}

func BenchmarkInt64_Each(b *testing.B) {
	benchFunc := func(n int64) { n = 0 }
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).Each(benchFunc)
			}
		})
	}
}

func BenchmarkInt64_Map(b *testing.B) {
	benchFunc := func(n int64) int64 {
		return n
	}
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).Map(benchFunc)
			}
		})
	}
}

func BenchmarkInt64_Unique(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []int64
	}{
		{
			name: "10 elements",
			slc:  genInt64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genInt64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genInt64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genInt64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genInt64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64(bm.slc).Unique()
			}
		})
	}
}
