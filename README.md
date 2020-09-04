[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Build Status](https://travis-ci.com/schigh/slice.svg?branch=master)](https://travis-ci.com/schigh/slice)
[![codecov](https://codecov.io/gh/schigh/slice/branch/master/graph/badge.svg)](https://codecov.io/gh/schigh/slice)
[![Go Report Card](https://goreportcard.com/badge/github.com/schigh/slice)](https://goreportcard.com/report/github.com/schigh/slice)
[![Godocs](https://img.shields.io/badge/go.dev-docs-blue.svg?logo=go)](https://pkg.go.dev/github.com/schigh/slice)

> Some great and powerful gopher once said you shouldn't put this kind of stuff in a library because it's so simple.  Until you have to write a contains method for the thousandth time.

# SLICE

Slice bolts on some generics-ish functionality to native Go data types when represented as a slice.
Currently, the following types are supported:

|     Go Type   |   Slice Type   |
| ------------- | -------------- |
| `[]int`       | `Int`          |
| `[]int8`      | `Int8`         |
| `[]int16`     | `Int16`        |
| `[]int32`     | `Int32`        |
| `[]int64`     | `Int64`        |
| `[]uint`      | `UInt`         |
| `[]uint8`     | `UInt8`        |
| `[]uint16`    | `UInt16`       |
| `[]uint32`    | `UInt32`       |
| `[]uint64`    | `UInt64`       |
| `[]float32`   | `Float32`      |
| `[]float64`   | `Float64`      |
| `[]string`    | `String`       |

For the above types, the following operations are supported (_x_ is the type in the slice):

| Function               | Description                                                  |
| ---------------------- | ------------------------------------------------------------ |
| `IndexOf(x)`           | Get the first index of the searched element in the slice.  If the element is not found, this function returns -1 |
| `Contains(x)`          | Test to see if slice contains element _x_                    |
| `Unique()`               | Returns unique items in slice.  The first occurence of an element is the one that is kept |
| `SortAsc()`              | Sort the slice in ascending order                            |
| `SortDesc()`             | Sort the slice in descending order                           |
| `Reverse()`              | Reverses the slice                                           |
| `Filter(func(x) bool)` | Applies a filter function to every item in the slice and return all items where the filter returns true |
| `Map(func(x) x)`	| Iterate over the slice and replace the current value with the computed value |
| `Each(func(x))`  | Iterate over the slice (non-mutating) |
| `TryEach(func(x) error) (int, error)` | Iterate over the slice, and halt if an error is returned from user func.  Return index of the failed member and the caught error |
| `IfEach(func(x) bool) (int, bool)` | Iterate over the slice, and halt if false is returned from user func.  Return the index of the element that caused the func to return false, and a bool that is true if every member of the slice returned true with the function applied.  If all elements return true, the index returned is `-1` |
| `Chunk(size)` | Splits the slice into chunks of a specified size |
| `Value()` | Returns the native type slice value |

#### Performance

Performance benchmarks can be found in [benchmarks.txt](./_data/benchmark.txt).

**A note about performance, benchmarks, etc**

Generics are a way to apply parametric polymorphism to common aggregates.  
They are generally very good at hiding complexity, and as a corollary, hide 
cost as well.  With that said, I encourage you to peek at the performance 
benchmarks I've included in this repo.  One thing that should stand out 
immediately is that cpu time and memory increase proportionally with load.  
This should not be surprising at all, but I should point out here that if you 
expect `Filter()` or `Unique()` to be performant with 100,000 elements, then it is 
likely you may need to pick a more bespoke and optimized approach.



#### Some examples...
```go
package main

import (
	"fmt"
	
	"github.com/schigh/slice"
)

func main() {
	orig := []string{"foo", "bar", "baz", "fizz", "buzz"}
	startsWithF := slice.String(orig).Filter(func(s string) bool {
		return len(s) > 0 && s[0] == 'f'
	}).Value()
	
	fmt.Println(startsWithF)
	// ["foo", "fizz"]
}
```
```go
package main

import (
	"fmt"
	
	"github.com/schigh/slice"
)

func main() {
	orig := []int{9,1,6,2,3,4,3,4,5,7,8,9,0}
	unique := slice.Int(orig).Unique().SortDesc().Value()
	
	fmt.Println(unique)
	// [9,8,7,6,5,4,3,2,1,0]
}
```

Check out the GoDoc for more information.

##  Generator

The slice generator uses `go:generate` to add slice functionality to your existing structs. You may choose which features you'd like to add by setting them in the `generate` command.  For example:

```
//go:generate slices User map filter each
```

Will generate the `Map`, `Filter`, and `Each` functionality (see below) on a User struct's slice type.  You could also just say you want everything:

```
//go:generate slices User all
```

This will generate all functions produced by the tool.

### Generator Options

#### Closures
Some generated functions take as their receiver a function that operates on every member of the slice.  By default, these receivers use a function where each member is passed by reference.  For example, the generated `Filter` function on a `User` struct:

```go
// Filter evaluates every element in the slice, and returns all User 
// instances where the eval function returns true
func (slc UserSlice) Filter(f func(*User) bool) UserSlice {
	out := make([]User, 0, len(slc))
	for _, v := range slc {
		if f(&v) {
			out = append(out, v)
		}
	}

	return UserSlice(out)
}
```

The user-defined function receives a pointer to `User` by default.

If instead you want your slice functions to operate by value, you can modify your generator tags like so:

```
//go:generate slices User filter(byvalue)
```

This would produce the following function:

```go
// Filter evaluates every element in the slice, and returns all User 
// instances where the eval function returns true
func (slc UserSlice) Filter(f func(User) bool) UserSlice {
	out := make([]User, 0, len(slc))
	for _, v := range slc {
		if f(v) {
			out = append(out, v)
		}
	}

	return UserSlice(out)
}
```

#### Pointer Slices
You can also generate pointer slices by prepending an asterisk to your struct name in the `go generate` directive, like so:

```
//go:generate slices *User all
```

This will generate a type called `UserPtrSlice`, which is an alias of `[]*User`.  All slice functions take a pointer receiver, for example:

```go
func (slc UserPtrSlice) Map(f func(*User) *User) {
	for i, v := range slc {
		slice[i] = f(v)
	}
}
```

