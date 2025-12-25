package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian[T ~uint16 | ~uint32 | ~uint64](val T) T {
	var res T
	//bs := unsafe.Sizeof(val)
	//			   56		48		40		  32		24		16		8
	// 11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111

	switch any(val).(type) {
	case uint64:
		a1 := byte(val >> 56)
		a2 := byte(val >> 48)
		a3 := byte(val >> 40)
		a4 := byte(val >> 32)
		a5 := byte(val >> 24)
		a6 := byte(val >> 16)
		a7 := byte(val >> 8)
		a8 := byte(val)

		res = T(a8)<<56 |
			T(a7)<<48 |
			T(a6)<<40 |
			T(a5)<<32 |
			T(a4)<<24 |
			T(a3)<<16 |
			T(a2)<<8 |
			T(a1)

		// 11111111 11111111 11111111 11111111
	case uint32:
		a1 := byte(val >> 24)
		a2 := byte(val >> 16)
		a3 := byte(val >> 8)
		a4 := byte(val)

		res = T(a4)<<24 |
			T(a3)<<16 |
			T(a2)<<8 |
			T(a1)
	case uint16:
		a1 := byte(val >> 8)
		a2 := byte(val)

		res = T(a2)<<8 | T(a1)
	}

	return res
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func runTests[T ~uint16 | ~uint32 | ~uint64](t *testing.T, cases map[string]struct{ number, result T }) {
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, ToLittleEndian(tc.number))
		})
	}
}

func TestToLittleEndian(t *testing.T) {
	runTests(t, map[string]struct{ number, result uint64 }{
		"uint64 case #1": {0x0102030401020304, 0x0403020104030201},
		"uint64 case #2": {0x0000FFFF0000FFFF, 0xFFFF0000FFFF0000},
	})

	runTests(t, map[string]struct{ number, result uint32 }{
		"uint32 case #1": {0x01020304, 0x04030201},
		"uint32 case #2": {0x0000FFFF, 0xFFFF0000},
	})

	runTests(t, map[string]struct{ number, result uint16 }{
		"uint16 case #1": {0x0102, 0x0201},
		"uint16 case #2": {0x1011, 0x1110},
	})
}
