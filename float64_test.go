// Code generated by github.com/schigh/slice/base/base.go. DO NOT EDIT.
package slice

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestFloat64_IndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []float64
		needle   float64
		expected int
	}{
		{
			name:     "only item",
			slc:      []float64{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slc:      []float64{0, 1.1, 1.1},
			needle:   1.1,
			expected: 1,
		},
		{
			name:     "missing",
			slc:      []float64{1.1, 2.1, 3.1, 4},
			needle:   5.1,
			expected: NotInSlice,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Float64(tt.slc).IndexOf(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestFloat64_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []float64
		needle   float64
		expected bool
	}{
		{
			name:     "present",
			slc:      []float64{1.1, 2.0, 3.2, 4.5},
			needle:   4.5,
			expected: true,
		},
		{
			name:     "not present",
			slc:      []float64{1.1, 2.0, 3.2, 4.5},
			needle:   5.6,
			expected: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Float64(tt.slc).Contains(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFloat64_SortAsc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []float64
		expected []float64
	}{
		{
			name:     "empty",
			slc:      []float64{},
			expected: []float64{},
		},
		{
			name:     "already sorted",
			slc:      []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
			expected: []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
		},
		{
			name:     "reversed",
			slc:      []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
			expected: []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
		},
		{
			name:     "random",
			slc:      []float64{3.0, 1.1, 4.32, 5.4, 0.123, 2.99},
			expected: []float64{0.123, 1.1, 2.99, 3.0, 4.32, 5.4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Float64(tt.slc).SortAsc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFloat64_SortDesc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []float64
		expected []float64
	}{
		{
			name:     "empty",
			slc:      []float64{},
			expected: []float64{},
		},
		{
			name:     "already sorted",
			slc:      []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
			expected: []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
		},
		{
			name:     "reversed",
			slc:      []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
			expected: []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
		},
		{
			name:     "random",
			slc:      []float64{3.0, 1.1, 4.32, 5.4, 0.123, 2.99},
			expected: []float64{5.4, 4.32, 3.0, 2.99, 1.1, 0.123},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Float64(tt.slc).SortDesc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFloat64_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []float64
		expected []float64
	}{
		{
			name:     "unaffected",
			slc:      []float64{0.123, 0.1, 0.3, 0.99, 0.432, 0.101},
			expected: []float64{0.123, 0.1, 0.3, 0.99, 0.432, 0.101},
		},
		{
			name:     "one extra 0.101",
			slc:      []float64{0.123, 0.1, 0.101, 0.3, 0.99, 0.432, 0.101},
			expected: []float64{0.123, 0.1, 0.101, 0.3, 0.99, 0.432},
		},
		{
			name:     "extras everywhere",
			slc:      []float64{0.1, 0.1, 1.2, 0.1, 1.2, 2.3, 2.3, 2.3, 3.4, 0.1, 3.4, 4.5, 2.3, 3.4, 4.5, 4.5, 2.3, 1.2, 0.1},
			expected: []float64{0.1, 1.2, 2.3, 3.4, 4.5},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Float64(tt.slc).Unique().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFloat64_Reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []float64
		expected []float64
	}{
		{
			name:     "even length",
			slc:      []float64{0.123, 1.234, 2.345, 3.456, 4.567, 5.678},
			expected: []float64{5.678, 4.567, 3.456, 2.345, 1.234, 0.123},
		},
		{
			name:     "odd length",
			slc:      []float64{0.123, 1.234, 2.345, 3.456, 4.567, 5.678, 6.789},
			expected: []float64{6.789, 5.678, 4.567, 3.456, 2.345, 1.234, 0.123},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Float64(tt.slc).Reverse().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFloat64_Filter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		slc        []float64
		expected   []float64
		filterFunc func(float64) bool
	}{
		{
			name:       "gt 10.5",
			slc:        []float64{1.2, 2.3, 5.6, 11.12, 13.14, 15.16},
			expected:   []float64{11.12, 13.14, 15.16},
			filterFunc: func(n float64) bool { return n > 10.5 },
		},
		{
			name:       "gt 0",
			slc:        []float64{0, -0.000001, 0.000001, -0.0000001, 0.0000001},
			expected:   []float64{0.000001, 0.0000001},
			filterFunc: func(n float64) bool { return n > 0 },
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Float64(tt.slc).Filter(tt.filterFunc).Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFloat64_Each(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(float64(0))

	tests := []struct {
		name     string
		slc      []float64
		expected float64
		eachFunc func(float64)
	}{
		{
			name:     "add n",
			slc:      []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected: 49.5,
			eachFunc: func(n float64) {
				v := rabbit.Load().(float64)
				v = v + n
				rabbit.Store(v)
			},
		},
		{
			name:     "subtract n",
			slc:      []float64{1.5, 2.25, 6.75, 8.5, 12.25},
			expected: 18.25,
			eachFunc: func(n float64) {
				v := rabbit.Load().(float64)
				v = v - n
				rabbit.Store(v)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			Float64(tt.slc).Each(tt.eachFunc)
			v := rabbit.Load().(float64)
			if tt.expected != v {
				t.Errorf("expected %v, got %v", tt.expected, rabbit)
			}
		})
	}
}

func TestFloat64_TryEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(float64(0))

	myErr := errors.New("i am an error")
	tests := []struct {
		name      string
		slc       []float64
		expected  int
		expected2 error
		before    func()
		eachFunc  func(float64) error
	}{
		{
			name:      "add n",
			slc:       []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n float64) error {
				v := rabbit.Load().(float64)
				v = v + n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name:      "subtract n",
			slc:       []float64{1.5, 2.25, 6.75, 8.5, 12.25},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n float64) error {
				v := rabbit.Load().(float64)
				v = v - n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name:      "errors",
			slc:       []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected:  3,
			expected2: myErr,
			before:    func() { rabbit.Store(float64(0)) },
			eachFunc: func(n float64) error {
				v := rabbit.Load().(float64)
				if n > 9 {
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
			e, i := Float64(tt.slc).TryEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func TestFloat64_IfEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store(float64(0))

	tests := []struct {
		name      string
		slc       []float64
		expected  int
		expected2 bool
		before    func()
		err       error
		eachFunc  func(float64) bool
	}{
		{
			name:      "all return true",
			slc:       []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n float64) bool {
				v := rabbit.Load().(float64)
				v = v + n
				rabbit.Store(v)
				return true
			},
		},
		{
			name:      "subtract n",
			slc:       []float64{1.5, 2.25, 6.75, 8.5, 12.25},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n float64) bool {
				v := rabbit.Load().(float64)
				v = v - n
				rabbit.Store(v)
				return true
			},
		},
		{
			name:      "breaking",
			slc:       []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected:  3,
			expected2: false,
			before:    func() { rabbit.Store(float64(0)) },
			eachFunc: func(n float64) bool {
				v := rabbit.Load().(float64)
				if n > 9 {
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
			e, i := Float64(tt.slc).IfEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func TestFloat64_Map(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slc      []float64
		expected []float64
		mapFunc  func(float64) float64
	}{
		{
			name:     "add 3.5",
			slc:      []float64{1.23, 2.34, 5.67, 11.5, 13.25, 15.25},
			expected: []float64{4.73, 5.84, 9.17, 15, 16.75, 18.75},
			mapFunc:  func(n float64) float64 { return n + 3.5 },
		},
		{
			name:     "multiply by 2",
			slc:      []float64{1.5, 2.5, 6.5, 8.5, 12.5, 15.5, 17.5},
			expected: []float64{3.0, 5.0, 13.0, 17.0, 25.0, 31.0, 35.0},
			mapFunc:  func(n float64) float64 { return n * 2 },
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Float64(tt.slc).Map(tt.mapFunc)
			if !reflect.DeepEqual(tt.expected, tt.slc) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func TestFloat64_Chunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		size     int
		slc      []float64
		expected [][]float64
	}{
		{
			name:     "chunks of 2 no remainder",
			size:     2,
			slc:      []float64{1.1, 2.2, 5.5, 11.11, 13.13, 15.15},
			expected: [][]float64{[]float64{1.1, 2.2}, []float64{5.5, 11.11}, []float64{13.13, 15.15}},
		},
		{
			name:     "chunks of 2 with remainder",
			size:     2,
			slc:      []float64{1.1, 2.2, 5.5, 11.11, 13.13, 15.15, 17.17},
			expected: [][]float64{[]float64{1.1, 2.2}, []float64{5.5, 11.11}, []float64{13.13, 15.15}, []float64{17.17}},
		},
		{
			name:     "chunks of 100",
			size:     100,
			slc:      []float64{1.1, 2.2, 5.5, 11.11, 13.13, 15.15, 17.17},
			expected: [][]float64{[]float64{1.1, 2.2, 5.5, 11.11, 13.13, 15.15, 17.17}},
		},
		{
			name:     "chunks of 4",
			size:     4,
			slc:      []float64{1.1, 2.2, 5.5, 11.11, 13.13, 15.15, 17.17},
			expected: [][]float64{[]float64{1.1, 2.2, 5.5, 11.11}, []float64{13.13, 15.15, 17.17}},
		},
		{
			name:     "chunks of 5",
			size:     5,
			slc:      []float64{1.1, 2.2, 5.5, 11.11, 13.13, 15.15, 17.17},
			expected: [][]float64{[]float64{1.1, 2.2, 5.5, 11.11, 13.13}, []float64{15.15, 17.17}},
		},
		{
			name:     "empty slice",
			size:     5,
			slc:      []float64{},
			expected: [][]float64{},
		},
		{
			name:     "invalid chunk size",
			size:     -1,
			slc:      []float64{1.1, 2.2, 5.5, 11.11, 13.13, 15.15, 17.17},
			expected: [][]float64{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := Float64(tt.slc).Chunk(tt.size)
			if !reflect.DeepEqual(tt.expected, out) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func BenchmarkFloat64_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).IndexOf(1)
			}
		})
	}
}

func BenchmarkFloat64_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).Contains(1)
			}
		})
	}
}

func BenchmarkFloat64_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).SortAsc()
			}
		})
	}
}

func BenchmarkFloat64_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).SortDesc()
			}
		})
	}
}

func BenchmarkFloat64_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).Reverse()
			}
		})
	}
}

func BenchmarkFloat64_Filter(b *testing.B) {
	benchFunc := func(n float64) bool {
		return true
	}
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).Filter(benchFunc)
			}
		})
	}
}

func BenchmarkFloat64_Each(b *testing.B) {
	benchFunc := func(n float64) { n = 0 }
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).Each(benchFunc)
			}
		})
	}
}

func BenchmarkFloat64_Map(b *testing.B) {
	benchFunc := func(n float64) float64 {
		return n
	}
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).Map(benchFunc)
			}
		})
	}
}

func BenchmarkFloat64_Unique(b *testing.B) {
	benchmarks := []struct {
		name string
		slc  []float64
	}{
		{
			name: "10 elements",
			slc:  genFloat64(b, 10),
		},
		{
			name: "100 elements",
			slc:  genFloat64(b, 100),
		},
		{
			name: "1000 elements",
			slc:  genFloat64(b, 1000),
		},
		{
			name: "10000 elements",
			slc:  genFloat64(b, 10000),
		},
		{
			name: "100000 elements",
			slc:  genFloat64(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Float64(bm.slc).Unique()
			}
		})
	}
}
