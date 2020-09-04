// Code generated by github.com/schigh/slice/base/base.go. DO NOT EDIT.
{{/*
	.BaseType
	-	the base go type (int, uint, etc)

	.PackageType
	-	the slice alias name (Int, Int16, etc)

*/}}package slice

import (
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
)

func Test{{.PackageType}}_IndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		needle {{.BaseType}}
		expected int
	}{
		{
			name: "only item",
			slc: []{{.BaseType}}{1},
			needle: 1,
			expected: 0,
		},
		{
			name: "at index 1",
			slc: []{{.BaseType}}{0,1,1},
			needle: 1,
			expected: 1,
		},
		{
			name: "missing",
			slc: []{{.BaseType}}{1,2,3,4},
			needle: 5,
			expected: NotInSlice,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := {{.PackageType}}(tt.slc).IndexOf(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func Test{{.PackageType}}_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		needle {{.BaseType}}
		expected bool
	}{
		{
			name: "present",
			slc: []{{.BaseType}}{1,2,3,4},
			needle: 4,
			expected: true,
		},
		{
			name: "not present",
			slc: []{{.BaseType}}{1,2,3,4},
			needle: 5,
			expected: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := {{.PackageType}}(tt.slc).Contains(tt.needle)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test{{.PackageType}}_SortAsc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected []{{.BaseType}}
	}{
		{
			name: "empty",
			slc: []{{.BaseType}}{},
			expected: []{{.BaseType}}{},
		},
		{
			name: "already sorted",
			slc: []{{.BaseType}}{0,1,2,3,4,5},
			expected: []{{.BaseType}}{0,1,2,3,4,5},
		},
		{
			name: "reversed",
			slc: []{{.BaseType}}{5,4,3,2,1,0},
			expected: []{{.BaseType}}{0,1,2,3,4,5},
		},
		{
			name: "random",
			slc: []{{.BaseType}}{3,1,4,5,0,2},
			expected: []{{.BaseType}}{0,1,2,3,4,5},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := {{.PackageType}}(tt.slc).SortAsc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test{{.PackageType}}_SortDesc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected []{{.BaseType}}
	}{
		{
			name: "empty",
			slc: []{{.BaseType}}{},
			expected: []{{.BaseType}}{},
		},
		{
			name: "already sorted",
			slc: []{{.BaseType}}{5,4,3,2,1,0},
			expected: []{{.BaseType}}{5,4,3,2,1,0},
		},
		{
			name: "reversed",
			slc: []{{.BaseType}}{0,1,2,3,4,5},
			expected: []{{.BaseType}}{5,4,3,2,1,0},
		},
		{
			name: "random",
			slc: []{{.BaseType}}{3,1,4,5,0,2},
			expected: []{{.BaseType}}{5,4,3,2,1,0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := {{.PackageType}}(tt.slc).SortDesc().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test{{.PackageType}}_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected []{{.BaseType}}
	}{
		{
			name: "unaffected",
			slc: []{{.BaseType}}{0,1,2,3,4,5},
			expected: []{{.BaseType}}{0,1,2,3,4,5},
		},
		{
			name: "one extra five",
			slc: []{{.BaseType}}{5,0,1,2,3,4,5},
			expected: []{{.BaseType}}{5,0,1,2,3,4},
		},
		{
			name: "extras everywhere",
			slc: []{{.BaseType}}{0,0,1,0,1,2,2,2,3,0,3,4,2,3,4,4,2,1,0},
			expected: []{{.BaseType}}{0,1,2,3,4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := {{.PackageType}}(tt.slc).Unique().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test{{.PackageType}}_Reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected []{{.BaseType}}
	}{
		{
			name: "even length",
			slc: []{{.BaseType}}{0,1,2,3,4,5},
			expected: []{{.BaseType}}{5,4,3,2,1,0},
		},
		{
			name: "odd length",
			slc: []{{.BaseType}}{0,1,2,3,4,5,6},
			expected: []{{.BaseType}}{6,5,4,3,2,1,0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := {{.PackageType}}(tt.slc).Reverse().Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test{{.PackageType}}_Filter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected []{{.BaseType}}
		filterFunc func({{.BaseType}}) bool
	}{
		{
			name: "gt 10",
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: []{{.BaseType}}{11, 13, 15},
			filterFunc: func(n {{.BaseType}}) bool { return n > 10 },
		},
		{
			name: "mod 3",
			slc: []{{.BaseType}}{1, 2, 6, 11, 12, 15, 17},
			expected: []{{.BaseType}}{6, 12, 15},
			filterFunc: func(n {{.BaseType}}) bool { return n%3 == 0 },
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := {{.PackageType}}(tt.slc).Filter(tt.filterFunc).Value()
			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test{{.PackageType}}_Each(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store({{.BaseType}}({{.ZeroValue}}))

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected {{.BaseType}}
		eachFunc func({{.BaseType}})
	}{
		{
			name: "add n",
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n {{.BaseType}}) {
				v := rabbit.Load().({{.BaseType}})
				v = v + n
				rabbit.Store(v)
			},
		},
		{
			name: "subtract n",
			slc: []{{.BaseType}}{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n {{.BaseType}}) {
				v := rabbit.Load().({{.BaseType}})
				v = v - n
				rabbit.Store(v)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			{{.PackageType}}(tt.slc).Each(tt.eachFunc)
			v := rabbit.Load().({{.BaseType}})
			if tt.expected != v {
				t.Errorf("expected %v, got %v", tt.expected, rabbit)
			}
		})
	}
}

func Test{{.PackageType}}_TryEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store({{.BaseType}}({{.ZeroValue}}))

	myErr := errors.New("i am an error")
	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected int
		expected2 error
		before func()
		eachFunc func({{.BaseType}}) error
	}{
		{
			name: "add n",
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: NotInSlice,
			expected2: nil,
			eachFunc: func(n {{.BaseType}}) error {
				v := rabbit.Load().({{.BaseType}})
				v = v + n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name: "subtract n",
			slc: []{{.BaseType}}{1, 2, 6, 8, 12},
			expected: NotInSlice,
			expected2: nil,
			eachFunc: func(n {{.BaseType}}) error {
				v := rabbit.Load().({{.BaseType}})
				v = v - n
				rabbit.Store(v)
				return nil
			},
		},
		{
			name: "errors",
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: 3,
			expected2: myErr,
			before: func() { rabbit.Store({{.BaseType}}({{.ZeroValue}})) },
			eachFunc: func(n {{.BaseType}}) error {
				v := rabbit.Load().({{.BaseType}})
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
			e, i := {{.PackageType}}(tt.slc).TryEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func Test{{.PackageType}}_IfEach(t *testing.T) {
	t.Parallel()

	var rabbit atomic.Value
	rabbit.Store({{.BaseType}}({{.ZeroValue}}))

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected int
		expected2 bool
		before func()
		err error
		eachFunc func({{.BaseType}}) bool
	}{
		{
			name: "all return true",
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: NotInSlice,
			expected2: true,
			eachFunc: func(n {{.BaseType}}) bool {
				v := rabbit.Load().({{.BaseType}})
				v = v + n
				rabbit.Store(v)
				return true
			},
		},
		{
			name: "subtract n",
			slc: []{{.BaseType}}{1, 2, 6, 8, 12},
			expected: NotInSlice,
			expected2: true,
			eachFunc: func(n {{.BaseType}}) bool {
				v := rabbit.Load().({{.BaseType}})
				v = v - n
				rabbit.Store(v)
				return true
			},
		},
		{
			name: "breaking",
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: 3,
			expected2: false,
			before: func() { rabbit.Store({{.BaseType}}({{.ZeroValue}})) },
			eachFunc: func(n {{.BaseType}}) bool {
				v := rabbit.Load().({{.BaseType}})
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
			e, i := {{.PackageType}}(tt.slc).IfEach(tt.eachFunc)
			if tt.expected != e {
				t.Errorf("expected %v, got %v", tt.expected, e)
			}
			if tt.expected2 != i {
				t.Errorf("expected %v, got %v", tt.expected2, i)
			}
		})
	}
}

func Test{{.PackageType}}_Map(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		slc []{{.BaseType}}
		expected []{{.BaseType}}
		mapFunc func({{.BaseType}}) {{.BaseType}}
	}{
		{
			name: "add 3",
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: []{{.BaseType}}{4, 5, 8, 14, 16, 18},
			mapFunc: func(n {{.BaseType}}) {{.BaseType}} { return n + 3 },
		},
		{
			name: "set mod 2",
			slc: []{{.BaseType}}{1, 2, 6, 8, 12, 15, 17},
			expected: []{{.BaseType}}{1, 0, 0, 0, 0, 1, 1},
			mapFunc: func(n {{.BaseType}}) {{.BaseType}} { return {{.BaseType}}(n%2) },
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			{{.PackageType}}(tt.slc).Map(tt.mapFunc)
			if !reflect.DeepEqual(tt.expected, tt.slc) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func Test{{.PackageType}}_Chunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		size int
		slc []{{.BaseType}}
		expected [][]{{.BaseType}}
	}{
		{
			name: "chunks of 2 no remainder",
			size: 2,
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15},
			expected: [][]{{.BaseType}}{[]{{.BaseType}}{1, 2}, []{{.BaseType}}{5, 11}, []{{.BaseType}}{13, 15}},
		},
		{
			name: "chunks of 2 with remainder",
			size: 2,
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15, 17},
			expected: [][]{{.BaseType}}{[]{{.BaseType}}{1, 2}, []{{.BaseType}}{5, 11}, []{{.BaseType}}{13, 15}, []{{.BaseType}}{17}},
		},
		{
			name: "chunks of 100",
			size: 100,
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15, 17},
			expected: [][]{{.BaseType}}{[]{{.BaseType}}{1, 2, 5, 11, 13, 15, 17}},
		},
		{
			name: "chunks of 4",
			size: 4,
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15, 17},
			expected: [][]{{.BaseType}}{[]{{.BaseType}}{1, 2, 5, 11}, []{{.BaseType}}{13, 15, 17}},
		},
		{
			name: "chunks of 5",
			size: 5,
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15, 17},
			expected: [][]{{.BaseType}}{[]{{.BaseType}}{1, 2, 5, 11, 13}, []{{.BaseType}}{15, 17}},
		},
		{
			name: "empty slice",
			size: 5,
			slc: []{{.BaseType}}{},
			expected: [][]{{.BaseType}}{},
		},
		{
			name: "invalid chunk size",
			size: -1,
			slc: []{{.BaseType}}{1, 2, 5, 11, 13, 15, 17},
			expected: [][]{{.BaseType}}{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := {{.PackageType}}(tt.slc).Chunk(tt.size)
			if !reflect.DeepEqual(tt.expected, out) {
				t.Errorf("expected %v, got %v", tt.expected, tt.slc)
			}
		})
	}
}

func Benchmark{{.PackageType}}_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).IndexOf(1)
			}
		})
	}
}

func Benchmark{{.PackageType}}_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).Contains(1)
			}
		})
	}
}

func Benchmark{{.PackageType}}_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).SortAsc()
			}
		})
	}
}

func Benchmark{{.PackageType}}_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).SortDesc()
			}
		})
	}
}

func Benchmark{{.PackageType}}_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).Reverse()
			}
		})
	}
}

func Benchmark{{.PackageType}}_Filter(b *testing.B) {
	benchFunc := func(n {{.BaseType}}) bool {
		return true
	}
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).Filter(benchFunc)
			}
		})
	}
}

func Benchmark{{.PackageType}}_Each(b *testing.B) {
	benchFunc := func(n {{.BaseType}}) { n = {{.ZeroValue}} }
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).Each(benchFunc)
			}
		})
	}
}

func Benchmark{{.PackageType}}_Map(b *testing.B) {
	benchFunc := func(n {{.BaseType}}) {{.BaseType}} {
		return n
	}
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).Map(benchFunc)
			}
		})
	}
}

func Benchmark{{.PackageType}}_Unique(b *testing.B) {
	benchmarks := []struct {
		name string
		slc []{{.BaseType}}
	}{
		{
			name: "10 elements",
			slc: gen{{.PackageType}}(b, 10),
		},
		{
			name: "100 elements",
			slc: gen{{.PackageType}}(b, 100),
		},
		{
			name: "1000 elements",
			slc: gen{{.PackageType}}(b, 1000),
		},
		{
			name: "10000 elements",
			slc: gen{{.PackageType}}(b, 10000),
		},
		{
			name: "100000 elements",
			slc: gen{{.PackageType}}(b, 100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				{{.PackageType}}(bm.slc).Unique()
			}
		})
	}
}