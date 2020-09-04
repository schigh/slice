package slice_test

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/schigh/slice"
)

func ExampleString_Filter() {
	slc := []string{"foo", "bar", "baz", "fizz", "buzz"}

	// starts with f
	fslice := slice.String(slc).Filter(func(s string) bool {
		return s != "" && s[0] == 'f'
	})

	// starts with b
	bslice := slice.String(slc).Filter(func(s string) bool {
		return s != "" && s[0] == 'b'
	})

	fmt.Printf("%v\n", fslice)
	fmt.Printf("%v\n", bslice)

	// output:
	// [foo fizz]
	// [bar baz buzz]
}

func ExampleString_Map() {
	slc := []string{"foo", "bar", "baz", "fizz", "buzz"}

	slice.String(slc).Map(func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Printf("%v\n", slc)

	// output:
	// [FOO BAR BAZ FIZZ BUZZ]
}

func ExampleString_Chunk() {
	slc := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"}

	chunks := slice.String(slc).Chunk(4)

	for _, chunk := range chunks {
		fmt.Printf("%v\n", chunk)
	}

	// output:
	// [a b c d]
	// [e f g h]
	// [i j k l]
	// [m]
}

func ExampleString_Each() {
	var v atomic.Value
	v.Store("")

	slc := []string{"h", "e", "l", "l", "o", ",", " ", "w", "o", "r", "l", "d", "!"}
	slice.String(slc).Each(func(s string) {
		vv := v.Load().(string)
		vv = vv + s
		v.Store(vv)
	})

	vv := v.Load().(string)
	fmt.Println(vv)

	// output:
	// hello, world!
}
