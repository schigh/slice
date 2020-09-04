package slice

import (
	"math/rand"
	"testing"
)

func genInt(b *testing.B, size int) []int {
	b.Helper()

	out := make([]int, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := 1
		if r%2 == 0 {
			sign = -1
		}
		out[i] = rand.Int() * sign
	}

	return out
}

func genInt8(b *testing.B, size int) []int8 {
	b.Helper()

	out := make([]int8, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := 1
		if r%2 == 0 {
			sign = -1
		}
		out[i] = int8(rand.Int() * sign)
	}

	return out
}

func genInt16(b *testing.B, size int) []int16 {
	b.Helper()

	out := make([]int16, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := 1
		if r%2 == 0 {
			sign = -1
		}
		out[i] = int16(rand.Int() * sign)
	}

	return out
}

func genInt32(b *testing.B, size int) []int32 {
	b.Helper()

	out := make([]int32, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := int32(1)
		if r%2 == 0 {
			sign = int32(-1)
		}
		out[i] = rand.Int31() * sign
	}

	return out
}

func genInt64(b *testing.B, size int) []int64 {
	b.Helper()

	out := make([]int64, size)
	for i := 0; i < size; i++ {
		r := rand.Int()
		sign := int64(1)
		if r%2 == 0 {
			sign = int64(-1)
		}
		out[i] = rand.Int63() * sign
	}

	return out
}

func genUInt(b *testing.B, size int) []uint {
	b.Helper()

	out := make([]uint, size)
	for i := 0; i < size; i++ {
		out[i] = uint(rand.Uint32())
	}

	return out
}

func genUInt8(b *testing.B, size int) []uint8 {
	b.Helper()

	out := make([]uint8, size)
	for i := 0; i < size; i++ {
		out[i] = uint8(rand.Uint32())
	}

	return out
}

func genUInt16(b *testing.B, size int) []uint16 {
	b.Helper()
	out := make([]uint16, size)
	for i := 0; i < size; i++ {
		out[i] = uint16(rand.Uint32())
	}

	return out
}

func genUInt32(b *testing.B, size int) []uint32 {
	b.Helper()

	out := make([]uint32, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Uint32()
	}

	return out
}

func genUInt64(b *testing.B, size int) []uint64 {
	b.Helper()

	out := make([]uint64, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Uint64()
	}

	return out
}

func genFloat32(b *testing.B, size int) []float32 {
	b.Helper()

	out := make([]float32, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Float32()
	}

	return out
}

func genFloat64(b *testing.B, size int) []float64 {
	b.Helper()

	out := make([]float64, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Float64()
	}

	return out
}

func genUIntPtr(b *testing.B, size int) []uintptr {
	b.Helper()

	out := make([]uintptr, size)
	for i := 0; i < size; i++ {
		out[i] = uintptr(rand.Uint64())
	}

	return out
}

func genString(b *testing.B, size int) []string {
	b.Helper()

	out := make([]string, size)
	charset := `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	ll := len(charset)

	for i := 0; i < size; i++ {
		l := rand.Intn(32)
		b := make([]byte, l)
		for k := 0; k < l; k++ {
			b[k] = charset[rand.Intn(ll)]
		}
		out[i] = string(b)
	}

	return out
}
